package palmetto

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cockroachdb/cockroach/client"
	"github.com/cockroachdb/cockroach/proto"
	"github.com/cockroachdb/cockroach/rpc"
)

type protoClient struct {
	kv *client.KV
}

func ProtoClient(srvAddr string) Client {
	snd := client.NewHTTPSender(srvAddr, &http.Transport{
		TLSClientConfig:       rpc.LoadInsecureTLSConfig().Config(),
		ResponseHeaderTimeout: 500 * time.Millisecond,
		TLSHandshakeTimeout:   500 * time.Millisecond,
	})
	return &protoClient{kv: client.NewKV(snd, nil)}
}

func makeKey(projID, name string) proto.Key {
	return proto.Key(fmt.Sprintf("%s|%s", projID, name))
}

func (c *protoClient) Get(key string) (*PlainResult, error) {
	protoKey := proto.Key(key)
	resp := &proto.GetResponse{}
	if err := c.kv.Call(proto.Get, proto.GetArgs(protoKey), resp); err != nil {
		return nil, err
	}
	return nil, ErrNotFound
}

func (c *protoClient) Put(key string, val []byte) (*PlainResult, error) {
	return nil, ErrNotFound
}
