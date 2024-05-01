package links

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func Init() {
	db, err := sql.Open("sqlite", "links.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS links (id INTEGER PRIMARY KEY, nick TEXT, url TEXT, readonly BOOLEAN DEFAULT FALSE, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)")
	if err != nil {
		panic(err)
	}
}

func getDb() *sql.DB {
	db, err := sql.Open("sqlite", "links.sqlite")
	if err != nil {
		panic(err)
	}
	return db
}
