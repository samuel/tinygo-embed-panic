package main

import (
	"embed"
	"fmt"
	"io"
	"path/filepath"
)

//go:embed embeds/*
var embedFS embed.FS

func init() {
	entries, err := embedFS.ReadDir("embeds")
	if err != nil {
		panic(err)
	}
	for _, e := range entries {
		name := e.Name()
		switch ext := filepath.Ext(name); ext {
		case ".txt":
			f, err := embedFS.Open("embeds/" + name)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			b, _ := io.ReadAll(f)
			_ = b
		case ".nfo":
			f, err := embedFS.Open("embeds/" + name)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			b, _ := io.ReadAll(f)
			_ = b
		default:
			panic(fmt.Sprintf("Unhandled extension %q", ext))
		}
	}
}

func main() {
}