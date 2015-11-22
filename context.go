package potto

import (
	"net/http"

	"github.com/m0t0k1ch1/ksatriya"
)

type Ctx interface {
	ksatriya.Ctx
	Actions() map[string]Action
}

type CtxBuilder func(w http.ResponseWriter, req *http.Request, args Args) Ctx

type Context struct {
	*ksatriya.Context
	actions map[string]Action
}

func (ctx *Context) Actions() map[string]Action {
	return ctx.actions
}
