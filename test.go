package main

import (
	"fmt"

	"github.com/openalpr/openalpr/src/bindings/go/openalpr"
)

func main() {
	alpr := openalpr.NewAlpr("eu", "", "")
	alpr.SetDefaultRegion("il")
	defer alpr.Unload()

	if !alpr.IsLoaded() {
		fmt.Println("OpenAlpr failed to load!")
		return
	}
	alpr.SetTopN(10)

	fmt.Println(alpr.IsLoaded())
	fmt.Println(openalpr.GetVersion())

	resultFromFilePath, err := alpr.RecognizeByFilePath("plate.jpg")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", resultFromFilePath)
	fmt.Printf("\n\n\n")
}
