package utils

import "github.com/astaxie/beego/context"

func SetErrorStatus(ctx *context.Context, errCode int, errMsg string) {
	if ctx == nil {
		return
	}

	ctx.Output.SetStatus(errCode)
	ctx.Output.Body([]byte(errMsg))
}
