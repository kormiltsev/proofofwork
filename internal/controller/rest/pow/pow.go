package pow

import (
	"context"
	"fmt"

	"github.com/inconshreveable/log15"
	"gopkg.in/tomb.v2"

	goasvc "github.com/kormiltsev/proofofwork/api/gen/words"
	"github.com/kormiltsev/proofofwork/internal/utils"
)

// Pow is a service to proof a work.
type Pow interface {
	NewTask() (string, int, error)
	Validate(string) (bool, error)
}

// Job makes some job after proof-of-work.
type Job interface {
	GetQuote(context.Context) string
}

// Controller implements the apps resource.
type Controller struct {
	log  log15.Logger
	tomb *tomb.Tomb

	pow Pow
	job Job
}

// NewController creates a job controller.
func NewController(t *tomb.Tomb, pow Pow, job Job) *Controller {
	return &Controller{
		log:  log15.New("controller", "job"),
		tomb: t,
		pow:  pow,
		job:  job,
	}
}

// Request is first step. Returns Task.
func (c *Controller) Request(_ context.Context) (*goasvc.WordsTask, error) {
	request, diff, err := c.pow.NewTask()
	if err != nil {
		c.log.Error("no task returned", "err", err)
		return nil, internal("no task generated", c.log, err)
	}

	res := goasvc.WordsTask{
		Hash:       utils.String(request),
		Difficulty: utils.Int(diff),
	}
	return &res, nil
}

// Words validates the solution and returns job's result.
func (c *Controller) Words(ctx context.Context, p *goasvc.WordsPayload) (*goasvc.WordsResult, error) {
	approved, err := c.pow.Validate(utils.StringUnref(p.Solution))
	if err != nil {
		c.log.Error("validation failed", "err", err)
		return nil, internal("validation failed", c.log, err)
	}

	if !approved {
		c.log.Debug("validation failed", "response", p.Solution)
		return nil, badRequest("validation failed", c.log, fmt.Errorf("validation failed"))
	}
	res := goasvc.WordsResult{
		Quote: utils.String(c.job.GetQuote(ctx)),
	}
	return &res, nil
}
