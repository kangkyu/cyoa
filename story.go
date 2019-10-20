package cyoa

import (
	"log"
	"net/http"
)

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v", h.s["intro"])
}

type Story map[string]Chapter

type Chapter struct {
	Title string `json:"title"`
	Paragraphs []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Chapter string `json:"arc"`
}
