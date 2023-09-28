/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
