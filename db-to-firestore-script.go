package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"strconv"

	"cloud.google.com/go/firestore"
	_ "github.com/go-sql-driver/mysql"
)

type article struct {
	ID       int    `firestore:"id,omitempty"`
	Title    string `firestore:"title,omitempty"`
	Date     string `firestore:"date,omitempty"`
	Category string `firestore:"category,omitempty"`
	Type     string `firestore:"type,omitempty"`
	Image    string `firestore:"image,omitempty"`
	Creator  string `firestore:"creator,omitempty"`
	Content  string `firestore:"content,omitempty"`
	Active   string `firestore:"active,omitempty"`
}

func main() {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	articles := loadArticles()
	for _, article := range articles {
		_, err := client.Collection("articles").Doc(strconv.Itoa(article.ID)).Set(ctx, article)
		if err != nil {
			// Handle any errors in an appropriate way, such as returning them.
			log.Printf("An error has occurred: %s", err)
		}
	}
}

func loadArticles() []article {
	user := os.Getenv("DBUSERNAME")
	password := os.Getenv("DBPASSWORD")
	host := os.Getenv("DBHOST")
	dbname := os.Getenv("DBNAME")

	connection := user + ":" + password + "@(" + host + ":3306)/" + dbname + "?parseTime=true&charset=utf8"

	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT contentId, title, creator, date, article, category, type, titleImage FROM content`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var articles []article
	for rows.Next() {
		var a article
		err = rows.Scan(&a.ID, &a.Title, &a.Creator, &a.Date, &a.Content, &a.Category, &a.Type, &a.Image)
		if err != nil {
			log.Fatal(err)
		}
		if a.Image == "" {
			a.Image = "/public/images/hambach_logo.png"
		}
		a.Active = ""
		articles = append(articles, a)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return articles
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
