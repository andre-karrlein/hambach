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
				app.A().Class("navbar-item").Href("/news.html").Text(
					"News",
				),
				app.A().Class("navbar-item").Href("/article/47").Text(
					"Termine",
				),
				app.A().Class("navbar-item").Href("/sportheim.html").Text(
					"Sportheim",
				),
				app.Div().Class("navbar-item has-dropdown").Class(n.dropdown).Body(
					app.Span().Class("navbar-link").Text(
						"Abteilungen",
					),
					app.Div().Class("navbar-dropdown").Body(
						app.A().Class("navbar-item").Text(
							"Fussball",
						).Href("/fussball.html"),
						app.A().Class("navbar-item").Text(
							"Korbball",
						).Href("/article/22"),
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
							"Kinderturnen",
						).Href("/article/34"),
						app.A().Class("navbar-item").Text(
							"Gymnastik",
						).Href("/article/ee40fc09-58be-486e-9bd5-9e8a7de96cfc"),
						app.A().Class("navbar-item").Text(
							"Freizetigruppe",
						).Href("/article/32"),
						app.A().Class("navbar-item").Text(
							"Fitness und Gesundheit",
						).Href("/article/09bb29a6-7eb2-4f80-90c7-c2fc414a05f6"),
						app.A().Class("navbar-item").Text(
							"Hallenbelegung",
						).Href("/article/87"),
					),
				).OnClick(n.dropdownClick),
				app.A().Class("navbar-item").Href("/article/3").Text(
					"Vorstandschaft",
				),
				app.A().Class("navbar-item").Href("/article/2").Text(
					"Vereinschronik",
				),
				app.A().Class("navbar-item").Href("/article/48").Text(
					"Mitgliedschaft",
				),
				app.A().Class("navbar-item").Href("/article/101").Body(
					app.Text("SPVGG als APP"),
					app.Span().Class("tag is-danger is-rounded").Text(
						"NEW",
					),
				),
			),
		).OnClick(n.onClick),
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
