package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	listenAddr := "localhost"
	addr := listenAddr + `:` + "8080"

	http.HandleFunc("/watch", stream)
	log.Printf("starting server at %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func stream(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query().Get("v")
	if v == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "use format /watch?v=...")
		return
	}

	err := downloadVideoAndExtractAudio(v, w)
	if err != nil {
		log.Printf("stream error: %v", err)
		fmt.Fprintf(w, "stream error: %v", err)
		return
	}
}
func downloadVideoAndExtractAudio(id string, out io.Writer) error {
	url := fmt.Sprintf("https://youtube.com/watch?v=" + id)
	//ytdl := exec.Command("youtube-dl", id)

	ytdl := exec.Command("youtube-dl", url, "-o", "-")
	ytdl.Stdout = out
	ytdl.Stderr = os.Stderr // show progress

	err := ytdl.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("stream finished")

	return err
}
