# Networking Projects with Go

This repository contains a collection of networking projects implemented in Go.
## Table of Contents

- [Introduction](#introduction)
- [Projects](#projects)
- [Getting Started](#getting-started)

## Introduction

As a software developer learning networking concepts, I've created this repository to document my journey of learning both networking theory and Go programming. The projects included here aim to cover a variety of networking topics and provide practical demonstrations of how these concepts can be implemented using Go.

Feel free to explore the projects, use them as learning resources, and even contribute your own projects to expand the collection.

## Projects

1. [UDP Client-Server Communication](/UdpClientServer) - Demonstrates a simple UDP client-server communication setup, where the server listens for incoming UDP messages on localhost and converts received text to uppercase and returns it to the client as a response, while the client sends user-input text to the server using UDP.

2. [TCP Client-Server Communication](/TcpClientServer) - Demonstrate a simple TCP client-server communication setup. The server listens for incoming connections and echoes received messages. The client connects to the server and allows users to send messages to it. Communication continues until either party sends "exit".

3. [Web Server Implementation](/WebServer) - Demonstrates a TCP server that listens on localhost, processes incoming requests to retrieve and send files based on HTTP-like responses, demonstrating basic web server-client interaction using Go's net package and custom helper libraries.

4. [Udp Pinger](/UdpPinger) - The client sends "Ping" messages to the UDP server on localhost and measures the round-trip time for a response and has a timeout for 1 sec. The server listens for incoming UDP messages on the localhost and responds with "Pong" messages.
## Getting Started

Each project is contained within its own directory.

Clone this repository:

```sh
git clone https://github.com/aayushbhan/awesomeProjectNetworks.git
cd awesomeProjectNetworks
