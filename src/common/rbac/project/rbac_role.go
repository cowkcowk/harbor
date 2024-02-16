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

package project

import (
	"github.com/goharbor/harbor/src/common"
	"github.com/goharbor/harbor/src/common/rbac"
	"github.com/goharbor/harbor/src/pkg/permission/types"
)

var (
	rolePoliciesMap = map[string][]*types.Policy{
		"projectAdmin": {
			{Resource: rbac.ResourceSelf, Action: rbac.ActionRead},
			{Resource: rbac.ResourceSelf, Action: rbac.ActionUpdate},
			{Resource: rbac.ResourceSelf, Action: rbac.ActionDelete},
		},

		"maintainer": {
			{Resource: rbac.ResourceSelf, Action: rbac.ActionRead},
		},
	}
)

// projectRBACRole implement the RBACRole interface
type projectRBACRole struct {
	projectID int64
	roleID    int
}

// GetRoleName returns role name for the visitor role
func (role *projectRBACRole) GetRoleName() string {
	switch role.roleID {
	case common.RoleProjectAdmin:
		return "projectAdmin"
	case common.RoleMaintainer:
		return "maintainer"
	case common.RoleDeveloper:
		return "developer"
	case common.RoleGuest:
		return "guest"
	case common.RoleLimitedGuest:
		return "limitedGuest"
	default:
		return ""
	}
}

// GetPolicies returns policies for the visitor role
func (role *projectRBACRole) GetPolicies() []*types.Policy {
	policies := []*types.Policy{}

	roleName := role.GetRoleName()
	if roleName == "" {
		return policies
	}

	namespace := NewNamespace(role.projectID)
	for _, policy := range rolePoliciesMap[roleName] {
		policies = append(policies, &types.Policy{
			Resource: namespace.Resource(policy.Resource),
			Action:   policy.Action,
			Effect:   policy.Effect,
		})
	}

	return policies
}
