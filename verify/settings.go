package verify

import (
	"io/ioutil"
	"log"
	"regexp"
	. "strings"
	"time"

	. "strconv"

	"gopkg.in/yaml.v2"
)

type file struct {
	Name    string
	MinSize string `yaml:"min_valid_size"`
}

type path struct {
	Name    string `yaml:"path"`
	Files   []file
	MinSize string `yaml:"min_valid_size,omitempty"`
}

type emailSettings struct {
	Login    string
	Password string
}

type emailList []string

type settings struct {
	EmailSetting emailSettings `yaml:"email_settings"`
	EmailList    emailList     `yaml:"email_list"`
	Paths        []path
}

var Set []PathVerifyInfo
var Emails settings

// InitSet load settings from file to Set.
func InitSet(path string) {
	if path == "" {
		path = "config.yaml"
	}

	f, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var set settings
	err = yaml.Unmarshal(f, &set)

	Emails = set

	if err != nil {
		log.Fatal(err.Error())
	}

	Set = normalizeSetting(set)
}

func normalizeSetting(filesPaths settings) []PathVerifyInfo {
	normalizeDate(filesPaths.Paths)
	normalizeSize(filesPaths.Paths)

	return convertToVerifyInfo(filesPaths.Paths)
}

func convertToVerifyInfo(filePaths []path) []PathVerifyInfo {
	var verifyInfo []PathVerifyInfo

	for _, p := range filePaths {
		pathVerify := PathVerifyInfo{}
		pathVerify.Path = p.Name

		if p.MinSize != "" {
			parseSize, err := ParseInt(p.MinSize, 10, 64)
			if err != nil {
				log.Fatalf("Size not parse: %v\n", p.MinSize)
				panic(err)
			}

			pathVerify.ValidSize = parseSize
		}

		for _, f := range p.Files {
			fileVerify := FileVerifyInfo{}
			fileVerify.FilePattern = f.Name

			parseSize, err := ParseInt(f.MinSize, 10, 64)
			if err != nil {
				log.Fatalf("Size not parsed: %v\n", parseSize)
				panic(err)
			}

			fileVerify.Size = parseSize

			pathVerify.FileVerifyInfos = append(pathVerify.FileVerifyInfos, fileVerify)
		}

		verifyInfo = append(verifyInfo, pathVerify)
	}

	return verifyInfo
}

func normalizeSize(filesPaths []path) {
	reg := regexp.MustCompile(`(\d+) ([KMG]B)`)

	for i := range filesPaths {
		filesPaths[i].MinSize = reg.ReplaceAllStringFunc(filesPaths[i].MinSize, insertSize)

		for j := range filesPaths[i].Files {
			filesPaths[i].Files[j].MinSize = reg.ReplaceAllStringFunc(
				filesPaths[i].Files[j].MinSize,
				insertSize)
		}
	}
}

func insertSize(str string) string {
	size, err := Atoi(str[:Index(str, "B")-2])
	if err != nil {
		return "ERROR"
	}

	label := str[Index(str, "B")-1:]

	switch label {
	case "KB":
		size *= 1024
	case "MB":
		size *= 1024 * 1024
	case "GB":
		size *= 1024 * 1024 * 1024
	}

	return Itoa(size)
}

func normalizeDate(filesPaths []path) {
	reg := regexp.MustCompile(`%\[(-?\d*)?([ymdYMD])\]`)

	for i := range filesPaths {
		currentTime := time.Now()

		filesPaths[i].Name = reg.ReplaceAllStringFunc(
			filesPaths[i].Name,
			func(m string) string {
				isNeg := Index(m, "-") > -1

				lastScopeInd := LastIndex(m, "]")
				label := m[lastScopeInd-1 : lastScopeInd]

				var shiftTmp string
				if isNeg {
					shiftTmp = m[2 : lastScopeInd-1]
				} else {
					shiftTmp = m[1 : lastScopeInd-1]
				}

				if shiftTmp == "[" {
					shiftTmp = "0"
				}

				shift, err := Atoi(shiftTmp)
				if err != nil {
					panic(err)
				}

				switch ToLower(label) {
				case "y":
					currentTime = currentTime.AddDate(shift, 0, 0)
				case "m":
					currentTime = currentTime.AddDate(0, shift, 0)
				case "d":
					currentTime = currentTime.AddDate(0, 0, shift)
				}

				return Join([]string{"%", label, "%"}, "")
			})

		filesPaths[i].Name = insertDate(filesPaths[i].Name, currentTime)
	}
}

func insertDate(str string, time time.Time) string {
	str = Replace(str, "%y%", Itoa(time.Year()), -1)
	str = Replace(str, "%Y%", Itoa(time.Year()), -1)

	upperTime := Split(time.Format("01@02"), "@")
	lowerTime := Split(time.Format("1@2"), "@")

	str = Replace(str, "%m%", lowerTime[0], -1)
	str = Replace(str, "%M%", upperTime[0], -1)

	str = Replace(str, "%d%", lowerTime[1], -1)
	str = Replace(str, "%D%", upperTime[1], -1)

	return str
}
