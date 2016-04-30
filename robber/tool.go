package robber

import (
	"bytes"
	"os/exec"
	"regexp"
	"strings"
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
	c := exec.Command(t.Path, strings.Split(t.Args, " ")...)
	var out bytes.Buffer
	c.Stdout = &out
	e := c.Run()
	c.Wait()

	reg := regexp.MustCompile(t.Regex)
	findAll := reg.FindAllStringSubmatch(out.String(), -1)

	var res [][]string
	for i, f := range findAll {
		res = append(res, make([]string, len(t.Groups)))

		for j := range t.Groups {
			res[i][j] = f[j+1]
		}
	}

	m = res

	if e != nil {
		m = append(m, []string{"std err: " + e.Error()})
	}

	err = nil
	return
}
