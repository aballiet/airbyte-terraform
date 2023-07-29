// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type SetWorkflowInAttemptResponse struct {
	ContentType string
	// Successful Operation
	InternalOperationResult *shared.InternalOperationResult
	StatusCode              int
	RawResponse             *http.Response
}
