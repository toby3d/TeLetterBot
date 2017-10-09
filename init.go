package main

import (
	"flag"

	"github.com/mhale/smtpd"
)

var flagToken = flag.String("token", "", "set telegram token")

func init() {
	flag.Parse()

	go func() {
		err := smtpd.ListenAndServe("127.0.0.1:2525", mailHandler, "TeLetterBot", "")
		checkError(err)
	}()
}
