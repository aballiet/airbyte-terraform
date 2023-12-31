// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type CreateConnectorBuilderProjectResponse struct {
	// Successful operation
	ConnectorBuilderProjectIDWithWorkspaceID *shared.ConnectorBuilderProjectIDWithWorkspaceID
	ContentType                              string
	StatusCode                               int
	RawResponse                              *http.Response
}
