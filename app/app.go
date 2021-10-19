package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type home struct {
	app.Compo

	article               [][]Content
	article_without_chunk []Content
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

func (h *home) Render() app.UI {
	return app.Div().Class("bg").Body(
		&navbar{},
		app.Section().Class("section is-medium").Body(
			app.Div().Class("tile is-ancestor is-vertical").Body(
				app.Range(h.article).Slice(func(i int) app.UI {
					return app.Div().Class("tile is-parent is-horizontal").Body(
						app.Range(h.article[i]).Slice(func(j int) app.UI {
							return app.Div().Class("tile is-parent is-3").Body(
								app.If(h.article[i][j].Link == "",
									app.A().Href("/article/"+strconv.Itoa(h.article[i][j].ID)).Body(
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
								).Else(
									app.A().Href(h.article[i][j].Link).Body(
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

func chunksOfStrings(slice []string, chunkSize int) [][]string {
	var chunks [][]string
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

func main() {
	app.Route("/", &home{})
	app.RouteWithRegexp("^/article.*", &article{})
	app.Route("/fussball", &fussball{})
	app.Route("/korbball", &korbball{})
	app.Route("/sportheim", &sportheim{})
	app.Run()
}
