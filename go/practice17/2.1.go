package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"strconv"
	"time"

	"os"
)

func main() {
	path := "9.jpg"
	path2 := "10.jpg"

	start := time.Now()
	Sequential(path, path2)
	elapsed := time.Since(start)
	fmt.Println("Sequential time: ", elapsed)
}

func Sequential(path string, path2 string) {
	imagePath, err := os.Open(path)
	defer imagePath.Close()
	myImage, _, err := image.Decode(imagePath)

	imagePath2, err := os.Open(path2)
	defer imagePath2.Close()
	myImage2, _, err := image.Decode(imagePath2)

	bounds := myImage2.Bounds()
	level := 0
	union := image.NewAlpha(bounds)

	for i := 0; i < 3; i++ {
		level += 25
		for j := 0; j < bounds.Dx(); j++ {
			for k := 0; k < bounds.Dy(); k++ {
				union.SetAlpha(j, k, color.Alpha{uint8(255 * level / 100)})
			}
		}
		myImage3 := image.NewRGBA(bounds)
		draw.Draw(myImage3, myImage3.Bounds(), myImage, image.ZP, draw.Src)
		draw.DrawMask(myImage3, bounds, myImage2, image.ZP, union, image.ZP, draw.Over)

		result, _ := os.Create("resultSequential2_" + strconv.Itoa(level) + ".jpg")
		defer result.Close()
		err = jpeg.Encode(result, myImage3, nil)
		check(err)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
