package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type article struct {
	app.Compo

	item Content
	navbar app.UI
}

func (a *article) Render() app.UI {

	image := strings.Replace(a.item.Image, "public", "web", 1)
	content := strings.ReplaceAll(a.item.Content, "/public", "")
	content = strings.ReplaceAll(content, "/images/", "/web/images/")
	image = strings.Replace(image, "https://storage.googleapis.com/hambach/", "https://hambach.s3.eu-central-1.amazonaws.com/", 1)
	content = strings.ReplaceAll(content, "https://storage.googleapis.com/hambach/", "https://hambach.s3.eu-central-1.amazonaws.com/")

	return app.Div().Class("bg").Body(
		a.navbar,
		app.Section().Class("section is-medium").Body(
			app.Div().Class("card").Body(
				app.Div().Class("card-content").Body(
					app.Br(),
					app.Div().Class("media").Body(
						app.Div().Class("media-left").Body(
							app.Figure().Class("image is-128x128").Body(
								app.Img().Src(image),
							),
						),
						app.Div().Class("media-content").Body(
							app.P().Class("title").Text(a.item.Title),
							app.P().Class("subtitle").Text(a.item.Date),
						),
					),
					app.Div().Class("content").Body(
						app.Raw("<div>"+content+"</div>"),
					),
				),
			),
		),
		&footer{},
	)
}

func getNavbar(category string) app.UI {
	if category == "Fussball" {
		return &fussballNavbar{}
	}
	if category == "Korbball" {
		return &korbballNavbar{}
	}
	return &navbar{}
}

func (article *article) OnNav(ctx app.Context) {
	path := strings.Split(ctx.Page().URL().Path, "/")
	id := path[2]
	// Launching a new goroutine:
	ctx.Async(func() {
		app_key := app.Getenv("READ_KEY")
		r, err := http.Get("https://api.spvgg-hambach.de/api/v1/content/" + id + "?appkey=" + app_key)
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

		var content Content
		json.Unmarshal([]byte(sb), &content)

		article.navabr = getNavbar(a.item.Category)
		article.item = content
		article.Update()
	})
}
