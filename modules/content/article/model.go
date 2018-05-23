package article

import (
	"time"

	"github.com/GuilhermeVendramini/golang-cms-mysql/config"
)

// Create a new Article
func Create(item Article) (Article, error) {
	stmtIns, err := config.DB.Prepare("INSERT INTO articles (Title, Teaser, Body, Image, Tags, Author, Url, Changed, Created) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	defer stmtIns.Close()
	if err != nil {
		return item, err
	}
	stmtIns.Exec(item.Title, item.Teaser, item.Body, item.Image, item.Tags, item.Author, item.URL, item.Changed, item.Created)
	return item, nil
}

// GetbyID return one article by ID
func GetbyID(ID string) (Article, error) {
	item := Article{}
	rows, err := config.DB.Query("SELECT Id, Title, Teaser, Body, Image, Tags, Author, Url, Changed, Created FROM articles WHERE Id = ?", ID)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&item.ID, &item.Title, &item.Teaser, &item.Body, &item.Image, &item.Tags, &item.Author, &item.URL, &item.Changed, &item.Created)
	}
	return item, err
}

// GetbyURL return one article by URL
func GetbyURL(URL string) (Article, error) {
	item := Article{}
	rows, err := config.DB.Query("SELECT Id, Title, Teaser, Body, Image, Tags, Author, Url, Changed, Created FROM articles WHERE Url = ?", URL)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&item.ID, &item.Title, &item.Teaser, &item.Body, &item.Image, &item.Tags, &item.Author, &item.URL, &item.Changed, &item.Created)
	}
	return item, err
}

// Remove article
func Remove(ID string) error {
	stmtDel, err := config.DB.Prepare("DELETE FROM articles WHERE Id = ?")
	defer stmtDel.Close()
	if err != nil {
		return err
	}
	stmtDel.Exec(ID)
	return nil
}

// Update article
func Update(item Article, ID string) (Article, error) {
	stmtUp, err := config.DB.Prepare("UPDATE articles SET Title = ?, Teaser = ?, Body = ?, Image = ?, Tags = ?, Author = ?, URL = ?, Changed = ? WHERE Id = ?")
	defer stmtUp.Close()
	if err != nil {
		return item, err
	}
	stmtUp.Exec(item.Title, item.Teaser, item.Body, item.Image, item.Tags, item.Author, item.URL, time.Now(), ID)
	return item, nil
}

// GetAll return all articles
func GetAll() ([]Article, error) {
	item := Article{}
	items := []Article{}
	rows, _ := config.DB.Query("SELECT Id, Title, Teaser, Body, Image, Tags, Author, Url, Changed, Created FROM articles")
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&item.ID, &item.Title, &item.Teaser, &item.Body, &item.Image, &item.Tags, &item.Author, &item.URL, &item.Changed, &item.Created)
		items = append(items, item)
	}
	return items, nil
}

// GetSkip return skipping articles
func GetSkip(s int) ([]Article, error) {
	item := Article{}
	items := []Article{}
	rows, _ := config.DB.Query("SELECT Id, Title, Teaser, Body, Image, Tags, Author, Url, Changed, Created FROM articles ORDER BY Id LIMIT 10 OFFSET ?", s)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&item.ID, &item.Title, &item.Teaser, &item.Body, &item.Image, &item.Tags, &item.Author, &item.URL, &item.Changed, &item.Created)
		items = append(items, item)
	}
	return items, nil
}

// GetNext article
func GetNext(s int) (Article, error) {
	item := Article{}
	rows, err := config.DB.Query("SELECT Id, Title, Teaser, Body, Image, Tags, Author, Url, Changed, Created FROM articles ORDER BY Id LIMIT 1 OFFSET ?", s)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&item.ID, &item.Title, &item.Teaser, &item.Body, &item.Image, &item.Tags, &item.Author, &item.URL, &item.Changed, &item.Created)
	}
	return item, err
}
