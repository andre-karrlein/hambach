package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type fussballNavbar struct {
	app.Compo

	active        string
	dropdown      string
	youthDropdown string
}

type fussball struct {
	app.Compo
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
				app.A().Class("navbar-item").Href("/fussball.html").Text(
					"Fussball",
				),
				app.A().Class("navbar-item").Href("/article/12").Text(
					"Mannschaft",
				),
				app.A().Class("navbar-item").Href("/article/13").Text(
					"Ergebnisse und Tabellen",
				),
				app.A().Class("navbar-item").Href("/article/11").Text(
					"Kontakt",
				),
				app.Div().Class("navbar-item has-dropdown").Class(fn.youthDropdown).Body(
					app.Span().Class("navbar-link").Text(
						"Nachwuchs",
					),
					app.Div().Class("navbar-dropdown").Body(
						app.A().Class("navbar-item").Href("/article/15").Text(
							"SG",
						),
						app.A().Class("navbar-item").Href("/article/16").Text(
							"U19",
						),
						app.A().Class("navbar-item").Href("/article/17").Text(
							"U15",
						),
						app.A().Class("navbar-item").Href("/article/19").Text(
							"U13",
						),
						app.A().Class("navbar-item").Href("/article/55").Text(
							"U11",
						),
						app.A().Class("navbar-item").Href("/article/57").Text(
							"U9",
						),
						app.A().Class("navbar-item").Href("/article/59").Text(
							"U7",
						),
					),
				).OnClick(fn.youthDropdownClick),

				app.A().Class("navbar-item").Href("/article/25c12a9f-5fb8-4d86-a7d2-7ee01fd5f02a").Text(
					"Senioren",
				),
				app.A().Class("navbar-item").Href("/article/72").Text(
					"Chronik",
				),
				app.Div().Class("navbar-item has-dropdown").Class(fn.dropdown).Body(
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
				).OnClick(fn.dropdownClick),
			),
		).OnClick(fn.onClick),
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
func (fn *fussballNavbar) youthDropdownClick(ctx app.Context, e app.Event) {
	if fn.youthDropdown == "is-active" {
		fn.youthDropdown = ""
	} else {
		fn.youthDropdown = "is-active"
	}
	fn.Update()
}

func (f *fussball) Render() app.UI {
	return app.Div().Class("bg").Body(
		&fussballNavbar{},
		app.Section().Class("section is-medium").Body(
			app.Figure().Class("image").Body(
				app.Img().Src("https://spvgghambach-969750a982653-staging.s3.eu-central-1.amazonaws.com/public/Unknown-4.jpg"),
			),
		),
		&footer{},
	)
}
