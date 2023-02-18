package app

import (
	"encoding/json"
	"log"

	"github.com/izaakdale/lib/listener"
	"github.com/izaakdale/lib/publisher"
	"github.com/izaakdale/service-event-order/internal/datastore"
	storeNotif "github.com/izaakdale/service-event-order/pkg/notifications"
	createdNotif "github.com/izaakdale/service-event/pkg/notifications"
)

func orderEventListener(m listener.Message) error {
	log.Printf("recieved messagge %+v\n", m)
	var n createdNotif.OrderCreatedPayload
	err := json.Unmarshal([]byte(m.Message), &n)
	if err != nil {
		return err
	}

	err = datastore.Insert(n)
	if err != nil {
		return err
	}

	// publish order store
	os := storeNotif.OrderStoredPayload{
		OrderID: n.OrderID,
	}

	// return whatever is inputted regardless of error
	oBytes, err := json.Marshal(os)
	if err != nil {
		return err
	}

	_, err = publisher.Publish(string(oBytes))
	if err != nil {
		return err
	}

	return nil
}
