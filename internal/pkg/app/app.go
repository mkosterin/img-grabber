package app

import (
	"fmt"
	"img-grabber/internal/app/argparser"
	"img-grabber/internal/app/image"
	"img-grabber/internal/app/manifest"
	"log"
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
		log.Fatal(err)
	}
	bbb := image.New(ma.Manifest)
	for key := range bbb.Count() {
		fmt.Printf("%v\n", key)
	}
	//bbb.Count()
	//a.args.Parse()
	//fmt.Print(*a.args)
	return nil
}
