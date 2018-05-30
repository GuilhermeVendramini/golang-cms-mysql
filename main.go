package main

import (
	"net/http"

	"github.com/GuilhermeVendramini/golang-cms-mysql/core/index"
	"github.com/GuilhermeVendramini/golang-cms-mysql/core/modules/admin"
	"github.com/GuilhermeVendramini/golang-cms-mysql/core/modules/users"
	"github.com/GuilhermeVendramini/golang-cms-mysql/core/utils"
	"github.com/GuilhermeVendramini/golang-cms-mysql/modules/contact"
	"github.com/GuilhermeVendramini/golang-cms-mysql/modules/content/article"
)

func main() {
	// Index Mux
	index.Mux()

	// Users Mux
	users.Mux()

	// Admin Mux
	admin.Mux()

	// Contact Mux
	contact.Mux()

	// Content Mux
	article.Mux()

	// Uncomment the line below to generate a demo user
	//demo.User()

	// Server Listen
	http.ListenAndServe(":8080", utils.Mux)
}
