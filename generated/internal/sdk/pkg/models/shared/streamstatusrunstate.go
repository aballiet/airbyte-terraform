// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// StreamStatusRunState - Values:
//   - `PENDING` - The stream operation has been selected to run
//   - `RUNNING` - The stream operation is running
//   - `COMPLETE` - The stream operation ran successfully to completion
//   - `INCOMPLETE` - The stream operation has terminated in an incomplete state.
//     See StreamStatusIncompleteRunCause for more details.
type StreamStatusRunState string

const (
	StreamStatusRunStatePending    StreamStatusRunState = "PENDING"
	StreamStatusRunStateRunning    StreamStatusRunState = "RUNNING"
	StreamStatusRunStateComplete   StreamStatusRunState = "COMPLETE"
	StreamStatusRunStateIncomplete StreamStatusRunState = "INCOMPLETE"
)

func (e StreamStatusRunState) ToPointer() *StreamStatusRunState {
	return &e
}

func (e *StreamStatusRunState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PENDING":
		fallthrough
	case "RUNNING":
		fallthrough
	case "COMPLETE":
		fallthrough
	case "INCOMPLETE":
		*e = StreamStatusRunState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StreamStatusRunState: %v", v)
	}
}