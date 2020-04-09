package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("Server is running")

	//half assed attempt to connect to mongodb atlas.
	connectToDB()

	//Silly little static handler thingy to handle static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//Look at these shitty ass "routes" LMAOOO
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/index.html", indexHandler)
	http.HandleFunc("/about.html", aboutHandler)
	http.HandleFunc("/map.html", mapHandler)
	http.HandleFunc("/contact.html", contactHandler)
	http.HandleFunc("/login.html", logInHandler)
	http.HandleFunc("/signup.html", signUpHandler)

	http.ListenAndServe(getPort(), nil)
}

func getPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "5000"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

func connectToDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://darienmiller88:nintendowiiu000@cluster0-wf6d0.mongodb.net/test?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}

	client.Connect(ctx)
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	myTemplate, _ := template.ParseFiles("templates/index.html")
	myTemplate.Execute(writer, request)
}

func aboutHandler(writer http.ResponseWriter, request *http.Request) {
	myTemplate, _ := template.ParseFiles("templates/about.html")
	myTemplate.Execute(writer, request)
}

func contactHandler(writer http.ResponseWriter, request *http.Request) {
	myTemplate, _ := template.ParseFiles("templates/contact.html")
	myTemplate.Execute(writer, request)
}

func logInHandler(writer http.ResponseWriter, request *http.Request) {
	myTemplate, _ := template.ParseFiles("templates/login.html")
	myTemplate.Execute(writer, request)
}

func signUpHandler(writer http.ResponseWriter, request *http.Request) {
	myTemplate, _ := template.ParseFiles("templates/signup.html")
	myTemplate.Execute(writer, request)
}

func mapHandler(writer http.ResponseWriter, request *http.Request) {
	myTemplate, _ := template.ParseFiles("templates/map.html")
	myTemplate.Execute(writer, request)
}
