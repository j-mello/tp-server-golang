package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func clockHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		currentTime := time.Now()
		fmt.Fprintf(w, "Il est %v", currentTime.Format("15h04"))
	default:
		fmt.Fprintf(w, "Requête non prise en charge")
	}
}

func addHandler(w http.ResponseWriter, req *http.Request) {
	//fmt.Fprintf(w, "Je suis l'ajout")
	switch req.Method {
	case http.MethodPost:
		if err := req.ParseForm(); err != nil {
			fmt.Println("Il y a une erreur")
			fmt.Fprintln(w, "Il y a une erreur")
			return
		}
		entry := req.PostForm.Get("entry")
		author := req.PostForm.Get("author")

		file, err := os.OpenFile("list.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
		defer file.Close()

		if err != nil {
			panic(err)
		}

		_, err = file.WriteString(req.PostForm.Get("author") + " : " + req.PostForm.Get("entry") + "\n")

		if err != nil {
			panic(err)
		}

		fmt.Fprintf(w, author+" : "+entry+"\n")
	default:
		fmt.Fprintf(w, "Requête non prise en charge")
	}

}

func entriesHandler(w http.ResponseWriter, req *http.Request) {
	//fmt.Fprintf(w, "Je suis la liste")
	if req.Method == http.MethodGet {
		file, err := os.OpenFile("list.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
		defer file.Close()

		if err != nil {
			fmt.Println(err)
		}

		data, err := ioutil.ReadFile("list.txt")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Fprintf(w, string(data))
	}
}

func main() {

	http.HandleFunc("/", clockHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/entries", entriesHandler)
	http.ListenAndServe(":4567", nil)
}
