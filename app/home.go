package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type home struct {
	app.Compo

	article               [][]Content
	article_without_chunk []Content
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
									app.A().Href("/article/"+h.article[i][j].ID).Body(
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

func (home *home) OnMount(ctx app.Context) {
	// Launching a new goroutine:
	ctx.Async(func() {
		app_key := app.Getenv("READ_KEY")
		r, err := http.Get("https://api.spvgg-hambach.de/api/v1/content?appkey=" + app_key)
		if err != nil {
			app.Log(err)
			return
		}
		defer r.Body.Close()

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			app.Log(err)
			return
		}

		sb := string(b)

		var result []Content
		json.Unmarshal([]byte(sb), &result)

		var content []Content
		for _, element := range result {
			if element.Type == "article" && element.Category == "Allgemein" {
				content = append(content, element)
			}
		}

		sort.Slice(content, func(i, j int) bool {
			return content[i].ID < content[j].ID
		})
		home.article = chunkSlice(content, 4)
		home.article_without_chunk = content

		home.Update()
	})
}
