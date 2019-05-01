package pedafytig

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", apiHomeH)
}

func apiHomeH(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	fmt.Fprint(w, `{"status":"running"}`)
}
