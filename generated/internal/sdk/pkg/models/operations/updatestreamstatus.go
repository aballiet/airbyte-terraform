// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type UpdateStreamStatusResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successfully updated stream status.
	StreamStatusRead *shared.StreamStatusRead
}