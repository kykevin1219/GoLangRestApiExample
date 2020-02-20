package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/pat"
)

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"Content"`
}

var Articles []Article

//getArticle returns an article in form of JSON
func getArticle(w http.ResponseWriter, r *http.Request) {
	queryId := r.URL.Query().Get(":id")
	fmt.Println(queryId)
	fmt.Println("End point Hit : getArticle no.", queryId)
	for _, article := range Articles {
		if article.Id == queryId {
			json.NewEncoder(w).Encode(article)
		}
	}
}

//returnAllArticles gives all articles in main function
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("End point Hit : returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

//writeArticle just returns input data in the form of JSON
func writeArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("End point Hit : writeArticle")
	body := json.NewDecoder(r.Body)
	var article Article
	body.Decode(&article)
	json.NewEncoder(w).Encode(article)
}

func main() {
	Articles = []Article{
		Article{Id: "1", Title: "title1", Desc: "desc1", Content: "This is contetn1"},
		Article{Id: "2", Title: "title2", Desc: "desc2", Content: "This is contetn2"},
	}
	m := pat.New()
	m.Get("/articles/{id}", getArticle)
	m.Get("/articles", returnAllArticles)
	m.Post("/articles", writeArticle)
	http.Handle("/", m)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
