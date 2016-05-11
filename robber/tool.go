package robber

import (
	"bytes"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

// Tool ...
type Tool struct {
	Path   string
	Args   string
	Groups []string
	Regex  string
	Time   int `yaml:"wait_time,omitempty"`
}

// ParseOut run tool and parse std out
// to compliance with regular expession.
func (t *Tool) ParseOut() (m [][]string, err error) {
	if t.Time == 0 {
		t.Time = 100
	}

	c := exec.Command(t.Path, strings.Split(t.Args, " ")...)

	var out bytes.Buffer
	c.Stdout = &out

	e := c.Start()

	waitTimer := time.NewTimer(time.Millisecond * time.Duration(t.Time))
	go func(timer *time.Timer, c *exec.Cmd) {
		<-timer.C
		err := c.Process.Signal(os.Kill)

		if err == nil {
			timer.Stop()
		}
	}(waitTimer, c)

	c.Wait()
	waitTimer.Stop()

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
