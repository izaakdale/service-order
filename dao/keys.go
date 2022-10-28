package dao

import "fmt"

const (
	OrderPrefixPK = "ORDER#"
	ItemSK        = "ITEMS"
	StatusSK      = "STATUS"
	AddressSK     = "ADDRESS"
)

func OrderPK(id string) string {
	return fmt.Sprintf("%s%s", OrderPrefixPK, id)
}
