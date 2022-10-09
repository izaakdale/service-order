package dao

import (
	"time"
)

type OrderRecord struct {
	OrderId    string    `dynamodbav:"PK" json:"order_id,omitempty"`
	RecordType string    `dynamodbav:"SK" json:"record_type,omitempty"`
	Items      []string  `json:"items,omitempty"`
	Created    time.Time `json:"created,omitempty"`
	Status     string    `json:"status,omitempty"`
}
