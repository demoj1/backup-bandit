package main

import (
	"flag"
	"github.com/diman3241/backupbandit/reporter"
)

func main() {
	var verifyPath = flag.String("verify", "", "Path to verify config file.")
	var robberPath = flag.String("robber", "", "Path to robber config file.")

	flag.Parse()

	reporter.CreateHtmlReport(*verifyPath, *robberPath)
}
