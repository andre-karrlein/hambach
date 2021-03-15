package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type navbar struct {
	app.Compo

	active   string
	dropdown string
}

type home struct {
	app.Compo

	article               [][]Content
	article_without_chunk []Content
	item                  Content
}

type footer struct {
	app.Compo
}

// Content struct
type Content struct {
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

type contentList struct {
	Content []Content `json:"content"`
}

func (n *navbar) Render() app.UI {
	return app.Nav().Class("navbar is-success is-fixed-top").Body(
		app.Div().Class("navbar-brand").Body(
			app.A().Class("navbar-item").Href("/").Body(
				app.Img().Src("/web/images/hambach_wappen.png"),
				app.H3().Class("title title-brand").Text("SpVgg Hambach"),
			),
			app.Span().Class("navbar-burger").Class(n.active).Body(
				app.Span(),
				app.Span(),
				app.Span(),
			).OnClick(n.onClick),
		),
		app.Div().Class("navbar-menu").ID("navbarMenu").Class(n.active).Body(
			app.Div().Class("navbar-start").Body(
				app.A().Class("navbar-item").Href("/article/47").Text(
					"Termine",
				),
				app.A().Class("navbar-item").Href("/article/1").Text(
					"Sportheim",
				),
				app.Div().Class("navbar-item has-dropdown").Class(n.dropdown).Body(
					app.Span().Class("navbar-link").Text(
						"Abteilungen",
					).OnClick(n.dropdownClick),
					app.Div().Class("navbar-dropdown").Body(
						app.A().Class("navbar-item").Text(
							"Fussball",
						).Href("/fussball"),
						app.A().Class("navbar-item").Text(
							"Korbball",
						).Href("/korbball"),
						app.A().Class("navbar-item").Text(
							"Theater",
						).Href("/article/42"),
						app.A().Class("navbar-item").Text(
							"Schützen",
						).Href("/article/41"),
						app.A().Class("navbar-item").Text(
							"Tischtennis",
						).Href("/article/39"),
						app.A().Class("navbar-item").Text(
							"Schach",
						).Href("/article/38"),
						app.A().Class("navbar-item").Text(
							"FitMixx",
						).Href("/article/37"),
						app.A().Class("navbar-item").Text(
							"Pilates",
						).Href("/article/36"),
						app.A().Class("navbar-item").Text(
							"Wirbelsäulengymnastik",
						).Href("/article/35"),
						app.A().Class("navbar-item").Text(
							"Kinderturnen",
						).Href("/article/34"),
						app.A().Class("navbar-item").Text(
							"Gymnastik",
						).Href("/article/33"),
						app.A().Class("navbar-item").Text(
							"Freizetigruppe",
						).Href("/article/32"),
						app.A().Class("navbar-item").Text(
							"Aqua-Fitness",
						).Href("/article/40"),
						app.A().Class("navbar-item").Text(
							"Hallenbelegung",
						).Href("/article/87"),
					),
				),
				app.A().Class("navbar-item").Href("/article/3").Text(
					"Vorstandschaft",
				),
				app.A().Class("navbar-item").Href("/article/2").Text(
					"Vereinschronik",
				),
				app.A().Class("navbar-item").Href("/article/48").Text(
					"Mitgliedschaft",
				),
			),
		),
	)
}

func (n *navbar) onClick(ctx app.Context, e app.Event) {
	if n.active == "is-active" {
		n.active = ""
	} else {
		n.active = "is-active"
	}
	n.Update()
}
func (n *navbar) dropdownClick(ctx app.Context, e app.Event) {
	if n.dropdown == "is-active" {
		n.dropdown = ""
	} else {
		n.dropdown = "is-active"
	}
	n.Update()
}

func (h *home) Render() app.UI {
	return app.Div().Class("bg").Body(
		&navbar{},
		app.Section().Class("section is-medium").Body(
			app.Div().Class("tile is-ancestor is-vertical").Body(
				app.Range(h.article).Slice(func(i int) app.UI {
					return app.Div().Class("tile is-parent is-horizontal").Body(
						app.Range(h.article[i]).Slice(func(j int) app.UI {
							return app.Div().Class("tile is-parent is-3").Body(
								app.A().Href("/article/" + strconv.Itoa(h.article[i][j].ID)).Body(
									app.Div().Class("tile is-child card").Style("background-color", "#008000").Body(
										app.Div().Class("card-image").Body(
											app.Figure().Class("image").Body(
												app.Img().Src(h.article[i][j].Image),
											),
										),
										app.Div().Class("card-content").Body(
											app.P().Class("subtitle").Style("color", "white").Text(h.article[i][j].Title),
										),
									),
								),
							)
						}),
					)
				}),
			),
		),
		&footer{},
	)
}

func (h *home) OnMount(ctx app.Context) {
	go h.doRequest()
}

func (h *home) onClick(ctx app.Context, e app.Event) {
	id := ctx.JSSrc.Get("id").String()
	log.Println(id)
	app.Navigate("/article/" + id)
}

func (h *home) doRequest() {
	resp, err := http.Get("/api/v1/articles")
	if err != nil {
		log.Println(err)
		var contentList contentList
		app.LocalStorage.Get("content", &contentList)
		var content []Content
		for _, element := range contentList.Content {
			if element.Type == "article" && element.Category == "Allgemein" {
				content = append(content, element)
			}
		}
		h.updateResponse(content)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	//Convert the body to type string
	sb := string(body)

	var result []Content
	json.Unmarshal([]byte(sb), &result)

	contentList := contentList{result}
	app.LocalStorage.Set("content", contentList)

	var content []Content
	for _, element := range result {
		if element.Type == "article" && element.Category == "Allgemein" {
			content = append(content, element)
		}
	}
	h.updateResponse(content)
}

func (h *home) updateResponse(content []Content) {
	app.Dispatch(func() {
		h.article = chunkSlice(content, 4)
		h.article_without_chunk = content
		h.Update()
	})
}

func chunkSlice(slice []Content, chunkSize int) [][]Content {
	var chunks [][]Content
	for {
		if len(slice) == 0 {
			break
		}

		// necessary check to avoid slicing beyond
		// slice capacity
		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}

func (f *footer) Render() app.UI {
	return app.Div().Class("footer is-success").Body(
		app.Div().Text("Andre Karrlein"),
	)
}

func main() {
	app.Route("/", &home{})
	app.RouteWithRegexp("^/article.*", &article{})
	app.Route("/fussball", &fussball{})
	app.Route("/korbball", &korbball{})
	app.Run()
}
