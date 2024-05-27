package pow

import (
	"runtime"

	goasvc "github.com/kormiltsev/proofofwork/api/gen/words"
)

// Logger implements logger for tracing errors occurred.
type Logger interface {
	Error(fmt string, ctx ...interface{})
}

// InternalError error generation using provided error, code and description
func internal(code string, lg Logger, err error, a ...interface{}) error {
	pc, fn, line, _ := runtime.Caller(1)
	lg.Error("internal error occurred", "err", err.Error(), "file", fn, "line", line, "function", runtime.FuncForPC(pc).Name(), "msg", err)

	return &goasvc.InternalError{
		MsgCode: code,
	}
}

// BadRequest error generation using provided error, code and description
func badRequest(code string, lg Logger, err error, a ...interface{}) error {
	pc, fn, line, _ := runtime.Caller(1)
	lg.Error("bad request", "err", err.Error(), "file", fn, "line", line, "function", runtime.FuncForPC(pc).Name())

	return &goasvc.BadRequestError{
		MsgCode: code,
	}
}

// NotFound error generation using provided error, code and description
func notFound(code string, lg Logger, err error, a ...interface{}) error {
	pc, fn, line, _ := runtime.Caller(1)
	lg.Error("not found", "err", err.Error(), "file", fn, "line", line, "function", runtime.FuncForPC(pc).Name())

	return &goasvc.NotFoundError{
		MsgCode: code,
	}
}

// Conflict error generation using provided error, code and description
func conflict(code string, lg Logger, err error, a ...interface{}) error {
	pc, fn, line, _ := runtime.Caller(1)
	lg.Error("conflict", "err", err.Error(), "file", fn, "line", line, "function", runtime.FuncForPC(pc).Name())

	return &goasvc.ConflictError{
		MsgCode: code,
	}
}
