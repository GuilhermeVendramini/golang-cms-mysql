package article

import (
	"github.com/GuilhermeVendramini/golang-cms-mysql/config"
	"gopkg.in/mgo.v2"
)

// Articles collection
var Articles *mgo.Collection

func init() {
	Articles = config.DB.C("articles")
}
