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

package metadata

import (
	"context"
)

// Controller defines the operations that a project metadata controller should implement
type Controller interface {
	// Add metadatas for project specified by projectID
	Add(ctx context.Context, projectID int64, meta map[string]string) error

	// Delete metadatas whose keys are specified in parameter meta, if it is absent, delete all
	Delete(ctx context.Context, projectID int64, meta ...string) error

	// Update metadatas
	Update(ctx context.Context, projectID int64, meta map[string]string) error

	// Get metadatas whose keys are specified in parameter meta, if it is absent, get all
	Get(ctx context.Context, projectID int64, meta ...string) (map[string]string, error)
}

func NewController() Controller {
	return
}

type controller struct {
	mgr meta
}
