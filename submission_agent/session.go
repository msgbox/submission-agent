package agent

import (
	"bufio"
	"github.com/msgbox/queue"
	"log"
	"net"
)

type session struct {
	rwc net.Conn
	br  *bufio.Reader
	bw  *bufio.Writer
}

func Session(rwc net.Conn) (s *session, err error) {
	s = &session{
		rwc: rwc,
		br:  bufio.NewReader(rwc),
		bw:  bufio.NewWriter(rwc),
	}

	return
}

// Read from the Session buffer and send to a handler function
func (s *session) Read() {
	defer s.rwc.Close()

	for {
		sl, err := s.br.ReadSlice('\n')
		if err != nil {
			log.Printf("Read error: %v", err)
			return
		}

		s.handleMessage(sl)
	}
}

// Connect to an AMQP Exchange and write the message
func (s *session) handleMessage(body []byte) {
	// Create an AMQP Connection
	conn, c_err := queue.Connect()
	if c_err != nil {
		// Handle Error
	}
	defer conn.Close()

	err := Send(body, conn)
	if err != nil {
		// Handle error
	}
}
