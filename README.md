# Testing grpc backpressure

The following repository contains a simple example of how to test the backpressure of a grpc server.
The server is a simple grpc that will indefinitely stream messages to the client. The client will receive the messages 
and process them slowly. The server will keep sending messages to the client until the client closes the connection.

One can witness the backpressure in action by running the server and the client and observing that the server will stop 
sending messages (logged to stdout) to the client when the client is not able to process the messages fast enough.