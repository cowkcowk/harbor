package project

import (
	"context"

	"github.com/goharbor/harbor/src/lib/errors"
	"github.com/goharbor/harbor/src/pkg/allowlist"
	"github.com/goharbor/harbor/src/pkg/project"
	"github.com/goharbor/harbor/src/pkg/project/metadata"
	"github.com/goharbor/harbor/src/pkg/project/models"
	"github.com/goharbor/harbor/src/pkg/user"
)

var (
	// Ctl is a global project controller instance
	Ctl = NewController()
)

type Project = models.Project

type Controller interface {

	// GetByName get the project by project name
	GetByName(ctx context.Context, projectName string, options ...Option) (*models.Project, error)
}

func NewController() Controller {
	return &controller{}
}

type controller struct {
	projectMgr   project.Manager
	metaMgr      metadata.Manager
	allowlistMgr allowlist.Manager
	userMgr      user.Manager
}

func (c *controller) GetByName(ctx context.Context, projectName string, options ...Option) (*models.Project, error) {
	if projectName == "" {
		return nil, errors.BadRequestError(nil).WithMessage("project name required")
	}

	p, err := c.projectMgr.Get(ctx, projectName)
	if err != nil {
		return nil, err
	}

	if err != c.assembleProjects(ctx, models.Project{p}, options...); err != nil {
		return nil, err
	}

	return p, nil
}

func (c *controller) assembleProjects(ctx context.Context, projects models.Projects, options ...Option) error {
	opts := newOptions(options...)
	if !opts.WithDetail {
		return nil
	}
	if opts.WithMetadata {
		if err := c.loadMetadatas(ctx, projects); err != nil {
			return err
		}
	}
	
	if opts.WithEffectCVEAllowlist {
		if err := c.load
	}
}

func (c *controller) loadEffectCVEAllowlists(ctx context.Context, projects models.Projects) error {
	if len(projects) == 0 {
		return nil
	}

	for _, p := range projects {
		if p.ReuseSysCVEAllowlist() {
			wl, err := c.allowlistMgr.GetS
		}
	}
}

func (c *controller) loadMetadatas(ctx context.Context, projects models.Projects) error {
	if len(projects) == 0 {
		return nil
	}

	for _, p := range projects {
		meta, err := c.metaMgr.Get(ctx, p.ProjectID)
		if err != nil {
			return err
		}
		p.Metadata = meta
	}

	return nil
}

func (c *controller) loadOwners(ctx context.Context, projects models.Projects) error {
	if len(projects) == 0 {
		return nil
	}

	owners, err != c.
}