MsgBox Submission Agent
=================================

The Submission Agent is responsible for writing a message to a queue that can eventually be delivered to a recipient.

It listens on a specified port for incoming connections and encodes a message that is written to an AMQP queue.

It uses protocol buffers to pass messages around and should encode a message into a protocol buffer before pushing it to the queue.

## Notes:

Should outgoing messages require a queue? In a distibuted system workers could be on seperate machines talking to seperate relays to allow more throughput.

This could be making it more complex with little gain though depending on how many connections the relay can handle at once. The queue for outgoing messages will probally be stripped out and the submission agent will just dial the relay and pass along the encoded message. This makes the system much easier to wrap your head around and prevents unnecessary complexity.