package performance

import (
	"context"
	"time"

	"github.com/lavanet/lava/protocol/chainlib/chainproxy"
	pairingtypes "github.com/lavanet/lava/x/pairing/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Cache struct {
	client  pairingtypes.RelayerCacheClient
	address string
}

func ConnectGRPCConnectionToRelayerCacheService(ctx context.Context, addr string) (*pairingtypes.RelayerCacheClient, error) {
	connectCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(connectCtx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(chainproxy.MaxCallRecvMsgSize)))
	if err != nil {
		return nil, err
	}
	/*defer conn.Close()*/

	c := pairingtypes.NewRelayerCacheClient(conn)
	return &c, nil
}

func InitCache(ctx context.Context, addr string) (*Cache, error) {
	relayerCacheClient, err := ConnectGRPCConnectionToRelayerCacheService(ctx, addr)
	if err != nil {
		return &Cache{client: nil, address: addr}, err
	}
	cache := Cache{client: *relayerCacheClient, address: addr}
	return &cache, nil
}

func (cache *Cache) GetEntry(ctx context.Context, relayCacheGet *pairingtypes.RelayCacheGet) (reply *pairingtypes.CacheRelayReply, err error) {
	if cache == nil {
		// TODO: try to connect again once in a while
		return nil, NotInitializedError
	}
	if cache.client == nil {
		return nil, NotConnectedError.Wrapf("No client connected to address: %s", cache.address)
	}
	// TODO: handle disconnections and error types here
	return cache.client.GetRelay(ctx, relayCacheGet)
}

func (cache *Cache) CacheActive() bool {
	return cache != nil
}

func (cache *Cache) SetEntry(ctx context.Context, cacheSet *pairingtypes.RelayCacheSet) error {
	if cache == nil {
		// TODO: try to connect again once in a while
		return NotInitializedError
	}
	if cache.client == nil {
		return NotConnectedError.Wrapf("No client connected to address: %s", cache.address)
	}
	// TODO: handle disconnections and SetRelay error types here
	_, err := cache.client.SetRelay(ctx, cacheSet)
	return err
}
