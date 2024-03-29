package app

import (
	"fmt"
	"img-grabber/internal/app/argparser"
	"img-grabber/internal/app/image"
	"img-grabber/internal/app/manifest"
	stat "img-grabber/internal/app/stats"
	"os"
	"sort"
	"text/tabwriter"
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
	mnfst := manifest.New(a.args.Filename)
	err := mnfst.Read() //if no manifest we will exit by timeout upper
	if err != nil {
		return err
	}
	switch a.args.Mode {
	case "image":
		m := image.New(mnfst.Manifest)
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
		s := stat.New(mnfst.Manifest)
		k, m := s.ObjectCounter()

		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 8, 8, 0, '\t', 0)
		defer w.Flush()

		fmt.Fprintf(w, "\n %s\t%s\t", "Object", "Count")
		fmt.Fprintf(w, "\n %s\t%s\t", "------", "-----")
		for _, key := range k {
			fmt.Fprintf(w, "\n %s\t%d\t", key, m[key])
		}

	}

	return nil
}
