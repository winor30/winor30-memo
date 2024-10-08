#!/bin/bash

# Start the server in the background
echo "Starting the server in the background..."
go run greeter_server/main.go &

# Wait for the server to start
echo "Waiting for the server to start..."
sleep 2

# Run the client
echo "Running the client..."
go run greeter_client/main.go

# Stop the server
echo "Stopping the server..."
kill -9 $(lsof -t -i:50051)
