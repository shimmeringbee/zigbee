package zigbee

import (
	"context"
)

type NodeOperations interface {
	QueryNodeDescription(ctx context.Context, networkAddress NetworkAddress) (NodeDescription, error)
	QueryNodeEndpoints(ctx context.Context, networkAddress NetworkAddress) ([]byte, error)
	QueryNodeEndpointDescription(ctx context.Context, networkAddress NetworkAddress, endpoint byte) (EndpointDescription, error)

	BindToNode(ctx context.Context, networkAddress NetworkAddress, sourceAddress IEEEAddress, sourceEndpoint byte, destinationAddress IEEEAddress, destinationEndpoint byte, cluster ZCLClusterID) error
	UnbindFromNode(ctx context.Context, networkAddress NetworkAddress, sourceAddress IEEEAddress, sourceEndpoint byte, destinationAddress IEEEAddress, destinationEndpoint byte, cluster ZCLClusterID) error
}
