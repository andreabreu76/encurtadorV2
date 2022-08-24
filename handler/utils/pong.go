package utils

import "net/http"

func Pong(w http.ResponseWriter, r *http.Request) {

	//JSON response shortener
	ResponseWithJSON(w, []byte(`[{"message":"shortener..."}]`), http.StatusOK)

}
