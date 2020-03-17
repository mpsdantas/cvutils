package cvutils

import (
	"errors"
	"fmt"
	"log"

	"gocv.io/x/gocv"
)

type ImageOptions struct {
	Name       string
	Flags      gocv.IMReadFlag
	WindowName string
	Draw       DrawOptions
}

type DrawOptions struct {
	DrawFunc      DrawFunc
	Color         Color
	StartingPoint Point
	EndPoint      Point
}

type Point struct {
	X int
	Y int
}

type DrawFunc func(image Image, color Color, point Point)

func ensureImageCanBeDraw(dimensions []int, points []Point) error {
	for _, point := range points {
		if point.X > dimensions[0] || point.Y > dimensions[1] {
			return errors.New(fmt.Sprintf("The point: %v is greater "+
				"than the upper limit of the image: x = %v y = %v", point, dimensions[0], dimensions[1]))
		}
	}

	return nil
}

func ShowIMG(opts ImageOptions) {
	img := Image{
		Mat: gocv.IMRead(opts.Name, opts.Flags),
	}

	if img.Mat.Empty() {
		log.Fatal(fmt.Sprintf("Error reading image from: %v\n", opts.Name))
	}

	defer img.Mat.Close()

	if err := ensureImageCanBeDraw(img.Mat.Size(), []Point{
		opts.Draw.StartingPoint,
		opts.Draw.EndPoint,
	}); err != nil {
		log.Fatal(err.Error())
	}

	window := gocv.NewWindow(opts.WindowName)

	for i := opts.Draw.StartingPoint.X; i < opts.Draw.EndPoint.X; i++ {
		for j := opts.Draw.StartingPoint.Y; j < opts.Draw.EndPoint.Y; j++ {
			opts.Draw.DrawFunc(img, opts.Draw.Color, Point{
				X: i,
				Y: j,
			})
		}
	}

	window.IMShow(img.Mat)
	window.WaitKey(0)
}
