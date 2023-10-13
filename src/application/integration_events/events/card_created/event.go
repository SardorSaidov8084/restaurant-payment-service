package cardcreated

import "time"

type Card struct {
	ID        string    `json:"id"`
	EaterID   string    `json:"eater_id"`
	Number    string    `json:"number"`
	CardToken string    `json:"card_token"`
	IsVerifed bool      `json:"is_verifed"`
	CreatedAt time.Time `json:"created_at"`
}
type Event struct {
	Card *Card `json:"card"`
}
func (e Event) Topic() string {
	return "card.added"
}