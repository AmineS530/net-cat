# Net-Cat Server-client

## Overview

Net-Cat is a Go-based implementation of a chat application mimicking the functionality of the NetCat command-line utility. It operates using a server-client architecture where the server listens for incoming connections and clients connect to the server to participate in a group chat.

## Objectives

The primary objective of this project is to create a chat application with the following features:

- **TCP Connections**: Support for multiple clients connected to a single server using TCP.
- **Client Identification**: Each client must provide a name upon connection.
- **Message Handling**: Clients can send messages, and these messages are time-stamped and tagged with the client's name.
- **Client Notifications**: The server notifies all clients when a new client joins or leaves the chat.
- **Message History**: New clients receive all previous messages upon joining.
- **Port Specification**: The server can listen on a user-specified port or default to port 8989.

## Features

- **TCP connection handling** for multiple clients (1-to-many relationship).
- **Client name requirement** before participating in the chat.
- **Message broadcasting** with timestamp and client name.
- **Previous messages synchronization** for new clients.
- **Client join/leave notifications** for all connected clients.
- **Maximum of 10 connections** allowed at a time.
- **Error handling** for both server and client sides.

## Setup and Installation

- **Clone the repository**:
    ```bash
    git clone <link> net-cat
    cd net-cat
    ```

- **Ensure you have Go installed**. You can download it from [the official Go website](https://golang.org/dl/).

## Usage

1. **Start the server**:
    ```bash
    go run . [port]
    ```
    or
    ```bash
    make [port]
    ```
   Replace `[port]` with the port number you wish to use (default is 8989 if not specified). For example:
    ```bash
    go run . 2525
    ```
    or
    ```bash
    make 2525
    ```

2. **Connect a client**:
    ```bash
    nc [server_ip] [port]
    ```
   Replace `[server_ip]` with the IP address of the server and `[port]` with the port number (e.g., `nc localhost 2525`).

3. **Follow the prompts** to enter your name and start chatting.

## Example Interaction

**Server:**
```bash
$ ./TCPChat  2525
Listening on the port 2525
```

**Client1 (ZED):**
```
$ nc $IP $port
Welcome to TCP-Chat!

        _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
    `-'       `--'

[ENTER YOUR NAME]:

[ENTER YOUR NAME]: ZED
[2020-01-20 16:03:43][ZED]:hello

```

**Client2 (AMINE):**

```
>:3 ❯ nc localhost 2525
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]: Amine
[2020-01-20 16:03:43][ZED]:hello
[2020-01-20 16:04:15][Amine]:Hi everyone!
[2020-01-20 16:04:32][Amine]:How are you?
[2020-01-20 16:04:35][Amine]:
[2020-01-20 16:04:41][ZED]:great, and you?
[2020-01-20 16:04:41][Amine]:good!
[2020-01-20 16:04:44][Amine]:alright, see ya!
[2020-01-20 16:04:50][Amine]:
[2020-01-20 16:04:57][ZED]:bye-bye!
[2020-01-20 16:04:57][Amine]:^C

```

**Client3 (Mohamed):**
```
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]: ZED
[2024-09-06 16:26:30][ZED]:
[2024-09-06 16:27:01][ZED]:hello!
[2024-09-06 16:27:23][ZED]:how are you     
[2024-09-06 16:27:37][ZED]:
Mohemad has joined our chat...
[2024-09-06 16:27:59][ZED]:
[2024-09-06 16:28:08][Mohemad]:HELLO ZED
[2024-09-06 16:28:08][ZED]:hello Mohemad
[2024-09-06 16:28:24][ZED]:how are you     
[2024-09-06 16:29:15][Mohemad]:i'm fine what about you
[2024-09-06 16:29:32][ZED]: fine, think you
[2024-09-06 16:29:32][ZED]: :)

```

## Bonus:
## Saving Logs and Messages

The Net-Cat project includes functionality to save chat history and logs. Here’s how these are managed:

### Client Logs

- **Log File**: `logs.txt`
- **Description**: This file records all significant events sent by clients, including join and leave notifications. Each entry is timestamped for reference.

### Message History

- **History File**: `messageHistory.txt`
- **Description**: This file stores all previous messages sent in the chat. When a new client joins, they receive the entire chat history from this file, ensuring they are updated with all past communications.

---

Net-Cat Group:

#### asadik |[Amine](https://github.com/AmineS530)|
#### mtawil |[Mohamed](https://github.com/twlmed212)|
#### zzitan |[ZED](https://github.com/0xZED88)|
