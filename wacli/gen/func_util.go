package gen

import (
	"regexp"
	"strings"
	"unicode"
)

const (
	funcStartPattern  = `^\s*func\s+`
	funcEndPattern    = `{\s*$`
	annotationPattern = `^\s*//\s*@WinterAOP\(template=([a-zA-Z0-9]+)\)\s*$`

	funcPattern = `^\s*func\s+(\(((?P<obj_name>[a-zA-Z0-9]+)?\s*(?P<pointer>\*)?\s*(?P<struct_name>[a-zA-Z0-9]+))\))?\s*(?P<func_name>[a-zA-Z0-9]+)(?P<parameter_result>.+)\s+{\s*$`
)

var (
	funcStartRegexp = regexp.MustCompile(funcStartPattern)
	funcEndRegexp   = regexp.MustCompile(funcEndPattern)
	funcRegexp      = regexp.MustCompile(funcPattern)
	templateRegexp  = regexp.MustCompile(annotationPattern)
)

func IsFuncStart(code string) bool {
	return funcStartRegexp.MatchString(code)
}

func IsFuncEnd(code string) bool {
	return funcEndRegexp.MatchString(code)
}

func ParseAnnotationStrForTemplate(annotation string) (template string, isFenggoAOP bool) {
	match := templateRegexp.FindStringSubmatch(annotation)
	if len(match) == 2 {
		return match[1], true
	}
	return "", false
}

type FuncStruct struct {
	FuncOrigin string

	ObjectName         string
	Point              string
	StructName         string
	FuncName           string
	ParameterAndResult string

	FuncVariableParameters string

	FuncNameLowCase string

	hasReturn bool
}

func (f *FuncStruct) Init(funcStr string) bool {
	f.FuncOrigin = strings.Replace(funcStr, "\n", " ", -1)
	matchs := funcRegexp.FindStringSubmatch(f.FuncOrigin)
	if len(matchs) < 3 {
		return false
	}
	groupNames := funcRegexp.SubexpNames()
	for i, v := range groupNames {
		switch v {
		case "obj_name":
			f.ObjectName = matchs[i]
		case "pointer":
			f.Point = matchs[i]
		case "struct_name":
			f.StructName = matchs[i]
		case "func_name":
			f.FuncName = matchs[i]
		case "parameter_result":
			f.ParameterAndResult = matchs[i]

		}
	}
	fnRune := []rune(f.FuncName)
	if !unicode.IsUpper(fnRune[0]) {
		return false
	}
	// get lowcase function Name
	f.parseLowcaseFuncName()

	// Parse variable names in parameters
	tmpArgs1 := strings.Split(f.ParameterAndResult, ")")
	if strings.TrimSpace(tmpArgs1[1]) != "" {
		f.hasReturn = true
	}
	tmpArgs2 := strings.Split(tmpArgs1[0], "(")
	tmpargs3 := strings.Split(tmpArgs2[1], ",")
	funcParams := []string{}
	for _, w := range tmpargs3 {
		w = strings.TrimSpace(w)
		fileds := strings.Fields(w)
		if len(fileds) != 0 {
			funcParams = append(funcParams, strings.Fields(w)[0])
		}
	}
	f.FuncVariableParameters = "(" + strings.Join(funcParams, ",") + ")"

	return true
}

func (f *FuncStruct) GenCode(templateName string) string {
	result := ""
	if f.StructName == "" {
		result = result + "func " + f.FuncName + f.ParameterAndResult + " {\n"
	} else {
		result = result + "func (" + f.ObjectName + f.Point + f.StructName + ")" + f.FuncName + f.ParameterAndResult + " {\n"
	}
	result = result + "\t"
	if f.hasReturn {
		result = result + "return "
	}
	result = result + templateName + ".Call("
	if f.ObjectName != "" {
		result = result + f.ObjectName + "."
	}
	result = result + f.FuncNameLowCase + ").(func" + f.ParameterAndResult + ")" + f.FuncVariableParameters + "\n"
	result = result + "}\n"

	if f.StructName == "" {
		result = result + "func " + f.FuncNameLowCase + f.ParameterAndResult + " {\n"
	} else {
		result = result + "func (" + f.ObjectName + f.Point + f.StructName + ")" + f.FuncNameLowCase + f.ParameterAndResult + " {\n"
	}

	return result
}

func (f *FuncStruct) parseLowcaseFuncName() {
	fnRune := []rune(f.FuncName)
	j := 0
	for i, r := range fnRune {
		if unicode.IsLower(r) {
			j = i
			break
		}
	}
	if j == 1 {
		fnRune[0] = unicode.ToLower(fnRune[0])
	} else if j > 1 {
		for i := 0; i < j-1; i++ {
			fnRune[i] = unicode.ToLower(fnRune[i])
		}
	}
	f.FuncNameLowCase = string(fnRune)
}
