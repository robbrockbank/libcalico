package api

import (
	"errors"

	"github.com/mohae/utilitybelt/deepcopy"
	"github.com/projectcalico/libcalico/lib/api/unversioned"
	"github.com/projectcalico/libcalico/lib/api/v1"
	"fmt"
)

type ResourceManager struct {
	ResourceHelper map[unversioned.TypeMetadata]ResourceHelper
}

type ResourceHelper struct {
	EmptyResource unversioned.Resource
}

func (rm *ResourceManager) registerResource(r unversioned.Resource) {
	rm.ResourceHelper[r.TypeMetadata] = ResourceHelper{r}
}

func CreateResourceManager() *ResourceManager {
	rm := &ResourceManager{}
	rm.registerResource(v1.ResourceTier(&v1.TierMetadata{}, &v1.TierSpec{}))
	rm.registerResource(v1.ResourcePolicy(&v1.PolicyMetadata{}, &v1.PolicySpec{}))
	rm.registerResource(v1.ResourceProfile(&v1.ProfileMetadata{}, &v1.ProfileSpec{}))
	rm.registerResource(v1.ResourceHostEndpoint(&v1.HostEndpointMetadata{}, &v1.HostEndpointSpec{}))
	return rm
}

func (rm *ResourceManager) NewResource(tm unversioned.TypeMetadata) (*unversioned.Resource, error) {
	var new unversioned.Resource
	rh, ok := rm.ResourceHelper[tm]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Unknown resource type (%s) and version (%s)", tm.Kind, tm.Version))
	}
	new = deepcopy.Iface(rh.EmptyResource).(unversioned.Resource)
	return &new, nil
}
