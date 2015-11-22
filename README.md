# potto

a tiny framework to handle Slack Outgoing WebHooks for golang

## Example

``` go
package main

import (
	"strings"

	"github.com/m0t0k1ch1/potto"
)

func Ping(ctx potto.Ctx, args potto.ActionArgs) (*potto.Response, error) {
	return potto.NewResponse("pong"), nil
}

func Say(ctx potto.Ctx, args potto.ActionArgs) (*potto.Response, error) {
	text := strings.Join(args, " ")
	return potto.NewResponse(text), nil
}

func main() {
	p := potto.New()
	p.AddAction("ping", Ping)
	p.AddAction("say", Say)
	p.Run(":8080")
}
```
