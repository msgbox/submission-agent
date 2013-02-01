package submission_agent

import (
	"code.google.com/p/goprotobuf/proto"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/msgbox/message"
	"github.com/msgbox/queue"
	"github.com/streadway/amqp"
	"time"
)

type Item struct {
	Header  Header
	Payload Payload
}

type Header struct {
	Creator    string
	Receiver   string
	Created_At int64
	MessageID  string
}

type Payload struct {
	Body string
}

func (m *Item) setCreatedAt() int64 {
	t := time.Now().UTC().Unix()
	return *&t
}

// Generate a UUID to use as a unique Message ID
// Need to work on the UUID library to ensure UUID's
// are created correctly
func (m *Item) generateUUID() string {
	uuid, err := genUUID()
	if err != nil {
		fmt.Errorf("Error Generating UUID: %s", err)
	}

	return *&uuid
}

// Accepts an input source and attempts to marshall it into a
// protocol buffer, then pushes the message to an AMQP Exchange.
//
// Currently accepts a JSON input but this should change.
func Send(data []byte, connection *amqp.Connection) error {
	pb, err := createProtocolBuffer(data)
	if err != nil {
		return fmt.Errorf("Protocol Buffer Error: %s", err)
	}

	// Send pb to AMQP Exchange
	p_err := queue.Publish("outgoing", pb, connection)
	if p_err != nil {
		return fmt.Errorf("Publishing Error: %s", p_err)
	}

	return nil
}

// Takes a byte slice and converts it into a Protocol Buffer.
//
// Returns the protocol buffer and error
//
// Currently parses JSON input but this should change
func createProtocolBuffer(data []byte) ([]byte, error) {

	var i Item
	err := json.Unmarshal(data, &i)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling error:", err)
	}

	// Now we want to create a Message from the Item
	msg := &messages.Message{
		Creator:   proto.String(*&i.Header.Creator),
		Receiver:  proto.String(*&i.Header.Receiver),
		CreatedAt: proto.Int64(i.setCreatedAt()),
		Id:        proto.String(i.generateUUID()),
		Payload:   proto.String(*&i.Payload.Body),
	}

	p, err := proto.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("marshaling error: ", err)
	}

	return p, nil
}

// UUID v4 Generator
// http://www.ashishbanerjee.com/home/go/go-generate-uuid
func genUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// TODO: verify the two lines implement RFC 4122 correctly
	uuid[8] = 0x80 // variant bits see page 5
	uuid[4] = 0x40 // version 4 Pseudo Random, see page 7

	return hex.EncodeToString(uuid), nil
}
