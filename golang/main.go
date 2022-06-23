package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	//"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	// "github.com/joho/godotenv"
)

type Bird struct {
	Species     string
	Description string
}

var connectionString = ("root:donQuiote2@tcp(127.0.0.1:3306)/test_api")

var db, err = sql.Open("mysql", connectionString)

type Article struct {
	Id       string `json:"Id,omitempty"`
	Title    string `json:"Title,omitempty"`
	Descript string `json:"Descript,omitempty"`
	Content  string `json:"Content,omitempty"`
}

// type Article struct{
// 	Id string
// 	Title string
// 	Descript string
// 	Content string
// }

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Homepage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	results, err := db.Query("Select * from Articles")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var article Article
		err = results.Scan(&article.Id, &article.Title, &article.Descript, &article.Content)
		if err != nil {
			panic(err.Error())
		}
		Articles = append(Articles, article)
	}
	fmt.Println("Endpoint hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keys := vars["id"]
	fmt.Println("hit endpoint: SingleArticle")
	fmt.Println("Key:" + keys)
	fmt.Println(fmt.Sprintf("Select * from Articles where Id= %s", &keys))
	results, err := db.Query(fmt.Sprintf("Select * from Articles where Id= %s", keys))
	var article Article
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		err = results.Scan(&article.Id, &article.Title, &article.Descript, &article.Content)
	}
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode(article)
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)

	fmt.Fprintf(w, "%+v", string(reqBody))
	var toStringBody = string(reqBody)
	var setList Article
	fmt.Println(toStringBody)
	if(toStringBody[0]=='\''){
		reqBody=reqBody[1:len(reqBody)-1]
	}
	fmt.Printf("%s\n",string(reqBody))
	var j_err=json.Unmarshal([]byte(reqBody), &setList)
	if j_err!=nil{
		
		panic(j_err.Error())
	}
	fmt.Printf("%+v\n", (setList))
	fmt.Printf("insert into Articles values(%s,%s,%s,%s)\n", setList.Id, setList.Title, setList.Descript, setList.Content)

	var _,err = db.Exec(fmt.Sprintf("insert into Articles Values(%s,\"%s\",\"%s\",\"%s\");", setList.Id, setList.Title, setList.Descript, setList.Content))
	// defer insert.Close()
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode(reqBody)

}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]
	var delStatement = fmt.Sprintf("delete from Articles where Id= %s ", id)
	fmt.Println(delStatement)
	var del, err = db.Query(delStatement)
	defer del.Close()
	if err != nil {
		panic(err.Error())
	}
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

func updateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var reqBody, _ = ioutil.ReadAll(r.Body)
	if(string(reqBody)[0]=='\''){
		reqBody=reqBody[1:len(reqBody)-1]
	}
	var article Article
	json.Unmarshal(reqBody, &article)
	var upStatement = fmt.Sprintf("update Articles set Descript=\"%s\", Title=\"%s\", Content=\"%s\" where Id= %s ",article.Descript,article.Title,article.Content, id)
	var _,err=db.Exec(upStatement)
	if err!=nil{
		panic(err.Error())
	}
}

func handleRequests() {
	// http.HandleFunc("/",homePage)
	// http.HandleFunc("/articles/{id}",returnAllArticles)
	// log.Fatal(http.ListenAndServe(":10000",nil))

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles/", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/articles/{id}", returnSingleArticle)
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println(connectionString)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	handleRequests()
}
