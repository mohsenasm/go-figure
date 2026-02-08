package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"figure"
)

func main() {
	files, err := filepath.Glob(filepath.Join("../fonts", "*.flf"))
	if err != nil {
		log.Fatal(err)
	}

	output, err := os.Create("gallery.md")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer output.Close()

	output.WriteString("# Font Gallary\n")
	for _, file := range files {
		fontName := strings.TrimSuffix(strings.TrimPrefix(file, "../fonts/"), ".flf")
		output.WriteString("## " + fontName + "\n")
		output.WriteString("<pre>\n")
		myFigure := figure.NewFigure("go-figure", fontName, true)
		_, err = output.WriteString(myFigure.String())
		output.WriteString("</pre>\n")
	}
}
