package models

type AddressBookEntry struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
}
