package links

import "time"

func Edit(id, nick, url string) error {
	db := getDb()
	defer db.Close()

	updated_at := time.Now().Format("2006-01-02 15:04:05")
	_, err := db.Exec("UPDATE links SET nick = ?, url = ?, updated_at = ? WHERE id = ? AND readonly = 0", nick, url, updated_at, id)
	return err
}
