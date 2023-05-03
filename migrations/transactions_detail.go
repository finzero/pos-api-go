package migrations

const TransactionsDetail = `
CREATE TABLE IF NOT EXISTS transactions_detail (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  transaction_no VARCHAR(20) NOT NULL,
  product_id INTEGER NOT NULL,
	quantity INTEGER NOT NULL,
  price NUMERIC NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
)
`
