package gen

import (
	"fmt"
	"github.com/fengya90/winter-aop/util"
	"github.com/fengya90/winter-aop/wacli/common"
)

func GenFunc(sourceDirs []string, configFilePath string) {
	allSourceDirs, err := common.GetCodeDirs(sourceDirs, configFilePath)
	if err != nil {
		fmt.Println("read configfile failed")
		return
	}
	allSourceFile := []string{}
	for _, dir := range allSourceDirs {
		files, err := util.GetFilePathsInTheDirWithSuffix(dir, ".go")
		if err != nil {
			panic(err)
		}
		allSourceFile = append(allSourceFile, files...)
	}
	for _, file := range allSourceFile {
		cf := CodeFile{}
		cf.InitFromFilePath(file)
		if cf.HasAOPAnnotation() {
			fmt.Println("generate " + cf.GetGenFilePath())
			cf.WriteGenFile()
		}
	}
}
