package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type navbar struct {
	app.Compo

	active   string
	dropdown string
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
				app.A().Class("navbar-item").Href("/47/article.html").Text(
					"Termine",
				),
				app.A().Class("navbar-item").Href("/sportheim.html").Text(
					"Sportheim",
				),
				app.Div().Class("navbar-item has-dropdown").Class(n.dropdown).Body(
					app.Span().Class("navbar-link").Text(
						"Abteilungen",
					).OnClick(n.dropdownClick),
					app.Div().Class("navbar-dropdown").Body(
						app.A().Class("navbar-item").Text(
							"Fussball",
						).Href("/fussball.html"),
						app.A().Class("navbar-item").Text(
							"Korbball",
						).Href("/22/article.html"),
						app.A().Class("navbar-item").Text(
							"Theater",
						).Href("/42/article.html"),
						app.A().Class("navbar-item").Text(
							"Schützen",
						).Href("/41/article.html"),
						app.A().Class("navbar-item").Text(
							"Tischtennis",
						).Href("/39/article.html"),
						app.A().Class("navbar-item").Text(
							"Schach",
						).Href("/38/article.html"),
						app.A().Class("navbar-item").Text(
							"Kinderturnen",
						).Href("/34/article.html"),
						app.A().Class("navbar-item").Text(
							"Gymnastik",
						).Href("/33/article.html"),
						app.A().Class("navbar-item").Text(
							"Freizetigruppe",
						).Href("/32/article.html"),
						app.A().Class("navbar-item").Text(
							"Hallenbelegung",
						).Href("/87/article.html"),
					),
				),
				app.A().Class("navbar-item").Href("/3/article.html").Text(
					"Vorstandschaft",
				),
				app.A().Class("navbar-item").Href("/2/article.html").Text(
					"Vereinschronik",
				),
				app.A().Class("navbar-item").Href("/48/article.html").Text(
					"Mitgliedschaft",
				),
				app.A().Class("navbar-item").Href("/101/article.html").Body(
					app.Text("SPVGG als APP"),
					app.Span().Class("tag is-danger is-rounded").Text(
						"NEW",
					),
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
