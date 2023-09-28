# go-playground/validator

This is a validator extended by [go-playground/validator@v10](https://github.com/go-playground/validator/tree/master).<br>
It can be used by hertz as a validator.

## Usage
* default validate tag: `binding`
* validate rule:  [go-playground/validator@v10](https://github.com/go-playground/validator/tree/master)

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/binding/go_playground"
)

type Test struct {
	Q int `query:"q" hhh:"gte=0,lte=130"`
}

func main() {
	vd := go_playground.NewValidator()
	vd.SetValidateTag("hhh") // the default validate tag is 'binding'
	// If you need to configure the validator, you can use the following
	//vdEngine := vd.Engine().(*validator.Validate)
	//vdEngine.RegisterCustomTypeFunc()
	h := server.New(server.WithCustomValidator(vd))
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		var req Test
		err := ctx.BindAndValidate(&req)
		if err != nil {
			fmt.Println(err)
			ctx.String(400, err.Error())
			return
		}
		fmt.Println(req)
		ctx.JSON(200, req)
	})

	go h.Spin()
	time.Sleep(100 * time.Millisecond)
	hc := http.Client{Timeout: 1000 * time.Second}
	hc.Get("http://127.0.0.1:8888/ping?q=12444")
	time.Sleep(1 * time.Second)
}

```