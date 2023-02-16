package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"strconv"

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
			app.Div().Class("columns is-multiline is-mobile").Body(
				app.Range(h.article_without_chunk).Slice(func(i int) app.UI {
					image := strings.Replace(h.article_without_chunk[i].Image, "https://storage.googleapis.com/hambach/", "https://hambach.s3.eu-central-1.amazonaws.com/", 1)
					link := strings.Replace(h.article_without_chunk[i].Link, "https://storage.googleapis.com/hambach/", "https://hambach.s3.eu-central-1.amazonaws.com/", 1)
			
					return app.Div().Class("column is-one-quarter").Body(
						app.If(h.article_without_chunk[i].Link == "",
							app.A().Href("/article/"+h.article_without_chunk[i].ID).Body(
								app.Div().Class("card").Style("background-color", "#008000").Body(
									app.Div().Class("card-image").Body(
										app.Figure().Class("image").Body(
											app.Img().Src(image),
										),
									),
									app.Div().Class("card-content").Body(
										app.P().Class("subtitle").Style("color", "white").Text(h.article_without_chunk[i].Title),
									),
								),
							),
						).Else(
							app.A().Href(link).Body(
								app.Div().Class("card").Style("background-color", "#008000").Body(
									app.Div().Class("card-image").Body(
										app.Figure().Class("image").Body(
											app.Img().Src(image),
										),
									),
									app.Div().Class("card-content").Body(
										app.P().Class("subtitle").Style("color", "white").Text(h.article_without_chunk[i].Title),
									),
								),
							),
						),
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
			content_i, _ := strconv.Atoi(content[i].ID)
			content_j, _ := strconv.Atoi(content[j].ID)
			return content_i > content_j
		})
		home.article = chunkSlice(content, 4)
		home.article_without_chunk = content

		app.Log(content)
		home.Update()
	})
}
