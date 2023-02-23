package main

import (
	"DailyEmoji/emoji"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		seed, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal("Given argument is not a int64 number")
			return
		}
		emoji.SetSeed(seed)
	} else {
		log.Println("Seed not given. Using standard seed.")
	}

	emoji.ShuffleEmojiList()
	indexTmpl, err := template.ParseFiles("site/templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	aboutTmpl, err := template.ParseFiles("site/templates/about.html")
	if err != nil {
		log.Fatal(err)
	}

	timeZone, _ := time.Now().Zone()

	http.HandleFunc("/api/emoji", func(writer http.ResponseWriter, request *http.Request) {
		err := json.NewEncoder(writer).Encode(map[string]string{"emoji": *emoji.GetCurrentEmoji()})
		if err != nil {
			log.Println("Error while loading /api/emoji", err)
			return
		}
	})

	http.HandleFunc("/about", func(writer http.ResponseWriter, request *http.Request) {
		err = aboutTmpl.Execute(writer, timeZone)
		if err != nil {
			log.Println("Error while loading about.html:", err)
		}
	})

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/" {
			http.NotFound(writer, request)
			return
		}
		err = indexTmpl.Execute(writer, emoji.GetCurrentEmoji())
		if err != nil {
			log.Println("Error while loading index.html:", err)
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
		return
	}
}
