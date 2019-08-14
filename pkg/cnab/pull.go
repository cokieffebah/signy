package cnab

import (
	"context"
	"fmt"

	"github.com/deislabs/cnab-go/bundle"
	"github.com/docker/cnab-to-oci/remotes"
	"github.com/docker/distribution/reference"
)

// Pull pulls a bundle from an OCI registry
func Pull(ref string) (*bundle.Bundle, error) {
	n, err := reference.ParseNormalizedNamed(ref)
	if err != nil {
		return nil, err
	}

	b, relocationMap, err := remotes.Pull(context.Background(), n, createResolver(nil))
	fmt.Println("\nRelocation map", relocationMap)
	if err != nil {
		return nil, err
	}
	return b, nil
}