package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type home struct {
	app.Compo

	article []Article
}

func (h *home) Render() app.UI {
	return app.Div().Class("bg").Body(
		&navbar{},
		app.Section().Class("section is-medium").Body(
			app.Div().Class("columns is-multiline").Body(
				app.Range(h.article).Slice(func(i int) app.UI {
					return app.Div().Class("column is-one-quarter").Body(
						app.If(h.article[i].Link == "",
							app.A().Href("/article/"+h.article[i].ID).Body(
								app.Div().Class("card equal-height").Style("background-color", "#008000").Body(
									app.Div().Class("card-image card-image-half").Body(
										app.Figure().Class("image").Body(
											app.Img().Src(h.article[i].Image),
										),
									),
									app.Div().Class("card-content").Body(
										app.P().Class("subtitle content-size").Style("color", "white").Text(h.article[i].Title),
									),
								),
							),
						).Else(
							app.A().Href(h.article[i].Link).Body(
								app.Div().Class("card equal-height").Style("background-color", "#008000").Body(
									app.Div().Class("card-image card-image-half").Body(
										app.Figure().Class("image").Body(
											app.Img().Src(h.article[i].Image),
										),
									),
									app.Div().Class("card-content").Body(
										app.P().Class("subtitle content-size").Style("color", "white").Text(h.article[i].Title),
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

func (home *home) OnMount(ctx app.Context) {
	// Launching a new goroutine:
	ctx.Async(func() {
		app_key := app.Getenv("READ_KEY")
		r, err := http.Get("https://api.spvgg-hambach.de/api/v1/article?appkey=" + app_key)
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

		var result []Article
		json.Unmarshal([]byte(sb), &result)

		var article []Article
		for _, element := range result {
			if element.Type == "ARTICLE" && element.Category == "ALLGEMEIN" {
				article = append(article, element)
			}
		}

		sort.Slice(article, func(i, j int) bool {
			dateString := "2021-11-22"
			app.Log(article[i].Date)
			app.Log(article[j].Date)
			date1, _ := time.Parse(article[i].Date, dateString)
			date2, _ := time.Parse(article[j].Date, dateString)
			app.Log(date1)
			app.Log(date2)
			return date1.After(date2)
		})
		home.article = article

		app.Log(article)
		home.Update()
	})
}
