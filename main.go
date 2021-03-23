package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"google.golang.org/api/iterator"
)

type article struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Date     string `json:"date"`
	Category string `json:"category"`
	Type     string `json:"type"`
	Image    string `json:"image"`
	Creator  string `json:"creator"`
	Content  string `json:"content"`
	Active   string `json:"active"`
	Link     string `json:"link"`
}

// The main function is the entry of the server. It is where the HTTP handler
// that serves the UI is defined and where the server is started.
//
// Note that because main.go and app.go are built for different architectures,
// this main() function is not in conflict with the one in
// app.go.
func main() {

	http.Handle("/", &app.Handler{
		Name:        "SpVgg Hambach",
		Title:       "SpVgg Hambach",
		Description: "Webiste of SpVgg Hambach",
		Icon: app.Icon{
			Default:    "/web/images/hambach_logo_192.png", // Specify default favicon.
			Large:      "/web/images/hambach_logo_512.png",
			AppleTouch: "/web/images/hambach_logo_192.png", // Specify icon on IOS devices.
		},
		Styles: []string{
			"https://cdn.jsdelivr.net/npm/bulma@0.9.1/css/bulma.min.css",
			"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.2/css/all.min.css",
			"/web/css/main.css",
		},
		ThemeColor: "#008000",
	})
	http.HandleFunc("/api/v1/articles", getArticles)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		contentJSON, err := json.Marshal(loadArticles())
		if err != nil {
			log.Fatal(err)
		}

		w.Write(contentJSON)
		return
	}

	id := keys[0]

	articlesJSON, err := json.Marshal(loadArticle(id))
	if err != nil {
		log.Fatal(err)
	}

	w.Write(articlesJSON)
}

func loadArticles() []article {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	var articles []article

	iter := client.Collection("articles").OrderBy("id", firestore.Desc).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		var a article
		doc.DataTo(&a)

		//if a.Category == "Allgemein" && a.Type == "article" {
		articles = append(articles, a)
		//}
	}

	return articles
}

func loadArticle(id string) article {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	result, err := client.Collection("articles").Doc(id).Get(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	var a article
	result.DataTo(&a)

	return a
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
