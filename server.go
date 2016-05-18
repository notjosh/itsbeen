package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"fmt"
	"os"
	"strings"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", nil)
	})

	m.Get("/:duration/since/:since", func(params martini.Params, r render.Render) {
		duration := strings.Replace(params["duration"], "-", "", -1)
		since := strings.Replace(params["since"], "-", "", -1)

		path := fmt.Sprintf(
			"public/audio/itsbeen-%s-%s.mp3",  // mp3 or ogg, doesn't matter
			duration,
			since)

		if _, err := os.Stat(path); os.IsNotExist(err) {
			r.Redirect("/")
		}

		r.HTML(200, "play", struct {
			Duration string
			Since string
			DurationTitle string
			SinceTitle string
		}{
			Duration: duration,
			Since: since,
			DurationTitle: strings.Replace(params["duration"], "-", " ", -1),
			SinceTitle: strings.Replace(params["since"], "-", " ", -1),
		})
	})

	m.NotFound(func(r render.Render) {
		r.Redirect("/")
	})

	m.Run()
}