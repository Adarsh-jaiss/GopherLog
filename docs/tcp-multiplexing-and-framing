## TCP Multiplexing 

- It means that the TCP stack can handle multiple different connections simultaneously, and it identifies each connection using the connection's 4-tuple:

```
(source IP, source port, destination IP, destination port)
```

## Framing 

- Framing tells the receiver where messages begin/end
```
Hello
How are you?
Goodbye
```

You need some way for the receiver to know:

```
┌─────────────┐
│ Hello       │
├─────────────┤
│ How are you?│
├─────────────┤
│ Goodbye     │
└─────────────┘
```

There are several ways to do this.
1. Method 1: Delimiter

You could use \n: 
```
Hello\n
How are you?\n
Goodbye\n
```

2. Method 2: Length prefix

[4 bytes: length][message]

eg : 

```
[0005][HELLO]
[0011][How are you?]
[0007][Goodbye]
```

| Concept          | Question it answers                            |
| ---------------- | ---------------------------------------------- |
| **Multiplexing** | "Which connection does these bytes belong to?" |
| **Framing**      | "Which message does these bytes belong to?"    |


### Architecture of TCP Multiplexing and Framing

```

                    TCP Server
                        │
                  net.Listener
                        │
              Accept connections
                        │
              ┌─────────┴─────────┐
              │                   │
        goroutine #1        goroutine #2
              │                   │
        TCP connection       TCP connection
              │                   │
        Frame Decoder        Frame Decoder
              │                   │
        Request Decoder      Request Decoder
              │                   │
              └─────────┬─────────┘
                        │
                  Request Router
                        │
       ┌────────┬───────┼────────┬────────┐
       │        │       │        │        │
    PRODUCE   FETCH   OFFSET   COMMIT   JOIN_GROUP

```
