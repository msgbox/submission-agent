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
	"fmt"
	"github.com/msgbox/queue"
	"github.com/msgbox/submission-agent/submission_agent"
	"github.com/streadway/amqp"
	"io/ioutil"
	"net"
)

type session struct {
	queue *amqp.Connection
}

func CreateAgent(port string) {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	// Create an AMQP Connection
	queueConn, err := queue.Connect()
	checkError(err)
	defer queueConn.Close()

	s := &session{
		queue: queueConn,
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleMessage(conn, s)
	}
}

func handleMessage(conn net.Conn, s *session) {
	// close connection on exit
	defer conn.Close()

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	err = agent.Send(result, s.queue)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("Fatal error: %s", err.Error())
	}
}
