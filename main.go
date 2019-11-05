package main

import (
	"fmt"
)

import "github.com/openalpr/openalpr/src/bindings/go/openalpr"

func main() {
	webcam()

	alpr := openalpr.NewAlpr("eu", "openalpr.conf", "");
	alpr.SetDefaultRegion("nl");
	defer alpr.Unload()

	if !alpr.IsLoaded() {
		fmt.Println("OpenAlpr failed to load!")
		return
	}
	alpr.SetTopN(10)

	fmt.Println(alpr.IsLoaded())
	fmt.Println(openalpr.GetVersion())

	//resultFromFilePath, err := alpr.RecognizeByFilePath("plate.jpg")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Printf("%+v\n", resultFromFilePath.Plates[0].BestPlate)
	//fmt.Printf("\n\n\n")

	//imageBytes, err := ioutil.ReadFile("lp.jpg")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//resultFromBlob, err := alpr.RecognizeByBlob(imageBytes)
	//fmt.Printf("%+v\n", resultFromBlob)
}