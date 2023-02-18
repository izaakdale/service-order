package datastore

import "fmt"

const (
	orderPK  = "order"
	metaSK   = "meta"
	ticketSK = "ticket"
)

func createKey(prefix, id string) string {
	return fmt.Sprintf("%s#%s", prefix, id)
}
