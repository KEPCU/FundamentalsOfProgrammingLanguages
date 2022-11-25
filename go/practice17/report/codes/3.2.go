package main

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"sync"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	path := "8.jpg"

	start := time.Now()
	Parallel(path)
	elapsed := time.Since(start)
	fmt.Println("Parallel time: ", elapsed)
}

func Parallel(path string) {
	imagePath, err := os.Open(path)
	check(err)
	defer imagePath.Close()

	myImage, _, err := image.Decode(imagePath)
	check(err)

	size := myImage.Bounds().Size()

	var redValues plotter.Values
	var greenValues plotter.Values
	var blueValues plotter.Values

	go func() {
		waitGroup := new(sync.WaitGroup)
		waitGroup.Add(1)
		for x := 0; x < size.X; x++ {
			for y := 0; y < size.Y; y++ {
				pixel := myImage.At(x, y)
				originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
				redValues = append(redValues, float64(originalColor.R))
			}
		}

		Histogram(redValues, "Red channel histogram (Parallel)", color.RGBA{R: 255, G: 0, B: 0, A: 255}, "redHistogramParallel.png")
	}()

	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			pixel := myImage.At(x, y)
			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
			greenValues = append(greenValues, float64(originalColor.G))
			blueValues = append(blueValues, float64(originalColor.B))
		}
	}

	Histogram(greenValues, "Green channel histogram (Parallel)", color.RGBA{R: 0, G: 255, B: 0, A: 255}, "greenHistogramParallel.png")
	Histogram(blueValues, "Blue channel histogram (Parallel)", color.RGBA{R: 0, G: 0, B: 255, A: 255}, "blueHistogramParallel.png")
}

func Histogram(values plotter.Values, title string, color color.Color, name string) {
	myPlot := plot.New()
	myPlot.Title.Text = title
	myPlot.X.Max = 260
	myPlot.X.Label.Text = "intensity"
	myPlot.Y.Label.Text = "number of pixel"

	histogram, err := plotter.NewHist(values, 300)
	check(err)
	histogram.Color = color

	myPlot.Add(histogram)
	myPlot.Save(4*vg.Inch, 6*vg.Inch, name)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
