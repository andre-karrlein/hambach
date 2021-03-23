package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
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

func (sp *sportheim) OnNav(ctx app.Context, u *url.URL) {
	go sp.doItemRequest()
}

func (sp *sportheim) doItemRequest() {
	resp, err := http.Get("/api/v1/articles?id=1")
	if err != nil {
		log.Println(err)
		var contentList contentList
		app.LocalStorage.Get("content", &contentList)

		contentKey := 0
		contentID, _ := strconv.Atoi("1")
		for index, element := range contentList.Content {
			if element.ID == contentID {
				contentKey = index
			}
		}
		sp.updateItemResponse(contentList.Content[contentKey])
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

	var content Content
	json.Unmarshal([]byte(sb), &content)

	sp.updateItemResponse(content)
}

func (sp *sportheim) updateItemResponse(content Content) {
	app.Dispatch(func() {
		sp.item = content
		sp.Update()
	})
}
