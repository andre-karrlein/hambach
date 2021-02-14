package main

import (
	"context"
	"html/template"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
)

// Article struct
type Article struct {
	ID       int
	Title    string
	Date     string
	Category string
	Type     string
	Image    string
	Creator  string
	Content  template.HTML
	Active   string
	Link     string
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../html/Layout.html", "../html/Main.html", "../html/MainNavbar.html", "../html/Carousel.html")
	if err != nil {
		log.Fatal(err.Error())
	}

	data := loadArticles("Allgemein")

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, data)
}

func handleSportheim(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../html/Layout.html", "../html/SportheimCarousel.html", "../html/MainNavbar.html")
	if err != nil {
		log.Fatal(err.Error())
	}

	article := loadArticle("1")

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, article)
}

func handleSoccer(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../html/Layout.html", "../html/IndexFussball.html", "../html/FussballNavbar.html", "../html/Carousel.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	w.WriteHeader(http.StatusOK)
	articles := loadArticles("Fussball")

	tmpl.Execute(w, articles)
}

func handleKorbball(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../html/Layout.html", "../html/Main.html", "../html/KorbballNavbar.html", "../html/Carousel.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	w.WriteHeader(http.StatusOK)
	articles := loadArticles("Korbball")

	tmpl.Execute(w, articles)
}

func handleArticle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	article := loadArticle(id)

	navbar := "../html/MainNavbar.html"
	if article.Category == "Fussball" {
		navbar = "../html/FussballNavbar.html"
	}
	if article.Category == "Korbball" {
		navbar = "../html/KorbballNavbar.html"
	}

	tmpl, err := template.ParseFiles("../html/Layout.html", navbar, "../html/Article.html")
	if err != nil {
		log.Fatal(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, article)
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
}

func loadArticles(category string) []Article {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	var articles []Article

	iter := client.Collection("articles").OrderBy("id", firestore.Desc).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		var a Article
		doc.DataTo(&a)

		if a.Category == category && a.Type == "article" {
			articles = append(articles, a)
		}
	}

	return articles
}

func loadArticle(id string) Article {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	result, err := client.Collection("articles").Doc(id).Get(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	var a Article
	result.DataTo(&a)

	return a
}

func main() {
	r := mux.NewRouter()

	staticDir := "/public/"

	r.HandleFunc("/", handleHome).Methods(http.MethodGet)
	r.HandleFunc("/sportheim", handleSportheim).Methods(http.MethodGet)
	r.HandleFunc("/fussball", handleSoccer).Methods(http.MethodGet)
	r.HandleFunc("/korbball", handleKorbball).Methods(http.MethodGet)
	r.HandleFunc("/article/{id}", handleArticle).Methods(http.MethodGet)
	r.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir(".."+staticDir))))
	r.
		PathPrefix("/images/").
		Handler(http.StripPrefix("/images/", http.FileServer(http.Dir(".."+staticDir+"images/"))))
	r.
		PathPrefix("/download-data/").
		Handler(http.StripPrefix("/download-data/", http.FileServer(http.Dir(".."+staticDir+"download-data/"))))
	

	log.Fatal(http.ListenAndServe(":8080", r))
}

func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := "hambach" // os.Getenv("PROJECT_ID")

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client
}
