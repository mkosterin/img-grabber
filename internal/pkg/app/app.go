package app

import (
	"fmt"
	"img-grabber/internal/app/argparser"
	"img-grabber/internal/app/image"
	"img-grabber/internal/app/manifest"
	stat "img-grabber/internal/app/stats"
	"sort"
)

type App struct {
	args *argparser.Argument
}

func New() *App {
	return &App{
		args: argparser.New(),
	}
}

func (a *App) Run() error {
	ma := manifest.New(a.args.Filename)
	err := ma.Read()
	if err != nil {
		return err
	}
	switch a.args.Mode {
	case "image":
		m := image.New(ma.Manifest)
		images := m.Count()
		switch a.args.Sort {
		case "asc":
			sort.Strings(images)
		case "desc":
			sort.Sort(sort.Reverse(sort.StringSlice(images)))
		}
		for i := 0; i < len(images); i++ {
			fmt.Println(images[i])
		}
	case "stat":
		s := stat.New(ma.Manifest)
		k, m := s.ObjectCounter()
		for _, key := range k {
			fmt.Printf("%s - %d\n", key, m[key])
		}
	}

	return nil
}
