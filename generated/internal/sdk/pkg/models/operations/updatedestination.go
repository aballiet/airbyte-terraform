// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type UpdateDestinationResponse struct {
	ContentType string
	// Successful operation
	DestinationRead *shared.DestinationRead
	// Input failed validation
	InvalidInputExceptionInfo *shared.InvalidInputExceptionInfo
	StatusCode                int
	RawResponse               *http.Response
}