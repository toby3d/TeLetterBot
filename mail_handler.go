package main

import (
	"bytes"
	"log"
	"net"

	"github.com/jhillyerd/enmime"
)

func mailHandler(origin net.Addr, from string, to []string, data []byte) {
	msg, err := enmime.ReadEnvelope(bytes.NewReader(data))
	checkError(err)

	subject := msg.GetHeader("Subject")
	log.Printf("Received mail from %s for %s with subject %s", from, to[0], subject)
}
