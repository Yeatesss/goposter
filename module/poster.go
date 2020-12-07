package module

import (
	"fmt"
	"os"

	"github.com/Yeate/goposter/file"
	"github.com/Yeate/gowheel"
	"github.com/fogleman/gg"
	"github.com/spf13/viper"
)

type Poster struct {
	Width           float64 `json:"width"`
	Height          float64 `json:"height"`
	BackgroundColor string  `json:"backgroundColor"`
	Background      string  `json:"background"`
	Texts           []Text  `json:"texts"`
	Images          []Image `json:"images"`
	Lines           []Line  `json:"lines"`
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
	if poster.Background != "" {
		poster.Height, poster.Width = gowheel.GetImageSizeFromUrl(poster.Background)
		poster.Images = append([]Image{{0, 0, poster.Background, int(poster.Width), int(poster.Height), 0, false}}, poster.Images...)
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
	err = ins.Background.SavePNG(ImgTempDir + name)
	if err == nil {
		err = FileSystem.Save(path, name)

	}
	return
}

//
//func (poster *Poster) initQueue() {
//	var queue DrawQueue
//	sort := []int{}
//
//	if poster.queue == nil {
//		queue = make(DrawQueue)
//	} else {
//		queue = poster.queue
//	}
//
//	for _, drawable := range poster.Images {
//		queue, sort = appendToQueue(queue, drawable, sort)
//	}
//
//	for _, drawable := range poster.Lines {
//		queue, sort = appendToQueue(queue, drawable, sort)
//	}
//
//	for _, drawable := range poster.Blocks {
//		queue, sort = appendToQueue(queue, drawable, sort)
//	}
//
//	for _, drawable := range poster.Texts {
//		queue, sort = appendToQueue(queue, drawable, sort)
//	}
//
//	poster.sort = sort
//	//utils2.QuickSort(poster.sort)
//	poster.queue = queue
//}
//
//func appendToQueue(queue DrawQueue, drawable Drawable, sort []int) (DrawQueue, []int) {
//	index := drawable.GetZIndex()
//	value, exists := queue[index]
//	if exists {
//		queue[index] = append(value, drawable)
//	} else {
//		sort = append(sort, index)
//		queue[index] = []Drawable{drawable}
//	}
//
//	return queue, sort
//}
