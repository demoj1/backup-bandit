package robber

import (
	"os/exec"
	"regexp"
)

// Tool ...
type Tool struct {
	Path   string
	Args   string
	Groups []string
	Regex  string
}

// ParseOut run tool and parse std out
// to compliance with regular expession.
func (t *Tool) ParseOut() (m [][]string, err error) {
	o, e := exec.Command(t.Path, t.Args).Output()
	if e != nil {
		m = nil
		err = e
		return
	}

	out := string(o)

	reg, e := regexp.Compile(t.Regex)
	if e != nil {
		m = nil
		err = e
		return
	}

	findAll := reg.FindAllStringSubmatch(out, -1)

	defer func() {
		if r := recover(); r != nil {
			m = nil
			err = r.(error)
			return
		}
	}()

	var res [][]string
	for i, f := range findAll {
		res = append(res, make([]string, len(t.Groups)))

		for j := range t.Groups {
			res[i][j] = f[j+1]
		}
	}

	m = res
	err = nil
	return
}
