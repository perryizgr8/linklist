package links

func Add(nick, url string) error {
	db := getDb()
	defer db.Close()

	_, err := db.Exec("INSERT INTO links (nick, url) VALUES (?, ?)", nick, url)
	return err
}
