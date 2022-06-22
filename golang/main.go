package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct{
	Id string `json:"Id"`
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w,"Welcome to the Homepage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter,r *http.Request){
	fmt.Println("Endpoint hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
	keys:=vars["id"]
	fmt.Println("hit endpoint: SingleArticle")
	fmt.Fprint(w,"Key:"+keys)
	for _,article:=range Articles{
		if(article.Id==keys){
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter,r *http.Request){
	reqBody,_:=ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody,&article)
	fmt.Printf("%+v",article)
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(reqBody)
}

func deleteArticle(w http.ResponseWriter,r *http.Request){
	 vars := mux.Vars(r)
    // we will need to extract the `id` of the article we
    // wish to delete
    id := vars["id"]

    // we then need to loop through all our articles
    for index, article := range Articles {
        // if our id path parameter matches one of our
        // articles
        if article.Id == id {
            // updates our Articles array to remove the 
            // article
            Articles = append(Articles[:index], Articles[index+1:]...)
        }
    }
}

func updateArticle(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
	id:=vars["id"]
	var reqBody,_=ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody,&article)
	for index, article := range Articles{
		if(article.Id==id){
			tmp:=append(Articles[:index],article)
			Articles=append(tmp,Articles[index+1:]... )
		}
	}
}

func handleRequests(){
	// http.HandleFunc("/",homePage)
	// http.HandleFunc("/articles/{id}",returnAllArticles)
	// log.Fatal(http.ListenAndServe(":10000",nil))

	myRouter:=mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/",homePage)
	myRouter.HandleFunc("/articles/",returnAllArticles)
	myRouter.HandleFunc("/article",createNewArticle).Methods("POST")
	myRouter.HandleFunc("/articles/{id}",returnSingleArticle)
	myRouter.HandleFunc("/article/{id}",deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}",updateArticle).Methods("PUT")
	log.Fatal(http.ListenAndServe(":10000",myRouter))
}

func main (){

	Articles=[]Article{
		Article{
			Id: "1",
			Title: "Hello",
			Desc: "Insert default",
			Content: "Insert Default",
		},
		Article{
			Id: "2",
			Title: "Hello_2",
			Desc: "Insert default",
			Content: "Insert Default",
		},
	}

	handleRequests()
}