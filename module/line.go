package module

import "github.com/fogleman/gg"

type Line struct {
	StartX float64 `json:"startX"`
	StartY float64 `json:"startY"`
	EndX   float64 `json:"endX"`
	EndY   float64 `json:"endY"`
	Width  float64 `json:"width"`
	Color  string  `json:"color"`
}

func (line *Line) Draw(dc *gg.Context) error {
	dc.SetLineWidth(line.Width)
	dc.SetHexColor(line.Color)
	dc.DrawLine(line.StartX, line.StartY, line.EndX, line.EndY)
	dc.Stroke()
	return nil
}
