package file

import "github.com/spf13/viper"

type FileSystem interface {
	PrintTest()
	Save(string, string) error
}

func GetFileSystem() (fileSystem FileSystem) {
	switch viper.GetString("disk") {
	case "oss":
		fileSystem = NewOssFile()
		return
	default:
		fileSystem = NewLocalFile()
		return
	}
}
