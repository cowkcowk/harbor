package api

import (
	"github.com/goharbor/harbor/src/common/api"
	"github.com/goharbor/harbor/src/common/rbac"
	"github.com/goharbor/harbor/src/common/utils"
)

const (
	yamlFileContentType = "application/x-yaml"
)

type BaseController struct {
	api.BaseApi
}

func (b *BaseController) RequireAuthenticated() bool {
	return true
}

func (b *BaseController) HasProjectPermission(projectIDOrName interface{}, action rbac.Action, subresource ...rbac.Resource) (bool, error) {
	_, _, err := utils.ParseProjectIDOrName(projectIDOrName)
	if err != nil {
		return false, err
	}

	project, err := b.ProjectCtl
}