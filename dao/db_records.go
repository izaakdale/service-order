package dao

import (
	"time"
)

type OrderRecord struct {
	OrderId    string `dynamodbav:"PK"`
	RecordType string `dynamodbav:"SK"`
	Items      []string
	Created    time.Time
}
