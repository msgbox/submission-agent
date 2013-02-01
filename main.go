// MsgBox Submission Agent
//
// The Submission Agent is responsible for writing a message
// to an outhoing queue that can eventually be delivered to a recipient.
//
// It uses protocol buffers to pass messages around and should
// encode a message into a protocol buffer before pushing it to
// the queue.

package main

import (
	"github.com/msgbox/submission-agent/submission-agent"
	"net"
)

var Addr string

func main() {

	addr := Addr
	if addr == "" {
		addr = ":1337"
	}

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		// handle error
	}
	defer ln.Close()

	for {
		rw, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}

		sess, err := submission_agent.Session(rw)
		if err != nil {
			continue
		}

		go sess.Read()
	}

}
