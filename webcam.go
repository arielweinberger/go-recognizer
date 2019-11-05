package main

import (
	"log"
	"gocv.io/x/gocv"
)

func webcam() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Fatalf("error opening web cam: %v", err)
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	window := gocv.NewWindow("brio")
	defer window.Close();

	for {
		if ok := webcam.Read(&img); !ok || img.Empty() {
			log.Println("Unable to read from the webcam")
			continue
		}

		//log.Println(img.ToBytes())
		window.IMShow(img)
		window.WaitKey(50)
	}
}