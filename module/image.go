package module

import (
	"os"

	"github.com/Yeate/gowheel"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
	"github.com/qbhy/go-utils"
)

type Image struct {
	X            int    `json:"x"`
	Y            int    `json:"y"`
	Url          string `json:"url"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	BorderRadius int    `json:"borderRadius"`
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
		img.CheckWh()
		imgInstance = resize.Resize(uint(img.Width), uint(img.Height), imgInstance, resize.Lanczos3)
		dc.DrawImage(imgInstance, img.X, img.Y)
		return nil
	} else {
		return err
	}
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
