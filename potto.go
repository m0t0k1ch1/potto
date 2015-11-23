package potto

import (
	"net/http"

	"github.com/m0t0k1ch1/ksatriya"
)

type Args struct {
	ksatriya.Args
}

type HandlerFunc func(Ctx)

type Action func(Ctx, ActionArgs) (*Response, error)
type ActionArgs []string

type Potto struct {
	*ksatriya.Ksatriya
	actions map[string]Action
}

func (p *Potto) SetCtxBuilder(f CtxBuilder) {
	p.Ksatriya.SetCtxBuilder(func(w http.ResponseWriter, req *http.Request, args ksatriya.Args) ksatriya.Ctx {
		return f(w, req, Args{args})
	})
}

func (p *Potto) AddRoute(method, path string, hf HandlerFunc) {
	p.Ksatriya.AddRoute(method, path, func(kctx ksatriya.Ctx) {
		hf(kctx.(Ctx))
	})
}

func (p *Potto) Actions() map[string]Action {
	return p.actions
}

func (p *Potto) AddAction(name string, action Action) {
	p.actions[name] = action
}

func (p *Potto) NewContext(w http.ResponseWriter, req *http.Request, args Args) Ctx {
	kctx := ksatriya.NewContext(w, req, args.Args)

	return &Context{
		Context: kctx.(*ksatriya.Context),
		actions: p.Actions(),
	}
}

func New() *Potto {
	p := &Potto{
		Ksatriya: ksatriya.New(),
		actions:  map[string]Action{},
	}

	p.SetCtxBuilder(p.NewContext)
	p.POST("/", ActionHandler)

	return p
}
