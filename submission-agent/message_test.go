package submission_agent

import (
	"code.google.com/p/goprotobuf/proto"
	"encoding/json"
	"fmt"
	"github.com/msgbox/message"
	"github.com/msgbox/queue"
	"github.com/streadway/amqp"
	"testing"
)

// Helper for creating an AMQP Connection
func createConnection() *amqp.Connection {
	conn, err := queue.Connect()
	if err != nil {
		fmt.Errorf("Error Connecting: %s", err)
	}

	return conn
}

// Helper for creating fake JSON data
// to use in the tests
func make_json() []byte {
	i := &Item{Header{}, Payload{}}

	i.Header.Creator = "sender:home@example.com"
	i.Header.Receiver = "particlebanana:home@example.com"
	i.Payload.Body = "Test Message Body"

	data, _ := json.Marshal(i)
	return data
}

// Ensure Protocol Buffers are Marshaled correctly
func Test_createPB_1(t *testing.T) {
	data := make_json()
	msg, _ := createProtocolBuffer(data)
	newTest := &messages.Message{}
	proto.Unmarshal(msg, newTest)
	if newTest.GetCreator() != "sender:home@example.com" {
		t.Error("Protocol Buffer Not Correct")
	}
}

// Test the message is sent to an AMQP Exchange
func Test_Send_1(t *testing.T) {
	conn := createConnection()
	defer conn.Close()

	data := make_json()

	err := Send(data, conn)
	if err != nil {
		t.Errorf("Send did not work as expected: %s", err)
	}
}
