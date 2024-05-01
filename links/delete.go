package links

func Delete(id string) error {
	db := getDb()
	defer db.Close()

	_, err := db.Exec("DELETE FROM links WHERE id = ? AND readonly = 0", id)
	return err
}
