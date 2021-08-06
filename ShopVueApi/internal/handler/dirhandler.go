package handler

import "net/http"

//
func dirhandler(patern, filedir string) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		handler := http.StripPrefix(patern, http.FileServer(http.Dir(filedir)))
		handler.ServeHTTP(w, req)

	}
}
