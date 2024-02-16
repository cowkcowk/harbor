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

package registry

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/goharbor/harbor/src/lib"

	"github.com/distribution/distribution"
	"github.com/goharbor/harbor/src/lib/log"

	"github.com/goharbor/harbor/src/pkg/registry/interceptor"
)

var ()

const (

	// DefaultHTTPClientTimeout is the default timeout for registry http client.
	DefaultHTTPClientTimeout = 30 * time.Minute
)

var (
	// registryHTTPClientTimeout is the timeout for registry http client.
	registryHTTPClientTimeout time.Duration
)

func init() {
	registryHTTPClientTimeout = DefaultHTTPClientTimeout
	// override it if read from environment variable, in minutes
	if env := os.Getenv("REGISTRY_HTTP_CLIENT_TIMEOUT"); len(env) > 0 {
		timeout, err := strconv.ParseInt(env, 10, 64)
		if err != nil {
			log.Errorf("Failed to parse REGISTRY_HTTP_CLIENT_TIMEOUT: %v, use default value: %v", err, DefaultHTTPClientTimeout)
		} else {
			if timeout > 0 {
				registryHTTPClientTimeout = time.Duration(timeout) * time.Minute
			}
		}
	}
}

// Client defines the methods that a registry client should implements
type Client interface {
	// Ping the base API endpoint "/v2/"
	Ping() (err error)
	// Catalog the repositories
	Catalog() (repositories []string, err error)
	// ListTags lists the tags under the specified repository
	ListTags(repository string) (tags []string, err error)
	// ManifestExist checks the existence of the manifest
	ManifestExist(repository, reference string) (exist bool, desc *distribution.Descriptor, err error)
	// PullManifest pulls the specified manifest
	PullManifest(repository, reference string, acceptedMediaTypes ...string) (manifest distribution.Manifest, digest string, err error)
	// PushManifest pushes the specified manifest
	PushManifest(repository, reference, mediaType string, payload []byte) (digest string, err error)
	// DeleteManifest deletes the specified manifest. The "reference" can be "tag" or "digest"
	DeleteManifest(repository, reference string) (err error)
	// BlobExist checks the existence of the specified blob
	BlobExist(repository, digest string) (exist bool, err error)
	// PullBlob pulls the specified blob. The caller must close the returned "blob"
	PullBlob(repository, digest string) (size int64, blob io.ReadCloser, err error)
	// PullBlobChunk pulls the specified blob, but by chunked
	PullBlobChunk(repository, digest string, blobSize, start, end int64) (size int64, blob io.ReadCloser, err error)
	// PushBlob pushes the specified blob
	PushBlob(repository, digest string, size int64, blob io.Reader) error
	// PushBlobChunk pushes the specified blob, but by chunked
	PushBlobChunk(repository, digest string, blobSize int64, chunk io.Reader, start, end int64, location string) (nextUploadLocation string, endRange int64, err error)
	// MountBlob mounts the blob from the source repository
	MountBlob(srcRepository, digest, dstRepository string) (err error)
	// DeleteBlob deletes the specified blob
	DeleteBlob(repository, digest string) (err error)
	// Copy the artifact from source repository to the destination. The "override"
	// is used to specify whether the destination artifact will be overridden if
	// its name is same with source but digest isn't
	Copy(srcRepository, srcReference, dstRepository, dstReference string, override bool) (err error)
	// Do send generic HTTP requests to the target registry service
	Do(req *http.Request) (*http.Response, error)
}

func NewClient(url, username, password string, insecure bool, interceptors ...interceptor.Interceptor) Client {

}

type client struct {
	url          string
	authorizer   lib.Authorizer
	interceptors []interceptor.Interceptor
	client       *http.Client
}

func (c *client) Ping() error {
	req, err := http.NewRequest(http.MethodGet)
}

func buildPingURL(endpoint string) string {
	return fmt.Sprintf("%s/v2/", endpoint)
}
