// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ConnectionStatus - Active means that data is flowing through the connection. Inactive means it is not. Deprecated means the connection is off and cannot be re-activated. the schema field describes the elements of the schema that will be synced.
type ConnectionStatus string

const (
	ConnectionStatusActive     ConnectionStatus = "active"
	ConnectionStatusInactive   ConnectionStatus = "inactive"
	ConnectionStatusDeprecated ConnectionStatus = "deprecated"
)

func (e ConnectionStatus) ToPointer() *ConnectionStatus {
	return &e
}

func (e *ConnectionStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "active":
		fallthrough
	case "inactive":
		fallthrough
	case "deprecated":
		*e = ConnectionStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionStatus: %v", v)
	}
}
