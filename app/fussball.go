package main

import "github.com/maxence-charriere/go-app/v7/pkg/app"

type fussballNavbar struct {
	app.Compo

	active   string
	dropdown string
}

type fussball struct {
	app.Compo

	article               [][]Content
	article_without_chunk []Content
	item                  Content
}

func (fn *fussballNavbar) Render() app.UI {
	return app.Nav().Class("navbar is-success is-fixed-top").Body(
		app.Div().Class("navbar-brand").Body(
			app.A().Class("navbar-item").Href("/").Body(
				app.Img().Src("/web/images/hambach_wappen.png"),
				app.H3().Class("title title-brand").Text("SpVgg Hambach"),
			),
			app.Span().Class("navbar-burger").Class(fn.active).Body(
				app.Span(),
				app.Span(),
				app.Span(),
			).OnClick(fn.onClick),
		),
		app.Div().Class("navbar-menu").ID("navbarMenu").Class(fn.active).Body(
			app.Div().Class("navbar-start").Body(
				app.Div().Class("navbar-item").Body(
					app.A().Text(
						"Fussball",
					).Href("/fussball"),
				),
				app.Div().Class("navbar-item").Body(
					app.A().Text(
						"Mannschaft",
					).Href("/article/12"),
				),
				app.Div().Class("navbar-item").Body(
					app.A().Text(
						"Ergebnisse und Tabellen",
					).Href("/article/13"),
				),
				app.Div().Class("navbar-item").Body(
					app.A().Text(
						"Kontakt",
					).Href("/article/11"),
				),
				// NACHWUCHS DROPDOWN
				app.Div().Class("navbar-item").Body(
					app.A().Text(
						"Senioren",
					).Href("/article/14"),
				),
				app.Div().Class("navbar-item").Body(
					app.A().Text(
						"Chronik",
					).Href("/article/72"),
				),
				app.Div().Class("navbar-item has-dropdown").Class(fn.dropdown).Body(
					app.Span().Class("navbar-link").Text(
						"Abteilungen",
					).OnClick(fn.dropdownClick),
					app.Div().Class("navbar-dropdown").Body(
						app.A().Class("navbar-item").Text(
							"Fussball",
						).Href("/fussball"),
						app.A().Class("navbar-item").Text(
							"Korbball",
						).Href("/korbball"),
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
							"FitMixx",
						).Href("/article/37"),
						app.A().Class("navbar-item").Text(
							"Pilates",
						).Href("/article/36"),
						app.A().Class("navbar-item").Text(
							"Wirbelsäulengymnastik",
						).Href("/article/35"),
						app.A().Class("navbar-item").Text(
							"Kinderturnen",
						).Href("/article/34"),
						app.A().Class("navbar-item").Text(
							"Gymnastik",
						).Href("/article/33"),
						app.A().Class("navbar-item").Text(
							"Freizetigruppe",
						).Href("/article/32"),
						app.A().Class("navbar-item").Text(
							"Aqua-Fitness",
						).Href("/article/40"),
						app.A().Class("navbar-item").Text(
							"Hallenbelegung",
						).Href("/article/87"),
					),
				),
			),
		),
	)
}

func (fn *fussballNavbar) onClick(ctx app.Context, e app.Event) {
	if fn.active == "is-active" {
		fn.active = ""
	} else {
		fn.active = "is-active"
	}
	fn.Update()
}
func (fn *fussballNavbar) dropdownClick(ctx app.Context, e app.Event) {
	if fn.dropdown == "is-active" {
		fn.dropdown = ""
	} else {
		fn.dropdown = "is-active"
	}
	fn.Update()
}

func (f *fussball) Render() app.UI {
	return app.Div().Class("bg").Body(
		&fussballNavbar{},
		app.Section().Class("section is-medium").Body(),
		&footer{},
	)
}
