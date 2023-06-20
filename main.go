package main

import (
	"fmt"
	"github.com/kkdai/youtube/v2"
	"io"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {
	var addr string
	greet(addr)

}
func greet(web string) {

	fmt.Println("Input the download link: ")
	fmt.Scanln(&web)
	if strings.Contains(web, "youtube") {
		ExampleClient(web)
		//startserver()
		//openBrowser(web)
	}

}
func ExampleClient(web string) {
	u, err := url.Parse(web)
	if err != nil {
		log.Fatal(err)
	}
	videoID := u.Query().Get("v")
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels()
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}

	file, err := os.Create(videoID + ".mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}

}

/*
func startserver() {
	listenAddr := "localhost"
	addr := listenAddr + `:` + "8080"
	http.HandleFunc("/watch", stream)
	log.Printf("starting server at %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
func openBrowser(web string) {
	u, err := url.Parse(web)
	if err != nil {
		log.Fatal(err)
	}
	query := u.Query()
	videoID := query.Get("v")
	url := fmt.Sprintf("localhost:8080/watch?v=" + videoID)

	// Command to open the link in the default browser
	cmd := exec.Command("xdg-open", url) // For Linux

	err2 := cmd.Start()
	if err2 != nil {
		log.Fatal(err)
	}
}

func stream(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query().Get("v")
	if v == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "use format /watch?v=...")
		return
	}

	err := dl(v, w)
	if err != nil {
		log.Printf("stream error: %v", err)
		fmt.Fprintf(w, "stream error: %v", err)
		return
	}
}
func dl(id string, out io.Writer) error {
	url := fmt.Sprintf("https://youtube.com/watch?v=" + id)

	ytdl := exec.Command("youtube-dl", url, "-o-")
	ytdl.Stdout = out
	ytdl.Stderr = os.Stderr

	err := ytdl.Run()
	log.Printf("stream finished")
	return err
}


*/
