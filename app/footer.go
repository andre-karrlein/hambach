package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type footer struct {
	app.Compo
}

func (f *footer) Render() app.UI {
	return app.Footer().Class("footer").Body(
		app.Div().Class("content has-text-centered").Body(
			app.Div().Class("columns").Body(
				app.Div().Class("column").Body(
					app.P().Class("matches").Text("WEITERES"),
					app.Br(),
					app.A().Href("/article/6").Text("Sponsoren"),
					app.Br(),
					app.A().Href("/article/5").Text("Impressum"),
					app.Br(),
					app.A().Href("/article/4").Text("Datenschutz"),
				),
				app.Div().Class("column").Body(
					app.P().Class("matches").Text("SOCIAL"),
					app.A().Href("https://www.facebook.com/SpVggHambachFussball").Body(
						app.Span().Body(
							app.I().Class("fab fa-facebook"),
						).Class("icon"),
						app.Span().Text(
							"Facebook",
						),
					),
				),
				app.Div().Class("column").Body(
					app.P().Class("matches").Text("SPONSORING"),
					app.Div().Class("columns").Body(
						app.Div().Class("column").Body(
							app.Figure().Class("image is-96x96").Body(
								app.Img().Src("/web/images/sponsoring/atb.png"),
							),
						),
						app.Div().Class("column").Body(
							app.Figure().Class("image is-96x96").Body(
								app.Img().Src("/web/images/sponsoring/pabst.jpg"),
							),
						),
						app.Div().Class("column").Body(
							app.Figure().Class("image is-96x96").Body(
								app.Img().Src("/web/images/sponsoring/delphi.jpg"),
							),
						),
						app.Div().Class("column").Body(
							app.Figure().Class("image is-96x96").Body(
								app.Img().Src("/web/images/sponsoring/tucher.png"),
							),
						),
					),
				),
			),
			app.P().Body(
				app.Text("Made by "),
				app.A().Href("https://karrlein.com").Target("__blank").Text("André Karrlein"),
			),
		),
	)
}
