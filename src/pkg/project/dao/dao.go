package dao

import (
	"time"

	"github.com/goharbor/harbor/src/lib/orm"
)

func init() {
	orm.RegisterModel()
}

// Project holds the details of a project.
type Project struct {
	ProjectID    int64                  `orm:"pk;auto;column(project_id)" json:"project_id"`
	OwnerID      int                    `orm:"column(owner_id)" json:"owner_id"`
	Name         string                 `orm:"column(name)" json:"name" sort:"default"`
	CreationTime time.Time              `orm:"column(creation_time);auto_now_add" json:"creation_time"`
	UpdateTime   time.Time              `orm:"column(update_time);auto_now" json:"update_time"`
	Deleted      bool                   `orm:"column(deleted)" json:"deleted"`
	OwnerName    string                 `orm:"-" json:"owner_name"`
	Role         int                    `orm:"-" json:"current_user_role_id"`
	RoleList     []int                  `orm:"-" json:"current_user_role_ids"`
	RepoCount    int64                  `orm:"-" json:"repo_count"`
	Metadata     map[string]string      `orm:"-" json:"metadata"`
	CVEAllowlist allowlist.CVEAllowlist `orm:"-" json:"cve_allowlist"`
	RegistryID   int64                  `orm:"column(registry_id)" json:"registry_id"`
}