package index

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// lUser := users.GetLoggedUser(r)
	// vars := make(map[string]interface{})
	// vars["LoggedUser"] = lUser
	// err := config.TPL.ExecuteTemplate(w, "index.html", vars)
	// HandleError(w, err)
	fmt.Println("Teste")
}

// HandleError return Status Internal Server Error
// func HandleError(w http.ResponseWriter, err error) {
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		log.Fatalln(err)
// 	}
// }
