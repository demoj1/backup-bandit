package reporter

import (
	"crypto/tls"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"gopkg.in/gomail.v2"

	"github.com/diman3241/backupbandit/robber"
	"github.com/diman3241/backupbandit/verify"

	tpl "github.com/diman3241/backupbandit/template"
)

func CreateHtmlReport() {
	tpl, err := template.New("report").Parse(tpl.Report)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("/tmp/report.html")
	if err != nil {
		log.Fatal(err)
	}

	allLog := make(map[string]string)

	verify.InitSet("")
	for _, v := range verify.Set {
		tmpMap := verify.Verify(v)

		for k, mapValue := range tmpMap {
			allLog[mapValue] = k
		}
	}

	robber.InitSet("")

	err = tpl.Execute(f, struct {
		Date     string
		Critical []string
		Error    []string
		Warning  []string
		Success  []string
		Tools    []robber.Tool
	}{
		Date:     time.Now().Format("2006-01-02 15:04"),
		Critical: critical(allLog),
		Error:    error(allLog),
		Warning:  warning(allLog),
		Success:  ok(allLog),
		Tools:    robber.Set.Tools,
	})

	f.Close()

	if err != nil {
		log.Fatal(err)
	}

	title := fmt.Sprintf(
		"%v C:%v/E:%v/W:%v/S:%v",
		time.Now().Format("2006-01-02 15:04"),
		len(critical(allLog))-1,
		len(error(allLog))-1,
		len(warning(allLog))-1,
		len(ok(allLog))-1,
	)

	reportText, err := ioutil.ReadFile("/tmp/report.html")
	if err != nil {
		log.Fatal(err)
	}

	sendEmail(reportText, verify.Emails.EmailList, title)
}

func sendEmail(msg []byte, to []string, title string) {
	m := gomail.NewMessage()
	m.SetHeader("From", verify.Emails.EmailSetting.Login)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", title)
	m.SetBody("text/html", string(msg))

	d := gomail.NewDialer(
		verify.Emails.EmailSetting.SMTPServer,
		verify.Emails.EmailSetting.SMTPPort,
		verify.Emails.EmailSetting.Login,
		verify.Emails.EmailSetting.Password,
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}
}

func critical(m map[string]string) []string {
	return filterMap(m, func(k string, v string) bool {
		return strings.HasPrefix(k, verify.CriticalLabel)
	})
}

func error(m map[string]string) []string {
	return filterMap(m, func(k string, v string) bool {
		return strings.HasPrefix(k, verify.ErrorLabel)
	})
}

func warning(m map[string]string) []string {
	return filterMap(m, func(k string, v string) bool {
		return strings.HasPrefix(k, verify.WarningLabel)
	})
}

func ok(m map[string]string) []string {
	return filterMap(m, func(k string, v string) bool {
		return strings.HasPrefix(k, verify.OkLabel)
	})
}

func filterMap(m map[string]string, f func(string, string) bool) []string {
	res := make([]string, 1)

	for k, v := range m {
		if f(k, v) {
			res = append(res, k)
		}
	}

	return res
}
