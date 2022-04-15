package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type korbballNavbar struct {
	app.Compo

	active        string
	dropdown      string
	youthDropdown string
}

type korbball struct {
	app.Compo
}

func (kn *korbballNavbar) Render() app.UI {
	return app.Nav().Class("navbar is-success is-fixed-top").Body(
		app.Div().Class("navbar-brand").Body(
			app.A().Class("navbar-item").Href("/").Body(
				app.Img().Src("/web/images/hambach_wappen.png"),
				app.H3().Class("title title-brand").Text("SpVgg Hambach"),
			),
			app.Span().Class("navbar-burger").Class(kn.active).Body(
				app.Span(),
				app.Span(),
				app.Span(),
			).OnClick(kn.onClick),
		),
		app.Div().Class("navbar-menu").ID("navbarMenu").Class(kn.active).Body(
			app.Div().Class("navbar-start").Body(
				app.A().Class("navbar-item").Href("/22/article.html").Text(
					"Kontakt",
				),
				app.A().Class("navbar-item").Href("/23/article.html").Text(
					"Chronik",
				),
				app.A().Class("navbar-item").Href("/24/article.html").Text(
					"Mannschaft I, II & III",
				),
				app.Div().Class("navbar-item has-dropdown").Class(kn.youthDropdown).Body(
					app.Span().Class("navbar-link").Text(
						"Nachwuchs",
					).OnClick(kn.youthDropdownClick),
					app.Div().Class("navbar-dropdown").Body(
						app.A().Class("navbar-item").Href("/25/article.html").Text(
							"Jugend 19",
						),
						app.A().Class("navbar-item").Href("/26/article.html").Text(
							"Jugend 15",
						),
						app.A().Class("navbar-item").Href("/27/article.html").Text(
							"Jugend 12",
						),
						app.A().Class("navbar-item").Href("/28/article.html").Text(
							"Jugend 9",
						),
					),
				),
				app.Div().Class("navbar-item has-dropdown").Class(kn.dropdown).Body(
					app.Span().Class("navbar-link").Text(
						"Abteilungen",
					).OnClick(kn.dropdownClick),
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

func (kn *korbballNavbar) onClick(ctx app.Context, e app.Event) {
	if kn.active == "is-active" {
		kn.active = ""
	} else {
		kn.active = "is-active"
	}
	kn.Update()
}
func (kn *korbballNavbar) dropdownClick(ctx app.Context, e app.Event) {
	if kn.dropdown == "is-active" {
		kn.dropdown = ""
	} else {
		kn.dropdown = "is-active"
	}
	kn.Update()
}

func (kn *korbballNavbar) youthDropdownClick(ctx app.Context, e app.Event) {
	if kn.youthDropdown == "is-active" {
		kn.youthDropdown = ""
	} else {
		kn.youthDropdown = "is-active"
	}
	kn.Update()
}

func (k *korbball) Render() app.UI {
	return app.Div().Class("bg").Body(
		&korbballNavbar{},
		app.Section().Class("section is-medium").Body(),
		&footer{},
	)
}
