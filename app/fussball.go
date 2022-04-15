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
				app.A().Class("navbar-item").Href("/12/article.html").Text(
					"Mannschaft",
				),
				app.A().Class("navbar-item").Href("/13/article.html").Text(
					"Ergebnisse und Tabellen",
				),
				app.A().Class("navbar-item").Href("/11/article.html").Text(
					"Kontakt",
				),
				app.Div().Class("navbar-item has-dropdown").Class(fn.youthDropdown).Body(
					app.Span().Class("navbar-link").Text(
						"Nachwuchs",
					).OnClick(fn.youthDropdownClick),
					app.Div().Class("navbar-dropdown").Body(
						app.A().Class("navbar-item").Href("/15/article.html").Text(
							"SG",
						),
						app.A().Class("navbar-item").Href("/16/article.html").Text(
							"U18",
						),
						app.A().Class("navbar-item").Href("/17/article.html").Text(
							"U15-1",
						),
						app.A().Class("navbar-item").Href("/18/article.html").Text(
							"U15-2",
						),
						app.A().Class("navbar-item").Href("/19/article.html").Text(
							"U13-1",
						),
						app.A().Class("navbar-item").Href("/20/article.html").Text(
							"U13-2",
						),
						app.A().Class("navbar-item").Href("/55/article.html").Text(
							"U11-1",
						),
						app.A().Class("navbar-item").Href("/56/article.html").Text(
							"U11-2",
						),
						app.A().Class("navbar-item").Href("/57/article.html").Text(
							"U9",
						),
						app.A().Class("navbar-item").Href("/58/article.html").Text(
							"U8",
						),
						app.A().Class("navbar-item").Href("/59/article.html").Text(
							"U7",
						),
					),
				),

				app.A().Class("navbar-item").Href("/14/article.html").Text(
					"Senioren",
				),
				app.A().Class("navbar-item").Href("/72/article.html").Text(
					"Chronik",
				),
				app.Div().Class("navbar-item has-dropdown").Class(fn.dropdown).Body(
					app.Span().Class("navbar-link").Text(
						"Abteilungen",
					).OnClick(fn.dropdownClick),
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
				app.Img().Src("https://storage.googleapis.com/hambach/cid%3A9853829C-82D4-4179-B37D-76F376B0E4E6.jpeg"),
			),
		),
		&footer{},
	)
}
