package potto

import (
	"log"
	"net/http"
)

func ActionHandler(ctx Ctx) {
	text := ctx.ParamSingle("text")
	trigger := ctx.ParamSingle("trigger_word")

	cmd := NewCommand(text, trigger)

	action, ok := ctx.Actions()[cmd.Name]
	if !ok {
		log.Printf("unknown action name: %s", cmd.Name)
		ctx.RenderJSON(http.StatusOK, NewResponse("unknown action name"))
		return
	}

	res, err := action(ctx, cmd.Args)
	if err != nil {
		log.Println(err)
		ctx.RenderJSON(http.StatusOK, NewResponse(err.Error()))
		return
	}

	ctx.RenderJSON(http.StatusOK, res)
}
