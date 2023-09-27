

-- merchant_settings
CREATE TABLE IF NOT EXISTS merchant_settings (
	entity_id varchar(36) PRIMARY KEY,
    cashbox_id VARCHAR(36),
    enabled boolean,
	created_at timestamp,
	updated_at timestamp
);

-- receipts
CREATE TABLE IF NOT EXISTS receipts (
	id varchar(36) PRIMARY KEY,
    order_id VARCHAR(36),
    card_id VARCHAR(36),
	amount int,
	data jsonb,
	created_at timestamp,
	updated_at timestamp
);
-- idx_order_receipt
CREATE INDEX IF NOT EXISTS idx_order_receipt ON receipts(order_id);

-- transactions
CREATE TABLE IF NOT EXISTS transactions (
	id varchar(36) PRIMARY KEY,
    restaurant_id varchar(36),
    currency varchar(20),
	status varchar(50),
	create_time int,
	pay_time int,
	cancel_time int,
	card_id varchar(36),
	amount int,
	account jsonb,
	created_at timestamp,
	updated_at timestamp
);

-- idx_restaurant_tx
CREATE INDEX IF NOT EXISTS idx_restaurant_tx ON transactions(restaurant_id);

-- payment_cards
CREATE TABLE IF NOT EXISTS payment_cards (
	id varchar(36) PRIMARY KEY,
    card_token VARCHAR(250),
	created_at timestamp
);

