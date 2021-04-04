package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type article struct {
	app.Compo

	item Content
}

func (a *article) Render() app.UI {

	image := strings.Replace(a.item.Image, "public", "web", 1)
	content := strings.ReplaceAll(a.item.Content, "/public", "")
	content = strings.ReplaceAll(content, "/images/", "/web/images/")

	navbar := getNavbar(a.item.Category)
	return app.Div().Class("bg").Body(
		navbar,
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

func (a *article) OnNav(ctx app.Context, u *url.URL) {
	path := strings.Split(u.Path, "/")
	id := path[2]
	go a.doItemRequest(id)
}

func (a *article) doItemRequest(id string) {
	resp, err := http.Get("/api/v1/articles?id=" + id)
	if err != nil {
		log.Println(err)
		var contentList contentList
		app.LocalStorage.Get("content", &contentList)

		contentKey := 0
		contentID, _ := strconv.Atoi(id)
		for index, element := range contentList.Content {
			if element.ID == contentID {
				contentKey = index
			}
		}
		a.updateItemResponse(contentList.Content[contentKey])
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

	a.updateItemResponse(content)
}

func (a *article) updateItemResponse(content Content) {
	app.Dispatch(func() {
		a.item = content
		a.Update()
	})
}
