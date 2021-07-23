package utils

import (
	"sync"

	"github.com/thedevsaddam/renderer"
)

var (
	rnd     *renderer.Render
	onceRnd sync.Once
)

func GetRenderer() *renderer.Render {
	onceRnd.Do(func() {
		opts := renderer.Options{
			ParseGlobPattern: "./public/*.html",
		}

		rnd = renderer.New(opts)
	})

	return rnd
}
