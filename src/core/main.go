// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/beego/beego/v2/server/web"

	"github.com/goharbor/harbor/src/core/middlewares"

	"github.com/goharbor/harbor/src/lib/config"

	"github.com/goharbor/harbor/src/lib/log"
)

const (
	adminUserID = 1
)

func main() {

	web.BConfig.WebConfig.Session.SessionOn = true
	web.BConfig.WebConfig.Session.SessionName = config.SessionCookieName

	log.Info("initializing cache ...")

	web.RunWithMiddleWares("", middlewares.MiddleWares()...)
}
