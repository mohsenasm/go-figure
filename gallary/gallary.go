package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"figure"
)

func main() {
	files, err := filepath.Glob(filepath.Join("../fonts", "*.flf"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("# Font Gallary")
	for _, file := range files {
		fontName := strings.TrimSuffix(strings.TrimPrefix(file, "../fonts/"), ".flf")
		fmt.Println("##", fontName)
		fmt.Println("```")
		myFigure := figure.NewFigure("go-figure", fontName, true)
		myFigure.Print()
		fmt.Println("```")
	}
}
