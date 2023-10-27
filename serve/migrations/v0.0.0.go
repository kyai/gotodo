package migrations

import "gotodo/serve/models"

func init() {
	models.DB.Exec(`
CREATE TABLE IF NOT EXISTS project (
	id INTEGER PRIMARY KEY
);
	`)
}
