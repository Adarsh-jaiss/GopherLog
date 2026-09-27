package client

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

const MaxFrameSize = 10 * 1024 * 1024

type Server struct {
	PORT string
	HOST string
}

func NewServer(s *Server) *Server {
	return &Server{
		PORT: s.PORT,
		HOST: s.HOST,
	}
}

func (s *Server) InitTCPServer() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", s.PORT))
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go HandleConnection(conn)
	}
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		// Read frames from conn
		payload, err := readFrame(conn)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("client disconnected")
				return
			}
			log.Println("read frame err:", err)
			return
		}

		log.Println("Received payload:", string(payload))

		err = writeFrame(conn, []byte("PONG"))
		if err != nil {
			return
		}

		// Decode requests
		// Route requests
		// Write responses
	}

}

func readFrame(r io.Reader) ([]byte, error) {
	var length uint32
	if err := binary.Read(r, binary.BigEndian, &length); err != nil {
		return nil, err
	}

	if length > MaxFrameSize {
		return nil, errors.New("frame too large")
	}

	payload := make([]byte, length)

	if _, err := io.ReadFull(r, payload); err != nil {
		return nil, err
	}
	return payload, nil
}

func writeFrame(conn net.Conn, payload []byte) error {
	frame := make([]byte, 4+len(payload))

	binary.BigEndian.PutUint32(frame[:4], uint32(len(payload)))

	copy(frame[4:], payload)

	_, err := conn.Write(frame)
	return err
}
