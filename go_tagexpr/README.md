# bytedance/go-tagexpr

This is a "binder and validator" extended by [bytedance/go-tagexpr](https://github.com/bytedance/go-tagexpr).<br>
It can be used by hertz as a "binder" or "validator".

## Usage
* default validate tag: `vd`
* validate rule:  [bytedance/go-tagexpr](https://github.com/bytedance/go-tagexpr/tree/master/validator)

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/binding/go_tagexpr"
)

type Test struct {
	Q int `query:"q" vd:"$>100"`
}

func main() {
	binder := go_tagexpr.NewBinder()
	// If you need to configure the validator, you can use the following
	//go_tagexpr.MustRegTypeUnmarshal
	h := server.New(server.WithCustomBinder(binder))
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
	hc.Get("http://127.0.0.1:8888/ping?q=99")
	time.Sleep(1 * time.Second)
}


```