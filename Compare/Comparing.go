package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

func compare(path string, path2 string) {
	var Path map[string]int
	Path = map[string]int{}
	memPath := path

	var Path2 map[string]int
	Path2 = map[string]int{}
	memPath2 := path2

	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			trimPath := path[len(memPath):]
			Path[trimPath] = int(info.Size())
			return nil
		})

	err2 := filepath.Walk(path2,
		func(path2 string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			trimPath := path2[len(memPath2):]
			Path2[trimPath] = int(info.Size())
			return nil
		})

	for key, val := range Path {
		_, found := Path2[key]
		if found {
			if val == Path2[key] {
			} else {
				fmt.Println(memPath, key, "  \t **MODIFIED**")
			}
		} else {
			fmt.Println(memPath, key, "  \t **NEW**")
		}

	}
	for key, _ := range Path2 {
		_, found := Path[key]
		if !found {
			fmt.Println(memPath2, key, "  \t: **DELETED**")
		}
	}

	if err != nil {
		log.Println(err)
	}
	if err2 != nil {
		log.Println(err2)
	}
}
func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "Compare",
			Aliases: []string{"c"},
			Usage:   "Compare Two Folder",
			Action: func(c *cli.Context) {
        var Folder1Path String
        var Folder2Path String
				fmt.Scan(Folder1Path)
				fmt.Scan(Folder2Path)
				compare(Folder1Path, Folder2Path)
			},
		},
	}
}
func main() {
	commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
