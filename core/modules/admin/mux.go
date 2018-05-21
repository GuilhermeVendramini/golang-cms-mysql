package admin

import "github.com/GuilhermeVendramini/golang-cms-mysql/core/utils"

// Mux admin
func Mux() {
	utils.Mux.GET("/admin", Admin)
	utils.Mux.GET("/admin/content", Content)
}
