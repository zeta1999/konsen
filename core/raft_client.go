package core

import (
	"context"

	konsen "github.com/lizhaoliu/konsen/v2/proto_gen"
)

// RaftClient is a generic interface for Raft client.
type RaftClient interface {
	// AppendEntries sends AppendEntries request to the remote server.
	AppendEntries(ctx context.Context, in *konsen.AppendEntriesReq) (*konsen.AppendEntriesResp, error)

	// RequestVote sends RequestVote request to the remote server.
	RequestVote(ctx context.Context, in *konsen.RequestVoteReq) (*konsen.RequestVoteResp, error)
}
