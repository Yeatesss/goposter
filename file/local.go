package file

import (
	"fmt"

	"github.com/Yeate/gowheel"

	"github.com/spf13/viper"
)

type LocalFile struct {
}

func NewLocalFile() *LocalFile {
	return &LocalFile{}
}
func (l *LocalFile) PrintTest() {
	fmt.Println("local")
}
func (l *LocalFile) Save(savePath, fileName string) (err error) {
	_, err = gowheel.CopyFile(savePath+fileName, viper.GetString("img_tmp_dir")+fileName)
	return

}
