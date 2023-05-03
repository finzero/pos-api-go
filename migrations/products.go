package migrations

const Products = `
CREATE TABLE IF NOT EXISTS products (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
	item_code VARCHAR(10) NULL,
  name VARCHAR(64) NOT NULL,
  price NUMERIC NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
)
`
