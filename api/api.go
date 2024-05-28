package api

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/inconshreveable/log15"

	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
	"gopkg.in/tomb.v2"

	powsrv "github.com/kormiltsev/proofofwork/api/gen/http/words/server"
	powsvc "github.com/kormiltsev/proofofwork/api/gen/words"
	powrest "github.com/kormiltsev/proofofwork/internal/controller/rest/pow"
	jobsc "github.com/kormiltsev/proofofwork/internal/service/job"
	powsc "github.com/kormiltsev/proofofwork/internal/service/pow"
)

// Run executes the service words of wosdome as REST API
func Run(ctx context.Context, t *tomb.Tomb, httpListener net.Listener) error {
	var err error
	httpParams := runParameters{
		l:          httpListener,
		tomb:       t,
		powService: powsc.New(),
		jobService: jobsc.New(),
	}

	runHTTP(ctx, &httpParams)

	select {
	case <-ctx.Done():
		log15.Info("application context is done")
		t.Kill(nil)
		return ctx.Err()
	case <-t.Dead():
		log15.Info("application exit due to error", "err", err)
		return t.Err()
	}
}

type runParameters struct {
	l          net.Listener
	tomb       *tomb.Tomb
	powService *powsc.ProofOfWork
	jobService *jobsc.QuoteService
}

func runHTTP(ctx context.Context, p *runParameters) {
	log15.Info("running http server", "addr", p.l.Addr())

	logger := log.New(os.Stderr, "[words-of-wisdom] ", log.Ltime)
	adapter := middleware.NewLogger(logger)
	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder
	mux := goahttp.NewMuxer()

	// handlers and middleware
	var handler http.Handler = newHandler(mux, "")
	handler = httpmdlwr.RequestID()(handler)
	handler = httpmdlwr.Log(adapter)(handler)

	// init controllers
	powSvc := powrest.NewController(p.tomb, p.powService, p.jobService)

	// init endpoints
	powEndpoints := powsvc.NewEndpoints(powSvc)

	eh := errorHandler(logger)

	// init servers
	powSrv := powsrv.New(powEndpoints, mux, dec, enc, eh, nil)

	servers := goahttp.Servers{
		powSrv,
	}

	servers.Use(httpmdlwr.Debug(mux, os.Stdout))

	// mount servers
	powsrv.Mount(mux, powSrv)

	srv := &http.Server{
		Handler: handler,
	}

	p.tomb.Go(func() error {
		return srv.Serve(p.l)
	})

	p.tomb.Go(func() error {
		<-p.tomb.Dying()
		log15.Info("shutdown server")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
		return nil
	})
}

func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}
