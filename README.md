## Cli-Chat Client

A lightweight command-line chat client written in Go. This project provides a simple and efficient way to communicate in real-time via a terminal-based interface.

## Features
- Real-time chat functionality.
- Lightweight and fast.
- Simple and intuitive command-line interface.
- Written in Go, ensuring high performance and portability.

## How to run
### Build Binary file
#### Prerequisites
before install or build this project you need to download following package
- [Go Lang](https://go.dev/)

```sh
git clone https://github.com/ThawThuHan/cli-chat-client.git
cd cli-chat-client
go build -o cli-chat main.go
./cli-chat <server-ip> <server-port>
```

### Run from Pre-Build binary file
#### Windows
Download pre-build binary file [Cli-Chat-Client](https://github.com/ThawThuHan/cli-chat-client/releases/download/v1.0/cli-chat.exe) from release and run
```sh
./cli-chat.exe <server-ip> <server-port>
```
OR
```sh
curl -o cli-chat.exe https://github.com/ThawThuHan/cli-chat-client/releases/download/v1.0/cli-chat.exe
./cli-chat.exe <server-ip> <server-port>
```