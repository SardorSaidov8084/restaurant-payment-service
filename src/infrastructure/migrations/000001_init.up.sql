

-- eaters
CREATE TABLE IF NOT EXISTS merchant_settings (
	entity_id varchar(36) PRIMARY KEY,
    cashbox_id VARCHAR(36),
    enabled boolean,
	created_at timestamp,
	updated_at timestamp
);

-- idx_eater_phone_number
CREATE INDEX IF NOT EXISTS idx_eater_phone_number ON eater.eaters(phone_number);

type MerchantSetting struct {
	EntityID  string    `json:"entity_id" gorm:"primaryKey"`
	CashboxID string    `json:"cashbox_id"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}