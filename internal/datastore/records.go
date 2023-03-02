package datastore

type OrderRecord struct {
	PK string `dynamodbav:"PK" json:"-"`
	SK string `dynamodbav:"SK" json:"-"`

	Type     string       `dynamodbav:"type" json:"type"`
	Metadata MetaRecord   `dynamodbav:"metadata" json:"metadata"`
	Ticket   TicketRecord `dynamodbav:"ticket" json:"ticket"`
}

type TicketRecord struct {
	TicketID  string `json:"ticket_id"`
	FirstName string `json:"first_name"`
	Surname   string `json:"surname"`
	Dob       string `json:"dob"`
	EventID   int64  `json:"event_id"`
	Type      string `json:"type"`
	QRPath    string `json:"qr_path"`
	QRScanned bool   `json:"qr_scanned"`
}
type MetaRecord struct {
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phone_number"`
	Address     Address `json:"address"`
}
type Address struct {
	NameOrNumber string `json:"name_or_number"`
	Street       string `json:"street"`
	City         string `json:"city"`
	Postcode     string `json:"postcode"`
}

type OrderTickets struct {
	Email   string
	Tickets []*TicketRecord
}
