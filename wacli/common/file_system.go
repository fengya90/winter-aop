package common

import (
	"github.com/fengya90/winter-aop/util"
	"path"

	"github.com/spf13/viper"
)

const (
	GenFileSuffix = "_winter_aop_gen.go"
	sourceDirList = "source_dir_list"
)

func GetCodeDirs(sourceDirs []string, configFilePath string) (allSourceDirs []string, err error) {
	allSourceDirs = sourceDirs
	if configFilePath != "" {
		vp := viper.New()
		vp.SetConfigFile(configFilePath)
		err = vp.ReadInConfig()
		if err != nil {
			return
		}
		dirList := vp.GetStringSlice(sourceDirList)
		for _, dir := range dirList {
			allSourceDirs = append(allSourceDirs, path.Join(path.Dir(configFilePath), dir))
		}
	}

	allSourceDirs = util.RemoveDuplicateStr(allSourceDirs)
	return allSourceDirs, nil
}
