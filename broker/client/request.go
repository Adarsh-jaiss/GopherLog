package client

import (
	"encoding/binary"
	"errors"
)

const (
	APIProduce uint16 = iota
	APIFetch
	APIOffset
	APICommit
	APIJoinGroup
)

type Request struct {
	ID      uint32
	APIKey  uint16
	Version uint16
	Body    []byte
}

type Response struct{}

func decodeRequest(data []byte) (*Request, error) {
	if len(data) < 8 {
		return nil, errors.New("invalid request")
	}

	req := &Request{
		ID:      binary.BigEndian.Uint32(data[0:4]),
		APIKey:  binary.BigEndian.Uint16(data[4:6]),
		Version: binary.BigEndian.Uint16(data[6:8]),
		Body:    data[8:],
	}

	return req, nil
}

func route(req Request) *Response {
	switch req.APIKey {
	case APIProduce:
		return handleProduce(req)

	case APIFetch:
		return handleFetch(req)

	case APIOffset:
		return handleOffset(req)

	case APICommit:
		return handleCommit(req)

	case APIJoinGroup:
		return handleJoinGroup(req)

	default:
		return errorResponse("unknown API")
	}
}
