package output

import (
	"github.com/zzztttkkk/sha"
	_ "glass/restapi/internal/init"
)

type M map[string]interface{}

type Msg struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func OK(ctx *sha.RequestCtx, v interface{}) { ctx.WriteJSON(Msg{Data: v}) }
