package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"strconv"
	"sync"
	"time"

	"os"
)

func main() {
	path := "1.jpg"
	path2 := "4.jpg"

	start := time.Now()
	Parallel(path, path2)
	elapsed := time.Since(start)
	fmt.Println("Parallel time: ", elapsed)
}

func Parallel(path string, path2 string) {
	imagePath, err := os.Open(path)
	defer imagePath.Close()
	myImage, _, err := image.Decode(imagePath)

	imagePath2, err := os.Open(path2)
	defer imagePath2.Close()
	myImage2, _, err := image.Decode(imagePath2)

	level := 0

	wg := new(sync.WaitGroup)

	goroutines := 3

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		level += 25
		go func(mylevel int) {
			bounds := myImage2.Bounds()
			union := image.NewAlpha(bounds)

			waitGroup := new(sync.WaitGroup)

			for h := 0; h < 2; h++ {
				startX := (h) * bounds.Dx() / 2
				endX := (h + 1) * bounds.Dx() / 2
				waitGroup.Add(1)
				go func() {
					for j := startX; j < endX; j++ {
						for k := 0; k < bounds.Dy(); k++ {
							union.SetAlpha(j, k, color.Alpha{uint8(255 * mylevel / 100)})
						}
					}
					defer waitGroup.Done()
				}()
			}
			waitGroup.Wait()

			myImage3 := image.NewRGBA(bounds)
			draw.Draw(myImage3, myImage3.Bounds(), myImage, image.ZP, draw.Src)
			draw.DrawMask(myImage3, bounds, myImage2, image.ZP, union, image.ZP, draw.Over)

			result, _ := os.Create("resultParallel2_" + strconv.Itoa(mylevel) + ".jpg")
			defer result.Close()
			err = jpeg.Encode(result, myImage3, nil)
			check(err)
			defer wg.Done()
		}(level)
	}
	wg.Wait()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
