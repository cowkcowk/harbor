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

package user

import (
	"context"

	commonmodels "github.com/goharbor/harbor/src/common/models"
	"github.com/goharbor/harbor/src/lib/q"

	"github.com/goharbor/harbor/src/pkg/user"
	"github.com/goharbor/harbor/src/pkg/user/models"
)

var ()

// Controller provides functions to support API/middleware for user management and query
type Controller interface {
	// SetSysAdmin ...
	SetSysAdmin(ctx context.Context, id int, adminFlag bool) error
	// VerifyPassword ...
	VerifyPassword(ctx context.Context, usernameOrEmail string, password string) (bool, error)
	// UpdatePassword ...
	UpdatePassword(ctx context.Context, id int, password string) error
	// List ...
	List(ctx context.Context, query *q.Query, options ...models.Option) ([]*commonmodels.User, error)
	// Create ...
	Create(ctx context.Context, u *commonmodels.User) (int, error)
	// Count ...
	Count(ctx context.Context, query *q.Query) (int64, error)
	// Get ...
	Get(ctx context.Context, id int, opt *Option) (*commonmodels.User, error)
	// GetByName gets the user model by username, it only supports getting the basic and does not support opt
	GetByName(ctx context.Context, username string) (*commonmodels.User, error)
	// GetBySubIss gets the user model by subject and issuer, the result will contain the basic user model and does not support opt
	GetBySubIss(ctx context.Context, sub, iss string) (*commonmodels.User, error)
	// Delete ...
	Delete(ctx context.Context, id int) error
	// UpdateProfile update the profile based on the ID and data in the model in parm, only a subset of attributes in the model
	// will be update, see the implementation of manager.
	UpdateProfile(ctx context.Context, u *commonmodels.User, cols ...string) error
	// SetCliSecret sets the OIDC CLI secret for a user
	SetCliSecret(ctx context.Context, id int, secret string) error
	// UpdateOIDCMeta updates the OIDC metadata of a user, if the cols are not provided, by default the field of token and secret will be updated
	UpdateOIDCMeta(ctx context.Context, ou *commonmodels.OIDCUser, cols ...string) error
	// OnboardOIDCUser inserts the record for basic user info and the oidc metadata
	// if the onboard process is successful the input parm of user model will be populated with user id
	OnboardOIDCUser(ctx context.Context, u *commonmodels.User) error
}

// Option  option for getting User info
type Option struct {
	WithOIDCInfo bool
}

type controller struct {
	mgr user.Manager
	odicMetaMgr
}
