package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Fungsi Log yang berguna sebagai middleware
func Auth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password Tidak boleh Kosong"))
			return
		}
		if uname == "admin" && pwd == "password" {
			next(w, r, ps)
			return
		}
		if uname == "editor" && pwd == "secret" {
			next(w, r, ps)
			return
		}
		if uname == "trainer" && pwd == "rahasia" {
			next(w, r, ps)
			return
		}
		w.Write([]byte("Username atau Password tidak sesuai"))
		return
	}
}
