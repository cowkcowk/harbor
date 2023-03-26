package artifact

import (
	"github.com/goharbor/harbor/src/controller/tag"
	"github.com/goharbor/harbor/src/pkg/label/model"
)

type Artifact struct {
	Tags         []*tag.Tag              `json:"tags"`
	AddtionLinks map[string]*AdditionLink `json:"addtional_links"`
	Labels []*model.Label
	Accessories []access
}

// AdditionLink is a link via that the addition can be fetched
type AdditionLink struct {
	HREF     string `json:"href"`
	Absolute bool   `json:"absolute"` // specify the href is an absolute URL or not
}

// Option is used to specify the properties returned when listing/getting artifacts
type Option struct {
	WithTag       bool
	TagOption     *tag.Option // only works when WithTag is set to true
	WithLabel     bool
	WithAccessory bool
}
