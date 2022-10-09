package dao

import "fmt"

const (
	OrderPrefixPK = "ORDER#"
	ItemSK        = "ITEMS"
	StatusSK      = "STATUS"
)

func OrderPK(id string) string {
	return fmt.Sprintf("%s%s", OrderPrefixPK, id)
}
