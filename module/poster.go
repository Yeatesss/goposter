package module

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"os"

	"github.com/Yeate/goposter/file"
	"github.com/Yeate/gowheel"
	"github.com/fogleman/gg"
	"github.com/spf13/viper"
)

type Poster struct {
	Width           float64     `json:"width"`
	Height          float64     `json:"height"`
	BackgroundColor string      `json:"backgroundColor"`
	Background      interface{} `json:"background"`
	Texts           []Text      `json:"texts"`
	Images          []Image     `json:"images"`
	Lines           []Line      `json:"lines"`
	SavePath        string
	SaveName        string
}

var FileSystem file.FileSystem
var ImgTempDir string

//func init() {
//	fontBytes, _ := ioutil.ReadFile(viper.GetString("font.path"))
//	Font, _ = truetype.Parse(fontBytes)
//
//}

func NewPoster() *Poster {
	viper.SetDefault("img_tmp_dir", "tmp/image/")
	FileSystem = file.GetFileSystem()
	ImgTempDir = viper.GetString("img_tmp_dir")
	return &Poster{}
}

//NewPoster 初始化画布
func (poster *Poster) NewPoster() (ins *Instantiation) {
	_ = gowheel.InitPath(ImgTempDir)
	ins = &Instantiation{}
	if poster.Background != nil {
		switch v := poster.Background.(type) {
		case string:
			background := v
			poster.Height, poster.Width = gowheel.GetImageSizeFromUrl(background)
			poster.Images = append([]Image{{X: 0, Y: 0, Url: background, Width: int(poster.Width), Height: int(poster.Height)}}, poster.Images...)
		case io.Reader:

			bs, _ := ioutil.ReadAll(v)
			ins, _, _ := image.Decode(bytes.NewBuffer(bs))
			g := ins.Bounds()
			height := float64(g.Dy())
			width := float64(g.Dx())
			if poster.Width == 0 {
				poster.Width = width
			}
			if poster.Height == 0 {
				poster.Height = height
			}
			poster.Images = append([]Image{{X: 0, Y: 0, ImgData: bytes.NewBuffer(bs), Width: int(poster.Width), Height: int(poster.Height)}}, poster.Images...)

		}

	}
	ins.Background = gg.NewContext(int(poster.Width), int(poster.Height))
	ins.Background.SetHexColor(poster.BackgroundColor)
	ins.Background.DrawRectangle(0, 0, poster.Width, poster.Height)
	ins.Background.Fill()
	return
}

func (poster *Poster) Draw() (err error) {
	if poster.SavePath == "" {
		err = fmt.Errorf("save_path is null")
		return
	}
	//初始画布，设置背景
	ins := poster.NewPoster()
	//向画布上添加元素
	for _, drawable := range poster.Images {
		err = ins.draw(&drawable)
		if err != nil {
			return
		}
	}
	for _, drawable := range poster.Texts {
		err = ins.draw(&drawable)
		if err != nil {
			return
		}
	}
	for _, drawable := range poster.Lines {
		err = ins.draw(&drawable)
		if err != nil {
			return
		}
	}
	//保存图片
	err = ins.save(poster.SavePath, poster.SaveName)

	if err != nil {
		return err
	}

	return err
}

type Instantiation struct {
	Background *gg.Context
}

func (ins *Instantiation) draw(element Drawable) (err error) {
	err = element.Draw(ins.Background)
	return
}

func (ins *Instantiation) save(path, name string) (err error) {
	defer func() {
		_ = os.Remove(ImgTempDir + name)
	}()
	err = gowheel.InitPath(ImgTempDir)

	err = ins.Background.SavePNG(ImgTempDir + name)
	if err == nil {
		err = FileSystem.Save(path, name)
	}
	return
}
