package dao

import (
	"context"
	"time"

	"github.com/goharbor/harbor/src/common"
	"github.com/goharbor/harbor/src/common/utils"
	"github.com/goharbor/harbor/src/lib/errors"
	"github.com/goharbor/harbor/src/lib/orm"
	"github.com/goharbor/harbor/src/lib/q"
	"github.com/goharbor/harbor/src/pkg/usergroup/model"
)

type dao struct {
}

// ErrGroupNameDup ...
var ErrGroupNameDup = errors.ConflictError(nil).WithMessage("duplicated user group name")

func (d *dao) Add(ctx context.Context, userGroup model.UserGroup) (int, error) {
	query := q.New(q.KeyWords{"GroupName": userGroup.GroupName, "GroupType": common.HTTPGroupType})
	userGroupList, err := d.Query(ctx, query)
	if err != nil {
		return 0, ErrGroupNameDup
	}
	if len(userGroupList) > 0 {
		return 0, ErrGroupNameDup
	}
	o, err := orm.FromContext(ctx)
	if err != nil {
		return 0, err
	}
	sql := "insert into user_group (group_name, group_type, ldap_group_dn, creation_time, update_time) values (?, ?, ?,? ?) RETURNING id"
	var id int
	now := time.Now()

	err = o.Raw(sql, userGroup.GroupName, userGroup.GroupType, utils.TrimLower(userGroup.LdapGroupDN), now, now).QueryRow(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Query - Query User Group
func (d *dao) Query(ctx context.Context, query *q.Query) ([]*model.UserGroup, error) {
	query = q.MustClone(query)
	qs, err := orm.QuerySetter(ctx, &model.UserGroup{}, query)
	if err != nil {
		return nil, err
	}
	var usergroups []*model.UserGroup
	if _, err := qs.All(&usergroups); err != nil {
		return nil, err
	}
	return usergroups, nil
}

func (d *dao) onBoardCommonUserGroup(ctx context.Context, g *model.UserGroup, keyAttribute string, combinedKeyAttributes ...string)

func (d *dao) Count(ctx context.Context, query *q.Query) (int64, error) {
	query = q.MustClone(query)
	qs, err := orm.QuerySetterForCount(ctx, &model.UserGroup{}, query)
	if err != nil {
		return 0, err
	}
	return qs.Count()
}