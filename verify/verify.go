package verify

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var (
	sizeSum   int64
	pathRegex *regexp.Regexp
)

// PathVerifyInfo ...
type PathVerifyInfo struct {
	Path            string           // Root directory
	ValidSize       int64            // Minimal valid size
	FileVerifyInfos []FileVerifyInfo // Files in dir
}

// FileVerifyInfo ...
type FileVerifyInfo struct {
	FilePattern string // File pattern
	Size        int64  // Minimal valid size
	Info        error  // Meta field
	regex       *regexp.Regexp
}

type tmpFile struct {
	Pattern string
	File    os.FileInfo
	Size    int64
}

var (
	size    int64
	files   []os.FileInfo
	logFile map[string]string
)

const (
	CriticalLabel string = "[CRITICAL]"
	ErrorLabel    string = "[ERROR]"
	WarningLabel  string = "[WARNING]"
	OkLabel       string = "[OK]"
)

// Verify file/path size.
func Verify(root PathVerifyInfo) (log map[string]string) {
	size = 0
	logFile = make(map[string]string, 0)

	defer func() {
		files = files[:0]
		logFile = make(map[string]string, 0)

		if r := recover(); r != nil {
			fmt.Println(r)

			logFile[r.(error).Error()] = logFormat(
				CriticalLabel,
				"%v",
				r.(error).Error(),
			)

			log = logFile
			return
		}
	}()

	_, ferr := os.Stat(root.Path)

	if os.IsNotExist(ferr) {
		logFile[root.Path] = logFormat(
			ErrorLabel,
			"path: %v error: %v",
			root.Path,
			ferr.Error(),
		)

		log = logFile
		return
	}

	if root.FileVerifyInfos == nil {
		err := filepath.Walk(root.Path, walkFunc)

		if err != nil {
			logFile[root.Path] = logFormat(
				ErrorLabel,
				"path: %v error: %v",
				root.Path,
				err.Error(),
			)
		}

		if size >= root.ValidSize {
			logFile[root.Path] = logFormat(
				OkLabel,
				"path: %v valid size: %v current size: %v",
				root.Path,
				root.ValidSize,
				size,
			)
		} else {
			logFile[root.Path] = logFormat(
				WarningLabel,
				"path: %v valid size: %v current size: %v",
				root.Path,
				root.ValidSize,
				size,
			)
		}
	}

	for i, v := range root.FileVerifyInfos {
		root.FileVerifyInfos[i].regex = regexp.MustCompile(v.FilePattern)
	}

	err := filepath.Walk(root.Path, getAllFiles)
	if err != nil {
		logFile[root.Path] = logFormat(
			ErrorLabel,
			"path: %v error: %v",
			root.Path,
			err.Error(),
		)
	}

	for _, v := range root.FileVerifyInfos {
		var tmp []tmpFile

		for _, f := range files {
			if !v.regex.MatchString(f.Name()) {
				continue
			}

			tmp = append(tmp, tmpFile{
				Pattern: v.FilePattern,
				File:    f,
				Size:    v.Size,
			})
		}

		if len(tmp) == 0 {
			logFile[v.FilePattern] = logFormat(
				ErrorLabel,
				"files suitable for the pattern: %v is not found.",
				v.FilePattern,
			)
		} else {
			for _, vf := range tmp {
				if vf.File.Size() < vf.Size {
					logFile[vf.File.Name()] = logFormat(
						WarningLabel,
						"path: %v valid size: %v current size: %v",
						root.Path+"/"+vf.File.Name(),
						vf.Size,
						vf.File.Size(),
					)
				} else {
					logFile[vf.File.Name()] = logFormat(
						OkLabel,
						"path: %v valid size: %v current size: %v",
						root.Path+"/"+vf.File.Name(),
						vf.Size,
						vf.File.Size(),
					)
				}
			}
		}

		tmp = tmp[:0]
	}

	return logFile
}

func logFormat(label string, format string, args ...interface{}) string {
	return label + " " + fmt.Sprintf(format, args...)
}

func getAllFiles(path string, info os.FileInfo, e error) error {
	if info == nil {
		return nil
	}

	if info.IsDir() {
		return nil
	}

	files = append(files, info)
	return nil
}

func walkFunc(path string, info os.FileInfo, e error) (err error) {
	if e != nil {
		err = e
		return
	}

	if info == nil {
		err = nil
		return
	}

	size += info.Size()

	err = nil
	return
}
