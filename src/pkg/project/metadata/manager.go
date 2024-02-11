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

	"github.com/goharbor/harbor/src/pkg/project/metadata/dao"
	"github.com/goharbor/harbor/src/pkg/project/metadata/models"
)

// Manager defines the operations that a project metadata manager should implement
type Manager interface {
	// Add metadatas for project specified by projectID
	Add(ctx context.Context, projectID int64, meta map[string]string) error

	// Delete metadatas whose keys are specified in parameter meta, if it is absent, delete all
	Delete(ctx context.Context, projectID int64, meta ...string) error

	// Update metadatas
	Update(ctx context.Context, projectID int64, meta map[string]string) error

	// Get metadatas whose keys are specified in parameter meta, if it is absent, get all
	Get(ctx context.Context, projectID int64, meta ...string) (map[string]string, error)

	// List metadata according to the name and value
	List(ctx context.Context, name, value string) ([]*models.ProjectMetadata, error)
}

type manager struct {
	dao dao.DAO
}

// Add metadatas for project specified by projectID
func (m *manager) Add(ctx context.Context, projectID int64, meta map[string]string) error {
	h := func(ctx context.Context) error {

	}
}
