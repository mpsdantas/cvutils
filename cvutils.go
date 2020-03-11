package cvutils

import "gocv.io/x/gocv"

type Image struct {
	Mat gocv.Mat
}

type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

func (m Image) GetColor(row int, col int) Color {
	ch := m.Mat.Channels()
	color := Color{}

	color.Blue = m.Mat.GetUCharAt(row, col*ch+0)
	color.Green = m.Mat.GetUCharAt(row, col*ch+1)
	color.Red = m.Mat.GetUCharAt(row, col*ch+2)

	return color
}

func (m Image) SetColor(row int, col int, color Color) {
	ch := m.Mat.Channels()

	m.Mat.SetUCharAt(row, col*ch+0, color.Blue)
	m.Mat.SetUCharAt(row, col*ch+1, color.Green)
	m.Mat.SetUCharAt(row, col*ch+2, color.Red)
}
