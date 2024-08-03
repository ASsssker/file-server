package server

import (
	"log"
	"net/http"
	"path/filepath"
)

type Handler http.Handler

func NewHandler(path string) Handler {

	path, err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("Не удалось получить абсолютный путь к каталогу: %s: %s", path, err.Error())
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir(path))
	http.Handle("/", fs)

	return mux
}