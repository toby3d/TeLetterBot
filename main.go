package main

import "github.com/mhale/smtpd"

func main() {
	err := smtpd.ListenAndServe("127.0.0.1:2525", mailHandler, "TeLetterBot", "")
	checkError(err)
}
