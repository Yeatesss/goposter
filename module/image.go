package module

import (
	"fmt"
	"image"
	"io"
	"math"
	"os"

	"github.com/Yeate/gowheel"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
	"github.com/qbhy/go-utils"
)

type Image struct {
	X            int       `json:"x"`
	Y            int       `json:"y"`
	Url          string    `json:"url"`
	ImgData      io.Reader `json:"img_data"`
	Width        int       `json:"width"`
	Height       int       `json:"height"`
	BorderRadius float64   `json:"border_radius"`
	CircleClip   bool      `json:"circle_clip"`
}

func (img *Image) Draw(dc *gg.Context) error {
	var (
		filename    = utils.Md5(img.Url)
		path        string
		err         error
		imgInstance image.Image
	)
	defer func() {
		if path != "" {
			_ = os.Remove(path)
		}
	}()
	imgPath := ImgTempDir + filename
	if exists, _ := utils.PathExists(imgPath); !exists && img.Url != "" {
		path, err = gowheel.DownloadFile(img.Url, ImgTempDir, filename)
		if err != nil {
			return err
		}
		imgInstance, err = gg.LoadImage(imgPath)
	} else if img.ImgData != nil {
		imgInstance, _, err = image.Decode(img.ImgData)

	} else {
		err = fmt.Errorf("图片数据不存在")
	}
	if err != nil {
		return err
	}
	if img.Width == 0 || img.Height == 0 {
		g := imgInstance.Bounds()
		// Get height and width
		height := float64(g.Dy())
		width := float64(g.Dx())
		if img.Width == 0 {
			img.Width = int(width)
		}
		if img.Height == 0 {
			img.Height = int(height)

		}
	}
	if img.CircleClip {
		imgInstance = img.CircleClipAction(imgInstance)
	} else {
		imgInstance = resize.Resize(uint(img.Width), uint(img.Height), imgInstance, resize.Lanczos3)
	}
	dc.DrawImage(imgInstance, img.X, img.Y)
	return nil
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
