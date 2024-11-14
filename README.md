RabbitMQMessageQueue

This project aims to facilitate interaction with RabbitMQ message queues using the Go programming language. RabbitMQ is an open-source message broker that enables communication between components in distributed systems.

Features

Message Sending and Receiving: Supports sending and receiving messages via RabbitMQ.

Flexible Structure: Designed to adapt to various messaging scenarios.

Modular Coding: Organized in a modular manner for extensibility and maintainability.

Installation

Clone the repository to your local machine:

git clone https://github.com/barisaydogdu/RabbitMQMessageQueue.git

Navigate to the project directory:

cd RabbitMQMessageQueue

Download the necessary dependencies:

go mod tidy

Usage

To perform message sending and receiving operations, follow these steps:

Start the RabbitMQ Server: Ensure that RabbitMQ is running on your local machine.

Sending Messages: Initiate the message-sending process with the following command:

go run cmd/send/main.go

Receiving Messages: In another terminal, start the message-receiving process with:

go run cmd/receive/main.go

These steps will enable you to send and receive messages via RabbitMQ.

Project Structure

The project consists of the following directories and files:

cmd/: Contains the entry points of the application.

send/: Main file for the message-sending process.

receive/: Main file for the message-receiving process.

internal/messaging/: Internal logic and functions related to messaging.

pkg/rabbitMQ/: Helper packages and functions related to RabbitMQ.

util/: Utility functions and tools.

go.mod and go.sum: Go modules and dependency management files.

Contribution

If you wish to contribute, please follow these steps:

Fork this repository.

Create a new branch:

git checkout -b feature-name

Make your changes and commit them:

git commit -m 'Added new feature'

Push your branch:

git push origin feature-name

Create a Pull Request.

Thank you in advance for your contributions!

