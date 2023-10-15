package carddeleted

type Event struct {
	CardID string `json:"card_id"`
}

func (e Event) Topic() string {
	return "card.deleted"
}