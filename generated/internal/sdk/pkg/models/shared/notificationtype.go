// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type NotificationType string

const (
	NotificationTypeSlack      NotificationType = "slack"
	NotificationTypeCustomerio NotificationType = "customerio"
)

func (e NotificationType) ToPointer() *NotificationType {
	return &e
}

func (e *NotificationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "slack":
		fallthrough
	case "customerio":
		*e = NotificationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NotificationType: %v", v)
	}
}
