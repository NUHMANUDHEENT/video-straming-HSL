package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	// Serve static files (HLS files)
	http.Handle("/hls/", http.StripPrefix("/hls/", http.FileServer(http.Dir("./hls"))))

	// Serve index.html with dynamic video source
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		videoType := strings.TrimPrefix(r.URL.Path, "/")
		if videoType == "" {
			videoType = "nature" // Default video
		}
		http.ServeFile(w, r, "index.html")
	})

	log.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
