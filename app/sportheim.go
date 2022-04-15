package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type sportheim struct {
	app.Compo

	item Content
}

func (sp *sportheim) Render() app.UI {
	image_list := []string{
		"/web/images/sportheim/Werb4.jpg",
		"/web/images/sportheim/Werb5.jpg",
		"/web/images/sportheim/Werb2.jpg",
		"/web/images/sportheim/PICT0040.jpg",
		"/web/images/sportheim/biergarten.jpg",
		"/web/images/sportheim/gerichte6.jpg",
		"/web/images/sportheim/kalamari.jpg",
		"/web/images/sportheim/spileplatz.jpg",
		"/web/images/sportheim/stammtisch.jpg",
		"/web/images/sportheim/terasse.jpg",
		"/web/images/sportheim/vorspeise.jpg",
	}
	images := chunksOfStrings(image_list, 4)

	return app.Div().Class("bg").Body(
		&navbar{},
		app.Section().Class("section is-medium").Body(
			app.Div().Class("card").Body(
				app.Div().Class("card-content").Body(
					app.Div().Class("media").Body(
						app.Div().Class("media-content").Body(
							app.P().Class("title").Text(sp.item.Title),
						),
					),
					app.Div().Class("content").Body(
						app.Raw(sp.item.Content),
					),
				),
			),
			app.Div().Class("tile is-ancestor is-vertical").Body(
				app.Range(images).Slice(func(i int) app.UI {
					return app.Div().Class("tile is-parent is-horizontal").Body(
						app.Range(images[i]).Slice(func(j int) app.UI {
							return app.Div().Class("tile is-parent is-3").Body(
								app.Div().Class("tile is-child card").Body(
									app.Div().Class("card-image").Body(
										app.Figure().Class("image").Body(
											app.Img().Src(images[i][j]),
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

func (sportheim *sportheim) OnNav(ctx app.Context) {
	// Launching a new goroutine:
	ctx.Async(func() {
		app_key := app.Getenv("READ_KEY")
		r, err := http.Get("https://api.spvgg-hambach.de/api/v1/content/1?appkey=" + app_key)
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

		sportheim.item = content
		sportheim.Update()
	})
}
