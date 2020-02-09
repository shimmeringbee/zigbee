package zigbee

import (
	"context"
)

type NodeOperations interface {
	QueryNodeEndpoints(ctx context.Context, networkAddress NetworkAddress) ([]byte, error)
	QueryNodeEndpointDescription(ctx context.Context, networkAddress NetworkAddress, endpoint byte) (EndpointDescription, error)
}
