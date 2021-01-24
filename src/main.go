package main

import (
	"context"
	"html/template"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
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
	Editor   string
	Content  string
	Active   string
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../html/Layout.html", "../html/Main.html", "../html/MainNavbar.html", "../html/Carousel.html")
	if err != nil {
		log.Fatal(err.Error())
	}

	data := []Article{
		{
			ID:       1,
			Title:    "TESTING",
			Date:     "01.01.2021",
			Category: "Allgemein",
			Type:     "article",
			Image:    "/public/images/hambach_logo.png",
			Creator:  "Andre Karrlein",
			Editor:   "Andre Karrlein",
			Content:  "<p>HAMBACH HTML DATA<p>",
			Active:   "active",
		},
		{
			ID:       2,
			Title:    "TESTING 2",
			Date:     "01.02.2021",
			Category: "Allgemein",
			Type:     "article",
			Image:    "/public/images/hambach_logo.png",
			Creator:  "Andre Karrlein",
			Editor:   "Andre Karrlein",
			Content:  "<p>HAMBACH HTML DATA<p>",
			Active:   "",
		},
		{
			ID:       3,
			Title:    "TESTING 3",
			Date:     "01.03.2021",
			Category: "Allgemein",
			Type:     "article",
			Image:    "/public/images/hambach_logo.png",
			Creator:  "Andre Karrlein",
			Editor:   "Andre Karrlein",
			Content:  "<p>HAMBACH HTML DATA<p>",
			Active:   "",
		},
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, data)
}

func handleSportheim(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../html/Layout.html", "../html/SportheimCarousel.html", "../html/MainNavbar.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, Article{
		ID:       4,
		Title:    "Sporthiem",
		Date:     "01.03.2021",
		Category: "Allgemein",
		Type:     "article",
		Image:    "/public/images/hambach_logo.png",
		Creator:  "Andre Karrlein",
		Editor:   "Andre Karrlein",
		Content:  "<p>HAMBACH HTML DATA<p>",
		Active:   "",
	})
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
}

/*func loadData() []data {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	var expirences []experience

	iter := client.Collection("resume").Doc("JrEvSIoWiSgTgXQRIC6I").Collection("experience").OrderBy("order", firestore.Asc).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		var e experience
		doc.DataTo(&e)

		expirences = append(expirences, e)
	}

	return expirences
}*/

func main() {
	r := mux.NewRouter()

	staticDir := "/public/"

	r.HandleFunc("/", handleHome).Methods(http.MethodGet)
	r.HandleFunc("/sportheim", handleSportheim).Methods(http.MethodGet)
	r.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir(".."+staticDir))))
		//Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	log.Fatal(http.ListenAndServe(":8080", r))
}

func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := "karrlein" // os.Getenv("PROJECT_ID")

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client
}
