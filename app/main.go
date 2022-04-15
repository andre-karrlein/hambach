package main

import (
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	app.Route("/", &home{})
	app.RouteWithRegexp("^/article.*", &article{})
	app.Route("/fussball", &fussball{})
	app.Route("/korbball", &korbball{})
	app.Route("/sportheim", &sportheim{})
	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite(".", &app.Handler{
		Name:        "SpVgg Hambach",
		Title:       "SpVgg Hambach",
		Description: "Webiste of SpVgg Hambach",
		Icon: app.Icon{
			Default:    "/web/images/hambach_logo_192.png", // Specify default favicon.
			Large:      "/web/images/hambach_logo_512.png",
			AppleTouch: "/web/images/hambach_logo_192.png", // Specify icon on IOS devices.
		},
		Styles: []string{
			"https://cdn.jsdelivr.net/npm/bulma@0.9.1/css/bulma.min.css",
			"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.2/css/all.min.css",
			"/web/css/main.css",
		},
		ThemeColor: "#008000",
	})

	if err != nil {
		log.Fatal(err)
	}
}
