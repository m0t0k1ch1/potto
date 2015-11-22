package potto

import (
	"log"
	"net/http"

	"github.com/m0t0k1ch1/ksatriya"
)

func ActionHandler(kctx ksatriya.Ctx) {
	ctx := kctx.(Ctx)

	text := ctx.ParamSingle("text")
	trigger := ctx.ParamSingle("trigger_word")

	cmd := NewCommand(text, trigger)

	action, ok := ctx.Actions()[cmd.Name]
	if !ok {
		log.Printf("unknown action name: %s", cmd.Name)
		ctx.JSON(http.StatusOK, NewResponse("unknown action name"))
		return
	}

	res, err := action(ctx, cmd.Args)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, NewResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, res)
}
