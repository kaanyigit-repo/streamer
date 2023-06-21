package main

import (
	"fmt"
	"github.com/kkdai/youtube/v2"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var Addr string

func main() {
	greet()
}

func greet() {
	fmt.Println("Input the download link: ")

	fmt.Scanln(&Addr)

	if strings.Contains(Addr, "youtube") {

		streamServer()
	}
}

func streamServer() {
	http.HandleFunc("/stream", streamBot)
	log.Printf("starting server at %s", "localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func streamBot(w http.ResponseWriter, r *http.Request) {

	u, err := url.Parse(Addr)
	if err != nil {
		log.Fatal(err)
	}
	videoID := u.Query().Get("v")

	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		log.Println("Error1:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	formats := video.Formats.WithAudioChannels()
	fmt.Println("check1")

	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		log.Println("Error2:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(w, stream)

	fmt.Println("check2")
	file, err := os.Create(videoID + ".mp4")
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println("check3")
	defer file.Close()

	_, err = io.Copy(file, stream)

	if err != nil {
		log.Println("Error3:", err)
	}
	fmt.Println("check4")

	/*
		netErr, ok := err.(net.Error)
		if ok && netErr.Timeout() {
			log.Println("Client disconnected")
			return
		}
	*/

	if err != nil {
		log.Println("Error4:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Println("Streaming completed")
}
