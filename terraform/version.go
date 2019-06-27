package terraform

import (
	"github.com/hashicorp/terraform-plugin-sdk/version"
)

// TODO: update providers to use the version package directly
func VersionString() string {
	return version.String()
}
