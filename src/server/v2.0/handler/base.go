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

package handler

// TODO move this file out of v2.0 folder as this is common for all versions of API

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/goharbor/harbor/src/common/rbac"
	"github.com/goharbor/harbor/src/common/security"
	"github.com/goharbor/harbor/src/common/utils"
	"github.com/goharbor/harbor/src/controller/project"

	"github.com/goharbor/harbor/src/lib/errors"
	lib_http "github.com/goharbor/harbor/src/lib/http"
	"github.com/goharbor/harbor/src/lib/log"
)

var (
	baseProjectCtl = project.Ctl
)

// BaseAPI base API handler
type BaseAPI struct{}

// Prepare default prepare for operation
func (*BaseAPI) Prepare(_ context.Context, _ string, _ interface{}) middleware.Responder {
	return nil
}

// SendError returns response for the err
func (*BaseAPI) SendError(_ context.Context, err error) middleware.Responder {
	return NewErrResponder(err)
}

// GetSecurityContext from the provided context
func (*BaseAPI) GetSecurityContext(ctx context.Context) (security.Context, error) {
	sc, ok := security.FromContext(ctx)
	if !ok {
		return nil, errors.UnauthorizedError(errors.New("security context not found"))
	}
	return sc, nil
}

// HasPermission returns true when the request has action permission on resource
func (b *BaseAPI) HasPermission(ctx context.Context, action rbac.Action, resource rbac.Resource) bool {
	s, err := b.GetSecurityContext(ctx)
	if err != nil {
		log.Warningf("security context not found")
		return false
	}
	return s.Can(ctx, action, resource)
}

func (b *BaseAPI) HasProjectPermission(ctx context.Context, projectIDOrName interface{}, action rbac.Action, subresource ...rbac.Resource) bool {
	projectID, projectName, err := utils.ParseProjectIDOrName(projectIDOrName)
	if err != nil {
		return false
	}

	if projectName != "" {
		p, err := base
	}
}

var _ middleware.Responder = &ErrResponder{}

// ErrResponder error responder
type ErrResponder struct {
	err error
}

// WriteResponse ...
func (r *ErrResponder) WriteResponse(rw http.ResponseWriter, _ runtime.Producer) {
	lib_http.SendError(rw, r.err)
}

// NewErrResponder returns responder for err
func NewErrResponder(err error) *ErrResponder {
	return &ErrResponder{err: err}
}
