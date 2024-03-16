package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"gitlab.com/arnaud-web/neli-forms/config"
	"gitlab.com/arnaud-web/neli-forms/password"
	"gitlab.com/arnaud-web/neli-forms/video"
)

var logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

func main() {
	logger.Println("Running server on port", *config.Port)

	r := chi.NewRouter()
	r.Get("/set-password/{code}", password.Set)
	r.Get("/confirmation", password.Confirmation)
	r.Get("/{userId:[0-9]+}/video/{content}", video.Get)

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "./__resources__/static/assets")
	fileServer(r, "/static", http.Dir(filesDir))

	filesDir = filepath.Join(workDir, "./__resources__/neli")
	fileServer(r, "/neli", http.Dir(filesDir))

	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", *config.Port), r))
}

func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}

	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
