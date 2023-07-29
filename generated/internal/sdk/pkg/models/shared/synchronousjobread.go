// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SynchronousJobRead struct {
	// only present if a config id was provided.
	ConfigID                      *string        `json:"configId,omitempty"`
	ConfigType                    JobConfigType  `json:"configType"`
	ConnectorConfigurationUpdated *bool          `json:"connectorConfigurationUpdated,omitempty"`
	CreatedAt                     int64          `json:"createdAt"`
	EndedAt                       int64          `json:"endedAt"`
	FailureReason                 *FailureReason `json:"failureReason,omitempty"`
	ID                            string         `json:"id"`
	Logs                          *LogRead       `json:"logs,omitempty"`
	Succeeded                     bool           `json:"succeeded"`
}
