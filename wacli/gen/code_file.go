package gen

import (
	"github.com/fengya90/winter-aop/wacli/common"
	"io/ioutil"
	"regexp"
	"strings"
)

const (
	buildFlag = `^\s*//\s*(\+build|go:build)`
)

var (
	buildFlagRegexP = regexp.MustCompile(buildFlag)
)

type CodeFile struct {
	filePath string
	lines    []string
}

func (cf *CodeFile) InitFromFilePath(filePath string) error {
	cf.filePath = filePath
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	content := string(bytes)
	cf.lines = strings.Split(content, "\n")
	return nil
}

func (cf *CodeFile) HasAOPAnnotation() bool {
	for _, line := range cf.lines {
		_, aop := ParseAnnotationStrForTemplate(line)
		if aop {
			return true
		}
	}
	return false
}

func (cf *CodeFile) WriteGenFile() error {
	newFilePath := cf.filePath[:len(cf.filePath)-3] + common.GenFileSuffix
	return ioutil.WriteFile(newFilePath, []byte(cf.generateProxyCode()), 0644)
}

func (cf *CodeFile) GetGenFilePath() string {
	newFilePath := cf.filePath[:len(cf.filePath)-3] + common.GenFileSuffix
	return newFilePath
}

func (cf *CodeFile) generateProxyCode() string {
	result := `//go:build winter_aop
// +build winter_aop
`

	template := ""
	funcStr := ""
	stage := 0
	for _, line := range cf.lines {
		if buildFlagRegexP.MatchString(line) {
			continue
		}

		tmpTemplate, isAop := ParseAnnotationStrForTemplate(line)
		if isAop {
			template = tmpTemplate
			continue
		}
		switch stage {
		case 0: // normal
			if IsFuncStart(line) && template != "" {
				funcStr = funcStr + line
				stage = 1
				if IsFuncEnd(line) {
					fs := &FuncStruct{}
					fs.Init(funcStr)
					result = result + fs.GenCode(template)
					funcStr = ""
					stage = 0
					template = ""
				}
			} else {
				result = result + line + "\n"
			}
		case 1: // read func
			if IsFuncEnd(line) {
				funcStr = funcStr + line
				fs := &FuncStruct{}
				fs.Init(funcStr)
				result = result + fs.GenCode(template)
				funcStr = ""
				stage = 0
				template = ""
			} else {
				funcStr = funcStr + line
			}

		}
	}
	return result
}
