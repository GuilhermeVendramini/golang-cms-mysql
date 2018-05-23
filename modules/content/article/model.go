package article

import (
	"github.com/GuilhermeVendramini/golang-cms-mysql/config"
)

// Create a new Article
func Create(item Article) (Article, error) {
	// err := Articles.Insert(item)
	// if err != nil {
	// 	return item, errors.New("internal server error" + err.Error())
	// }
	// return item, nil
	stmtIns, err := config.DB.Prepare("INSERT INTO articles (Title, Teaser, Body, Image, Tags, Author, Url, Changed, Created) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	defer stmtIns.Close()
	if err != nil {
		panic(err.Error())
	}
	stmtIns.Exec(item.Title, item.Teaser, item.Body, item.Image, item.Tags, item.Author, item.URL, item.Changed, item.Created)
	return item, nil
}

// // GetbyID return one article by ID
// func GetbyID(ID string) (Article, error) {
// 	item := Article{}
// 	err := Articles.Find(bson.M{"_id": bson.ObjectIdHex(ID)}).One(&item)
// 	return item, err
// }

// GetbyURL return one article by URL
func GetbyURL(URL string) (Article, error) {
	// item := Article{}
	// err := Articles.Find(bson.M{"url": URL}).One(&item)
	// return item, err
	item := Article{}
	rows, err := config.DB.Query("SELECT Id, Title, Teaser, Body, Image, Tags, Author, Url, Changed, Created FROM articles WHERE Url = ?", URL)
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&item.ID, &item.Title, &item.Teaser, &item.Body, &item.Image, &item.Tags, &item.Author, &item.URL, &item.Changed, &item.Created)
	}
	return item, err
}

// // Remove article
// func Remove(ID string) error {
// 	err := Articles.Remove(bson.M{"_id": bson.ObjectIdHex(ID)})
// 	if err != nil {
// 		return errors.New("500 internal server error")
// 	}
// 	return nil
// }

// Update article
func Update(item Article, ID string) (Article, error) {
	// err := Articles.Update(bson.M{"_id": bson.ObjectIdHex(ID)}, &item)
	// if err != nil {
	// 	return item, err
	// }
	return item, nil
}

// GetAll return all articles
func GetAll() ([]Article, error) {
	// items := []Article{}
	// err := Articles.Find(bson.M{}).Sort("-_id").All(&items)
	// if err != nil {
	// 	return nil, err
	// }
	// return items, nil
	item := Article{}
	items := []Article{}
	rows, _ := config.DB.Query("SELECT Id, Title, Teaser, Body, Image, Tags, Author, Url, Changed, Created FROM articles")
	defer rows.Close()

	for rows.Next() {
		//var created time.Time
		rows.Scan(&item.ID, &item.Title, &item.Teaser, &item.Body, &item.Image, &item.Tags, &item.Author, &item.URL, &item.Changed, &item.Created)
		//fmt.Println(created)
		items = append(items, item)
	}
	return items, nil
}

// // GetSkip return skipping articles
// func GetSkip(s int) ([]Article, error) {
// 	items := []Article{}
// 	err := Articles.Find(bson.M{}).Sort("-_id").Skip(s).Limit(10).All(&items)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return items, nil
// }

// // GetNext article
// func GetNext(s int) (Article, error) {
// 	item := Article{}
// 	err := Articles.Find(bson.M{}).Sort("-_id").Skip(s).Limit(1).One(&item)
// 	if err != nil {
// 		return item, err
// 	}
// 	return item, nil
// }
