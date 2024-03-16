package password

import (
	"net/http"

	"github.com/go-chi/chi"
	"gitlab.com/arnaud-web/neli-forms/config"
	"gitlab.com/arnaud-web/neli-forms/page"
)

func Set(w http.ResponseWriter, r *http.Request) {
	c := chi.URLParam(r, "code")

	res := struct {
		URL  string
		Code string
	}{
		*config.ApiURL,
		c,
	}

	page.Display(w, r, "reset.gohtml", res)
}

func Confirmation(w http.ResponseWriter, r *http.Request) {
	var i interface{}
	page.Display(w, r, "confirmation.gohtml", i)
}
