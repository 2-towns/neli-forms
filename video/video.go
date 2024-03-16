package video

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"gitlab.com/arnaud-web/neli-forms/config"
	"gitlab.com/arnaud-web/neli-forms/page"
)

type share struct {
	Message        string
	ExpirationDate time.Time `json:"expirationDate"`
	Path           string
	Name           string
}

func Get(w http.ResponseWriter, r *http.Request) {
	// Ignore error because uid is necessarily an int due to match param url
	// uid, _ := strconv.ParseInt(chi.URLParam(r, "userId"), 10, 64)
	// Ignore error because uid is necessarily an int due to match param url
	c := chi.URLParam(r, "content")

	b, err := base64.URLEncoding.DecodeString(c)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf := bytes.Buffer{}
	buf.Write(b)

	s := share{}
	if err := gob.NewDecoder(&buf).Decode(&s); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if time.Now().After(s.ExpirationDate) {
		page.Display(w, r, "expired.gohtml","")
		return
	}

	url := fmt.Sprintf("%s/%s/%s.mp4", *config.VideosURL, s.Path, s.Path)

	content := struct {
		URL     string
		Name    string
		Message string
	}{
		url,
		s.Name,
		s.Message,
	}

	page.Display(w, r, "video.gohtml", content)
}
