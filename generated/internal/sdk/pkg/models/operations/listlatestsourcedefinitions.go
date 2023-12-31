// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type ListLatestSourceDefinitionsResponse struct {
	ContentType string
	// Successful operation
	SourceDefinitionReadList *shared.SourceDefinitionReadList
	StatusCode               int
	RawResponse              *http.Response
}
