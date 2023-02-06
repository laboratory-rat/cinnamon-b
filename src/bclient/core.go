package bclient

import (
	"api/main/src/bclient/proto"
)

type Core struct {
	collection map[string]proto.ConnectorServer
}

func NewBClientCore() *Core {
	return &Core{
		collection: make(map[string]proto.ConnectorServer),
	}
}

func (c Core) AddClient(name string, client proto.ConnectorServer) {
	c.collection[name] = client
}

func (c Core) GetClient(name string) (proto.ConnectorServer, error) {
	client, ok := c.collection[name]
	if !ok {
		return nil, ClientNotFound
	}

	return client, nil
}
