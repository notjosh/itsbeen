///usr/bin/env go run $0 $@ ; exit

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
	m.Use(render.Renderer())

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
		}{
			Duration: duration,
			Since: since,
		})
	})

	m.NotFound(func(r render.Render) {
		r.Redirect("/")
	})

	m.Run()
}