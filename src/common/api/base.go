package api

import (
	"context"
	"encoding/json"

	"github.com/beego/beego/v2/server/web"

	"github.com/cowkcowk/harbor/src/lib/log"
)

const (
	defaultPageSize int64 = 500
	maxPageSize     int64 = 500

	// APIVersion is the current core api version
	APIVersion = "v2.0"
)

type BaseApi struct {
	web.Controller
}

func (b *BaseApi) Context() context.Context {
	return b.Ctx.Request.Context()
}

func (b *BaseApi) GetStringFromPath(key string) string {
	return b.Ctx.Input.Param(key)
}

func (b *BaseApi) Render() error {
	return nil
}

func (b *BaseApi) DecodeJSONReq(v interface{}) error {
	err := json.Unmarshal(b.Ctx.Input.CopyBody(1<<35), v)
	if err != nil {
		log.Errorf()
	}
}
