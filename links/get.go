package links

import (
	"fmt"
	"linklist/structs"
)

func Get() []structs.Link {
	db := getDb()
	defer db.Close()

	rows, err := db.Query("SELECT id, nick, url, readonly FROM links")
	if err != nil {
		fmt.Println(err)
		return []structs.Link{}
	}
	defer rows.Close()

	links := []structs.Link{}
	for rows.Next() {
		link := structs.Link{}
		err := rows.Scan(&link.Id, &link.Nick, &link.Url, &link.Readonly)
		if err != nil {
			fmt.Println(err)
			return []structs.Link{}
		}
		links = append(links, link)
	}

	return links
}

func GetUrlByNick(nick string) (string, error) {
	db := getDb()
	defer db.Close()

	row := db.QueryRow("SELECT url FROM links WHERE nick = ?", nick)

	var url string
	err := row.Scan(&url)
	if err != nil {
		return "", err
	}

	return url, nil
}

func GetUrlById(id string) (string, error) {
	db := getDb()
	defer db.Close()

	row := db.QueryRow("SELECT url FROM links WHERE id = ?", id)

	var url string
	err := row.Scan(&url)
	if err != nil {
		return "", err
	}

	return url, nil
}
