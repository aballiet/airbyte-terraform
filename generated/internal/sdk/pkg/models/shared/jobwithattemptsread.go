// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type JobWithAttemptsRead struct {
	Attempts []AttemptRead `json:"attempts,omitempty"`
	Job      *JobRead      `json:"job,omitempty"`
}
