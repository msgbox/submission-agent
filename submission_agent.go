// MsgBox Submission Agent
//
// The Submission Agent is responsible for writing a message
// to an outhoing queue that can eventually be delivered to a recipient.
//
// It uses protocol buffers to pass messages around and should
// encode a message into a protocol buffer before pushing it to
// the queue.

package submission_agent

import (
	"github.com/msgbox/submission-agent/submission_agent"
	"net"
)

func CreateAgent(port string) {
	ln, err := net.Listen("tcp", port)
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

		sess, err := agent.Session(rw)
		if err != nil {
			continue
		}

		go sess.Read()
	}
}
