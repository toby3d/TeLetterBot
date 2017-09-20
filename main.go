package main

import (
	"bytes"
	"log"
	"net"

	enmime "github.com/jhillyerd/enmime"
	smtpd "github.com/mhale/smtpd"
)

func hardCheck(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func mailHandler(origin net.Addr, from string, to []string, data []byte) {
	msg, err := enmime.ReadEnvelope(bytes.NewReader(data))
	hardCheck(err)

	subject := msg.GetHeader("Subject")
	log.Printf("Received mail from %s for %s with subject %s", from, to[0], subject)
}

func main() {
	err := smtpd.ListenAndServe("127.0.0.1:2525", mailHandler, "TeLetterBot", "")
	hardCheck(err)
}
