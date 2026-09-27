package client

const (
	APIProduce   uint16 = 1
	APIFetch     uint16 = 2
	APIOffset    uint16 = 3
	APICommit    uint16 = 4
	APIJoinGroup uint16 = 5
)

type Request struct {
	ID      uint64
	APIKey  uint16
	Version uint16
	Body    []byte
}

// func route(req Request) Response {
// 	switch req.APIKey {
// 	case APIProduce:
// 		return handleProduce(req)

// 	case APIFetch:
// 		return handleFetch(req)

// 	case APIOffset:
// 		return handleOffset(req)

// 	case APICommit:
// 		return handleCommit(req)

// 	case APIJoinGroup:
// 		return handleJoinGroup(req)

// 	default:
// 		return errorResponse("unknown API")
// 	}
// }
