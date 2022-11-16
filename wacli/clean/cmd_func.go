package clean

import (
	"fmt"
	"github.com/fengya90/winter-aop/util"
	"github.com/fengya90/winter-aop/wacli/common"
	"os"
)

func CleanFunc(sourceDirs []string, configFilePath string) {
	allSourceDirs, err := common.GetCodeDirs(sourceDirs, configFilePath)
	if err != nil {
		fmt.Println("read configfile failed")
		return
	}
	allSourceFile := []string{}
	for _, dir := range allSourceDirs {
		files, err := util.GetFilePathsInTheDirWithSuffix(dir, common.GenFileSuffix)
		if err != nil {
			panic(err)
		}
		allSourceFile = append(allSourceFile, files...)
	}
	for _, file := range allSourceFile {
		os.Remove(file)
	}
}
