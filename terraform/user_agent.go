package terraform

import (
	"github.com/hashicorp/terraform-plugin-sdk/httpclient"
)

// Generate a UserAgent string
//
// Deprecated: Use httpclient.UserAgentString if you are setting your
// own User-Agent header.
func UserAgentString() string {
	return httpclient.UserAgentString()
}
