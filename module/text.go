package module

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"golang.org/x/image/font"

	"github.com/golang/freetype/truetype"

	"github.com/fogleman/gg"
)

type DeleteLine struct {
	Draw   bool
	StartY float64 `json:"startY"`
	EndY   float64 `json:"endY"`
	Width  float64 `json:"width"`
	Color  string  `json:"color"`
}
type WordWarp struct {
	Open    bool
	MaxLine int
}
type Style struct {
	DeleteLine DeleteLine
	WordWarp   WordWarp
}
type Text struct {
	X          int     `json:"x"`
	Y          int     `json:"y"`
	Text       string  `json:"text"`
	Width      float64 `json:"width"`
	FontSize   int     `json:"fontSize"`
	Color      string  `json:"color"`
	LineHeight int     `json:"lineHeight"`
	TextAlign  string  `json:"textAlign"`
	Style      Style   `json:"style"`
	Font       *font.Face
}

const TextAlignCenter = "center"

func (t *Text) SetFont(reader io.Reader) {
	fontBytes, _ := ioutil.ReadAll(reader)
	reader = bytes.NewBuffer(fontBytes)
	f, _ := truetype.Parse(fontBytes)
	if f != nil {
		fontFace := truetype.NewFace(f, &truetype.Options{
			Size: float64(t.FontSize),
		})
		t.Font = &fontFace

	}
}

func (t *Text) DrawX(w float64) float64 {
	if t.TextAlign == TextAlignCenter {
		return float64(t.X) - w/2
	}
	return float64(t.X)
}

func (text *Text) Draw(dc *gg.Context) error {
	var words []string
	if text.LineHeight == 0 {
		text.LineHeight = text.FontSize
	}
	if text.Font == nil {
		err := fmt.Errorf("text not set font")
		return err
	}
	dc.SetFontFace(*text.Font)
	//dc.FontHeight()
	////_ = dc.LoadFontFace(viper.GetString("font.path"), float64(text.FontSize))
	dc.SetHexColor(text.Color)
	w, _ := dc.MeasureString(text.Text)
	if text.Style.WordWarp.Open {
		words = wordWrap(dc, text.Text, text.Width, text.Style.WordWarp.MaxLine)

	} else {
		words = append(words, text.Text)
	}
	for index, word := range words {
		fmt.Println(float64(text.Y + text.LineHeight*(index+1)))
		dc.DrawString(word, text.DrawX(w), float64(text.Y+text.LineHeight*(index+1)-5))
		if text.Style.DeleteLine.Draw {
			deleteLine := Line{}
			deleteLine.Color = text.Style.DeleteLine.Color
			deleteLine.StartY = text.Style.DeleteLine.StartY
			deleteLine.StartX = text.DrawX(w)
			deleteLine.EndX = text.DrawX(w) + w
			deleteLine.EndY = text.Style.DeleteLine.EndY
			deleteLine.Width = text.Style.DeleteLine.Width
			_ = deleteLine.Draw(dc)
		}
	}
	return nil
}

type measureStringer interface {
	MeasureString(s string) (w, h float64)
}

func wordWrap(m measureStringer, s string, width float64, maxLine int) []string {
	var result []string
	if maxLine == 0 {
		maxLine = 10000
	}
	for _, line := range strings.Split(s, "\n") {

		var fields []string
		for _, v := range []rune(line) {
			fields = append(fields, string(v))
		}
		if len(fields)%2 == 1 {
			fields = append(fields, "")
		}
		x := ""
		for i := 0; i < len(fields); i += 2 {
			w, _ := m.MeasureString(x + fields[i])
			if w > width {
				if x == "" {
					result = append(result, fields[i])
					x = ""
					continue
				} else {
					result = append(result, x)
					x = ""
				}
			}
			x += fields[i] + fields[i+1]
		}
		if x != "" {
			result = append(result, x)
		}
	}
	for i, line := range result {
		result[i] = strings.TrimSpace(line)
	}
	if maxLine < len(result) {
		result = result[0:maxLine]
	}
	return result
}
