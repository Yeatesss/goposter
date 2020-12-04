package module

import (
	"image"
	"math"
	"os"

	"github.com/Yeate/gowheel"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
	"github.com/qbhy/go-utils"
)

type Image struct {
	X            int     `json:"x"`
	Y            int     `json:"y"`
	Url          string  `json:"url"`
	Width        int     `json:"width"`
	Height       int     `json:"height"`
	BorderRadius float64 `json:"border_radius"`
	CircleClip   bool    `json:"circle_clip"`
}

func (img *Image) Draw(dc *gg.Context) error {
	var (
		filename = utils.Md5(img.Url)
		fullPath string
	)
	defer func() {
		if fullPath != "" {
			_ = os.Remove(fullPath)
		}
	}()
	imgPath := ImgTempDir + filename
	if exists, _ := utils.PathExists(imgPath); !exists {
		path, err := gowheel.DownloadFile(img.Url, ImgTempDir, filename)
		if err != nil {
			return err
		}
		fullPath = path
	}

	if imgInstance, err := gg.LoadImage(imgPath); err == nil {
		if img.CircleClip {
			imgInstance = img.CircleClipAction(imgInstance)
		} else {
			img.CheckWh()
			imgInstance = resize.Resize(uint(img.Width), uint(img.Height), imgInstance, resize.Lanczos3)
		}
		dc.DrawImage(imgInstance, img.X, img.Y)
		return nil
	} else {
		return err
	}
}
func (img *Image) CircleClipAction(imgInstance image.Image) image.Image {
	w := imgInstance.Bounds().Size().X
	h := imgInstance.Bounds().Size().Y

	if img.BorderRadius == 0 {
		img.BorderRadius = math.Min(float64(w), float64(h)) / 2

	}
	dc := gg.NewContext(int(img.BorderRadius*2), int(img.BorderRadius*2))

	dc.DrawCircle(float64(w/2), float64(h/2), img.BorderRadius)
	dc.Clip()
	dc.DrawImage(imgInstance, 0, 0)
	return dc.Image()
}

func (img *Image) CheckWh() {
	if img.Width == 0 || img.Height == 0 {
		width, height := gowheel.GetImageSizeFromUrl(img.Url)
		if img.Width == 0 {
			img.Width = int(width)
		}
		if img.Height == 0 {
			img.Height = int(height)

		}
	}
}
