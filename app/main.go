package main

import (
	"log"
	"os"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	app.Route("/", &home{})
	app.RouteWithRegexp("^/article.*", &article{})
	app.Route("/fussball.html", &fussball{})
	app.Route("/korbball.html", &korbball{})
	app.Route("/sportheim.html", &sportheim{})
	app.Route("/news.html", &news{})
	app.RunWhenOnBrowser()

	if os.Getenv("GITHUB") == "TRUE" {
		err := app.GenerateStaticWebsite(".", &app.Handler{
			Name:        "SpVgg Hambach",
			Title:       "SpVgg Hambach",
			Description: "Webiste of SpVgg Hambach",
			Icon: app.Icon{
				Default:    "/web/images/hambach_logo_192.png", // Specify default favicon.
				Large:      "/web/images/hambach_logo_512.png",
				AppleTouch: "/web/images/hambach_logo_192.png", // Specify icon on IOS devices.
			},
			Resources: app.GitHubPages("hambach"),
			Styles: []string{
				"https://cdn.jsdelivr.net/npm/bulma@0.9.1/css/bulma.min.css",
				"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.2/css/all.min.css",
				"/web/css/main.css",
				"/web/css/background-test.css",
			},
			ThemeColor: "#008000",
			Env: app.Environment{
				"READ_KEY": os.Getenv("READ_KEY"),
			},
		})
		if err != nil {
			log.Fatal(err)
		}
	} else {
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
			Env: app.Environment{
				"READ_KEY": os.Getenv("READ_KEY"),
			},
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}
