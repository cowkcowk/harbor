package api

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/beego/beego/v2/server/web"

	lib_http "github.com/goharbor/harbor/src/lib/http"
	"github.com/goharbor/harbor/src/lib/log"
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

func (b *BaseApi) GetInt64FromPath(key string) (int64, error) {
	value := b.Ctx.Input.Param(key)
	return strconv.ParseInt(value, 10, 64)
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

func (b *BaseApi) SendError(err error) {
	lib_http.SendError(b.Ctx.ResponseWriter, err)
}
