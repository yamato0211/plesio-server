package ws

import (
	"context"

	"github.com/yamato0211/plesio-server/pkg/domain/repository"
)

const broadCastChannel = "broadcast"

type Hub struct {
	Clients      map[*Client]bool
	RegisterCh   chan *Client
	UnRegisterCh chan *Client
	BroadcastCh  chan []byte
	pubsub       repository.RedisRepository
}

func NewHub(pubsub repository.RedisRepository) *Hub {
	return &Hub{
		Clients:      make(map[*Client]bool),
		RegisterCh:   make(chan *Client),
		UnRegisterCh: make(chan *Client),
		BroadcastCh:  make(chan []byte),
		pubsub:       pubsub,
	}
}

func (h *Hub) RunLoop() {
	for {
		select {
		case client := <-h.RegisterCh:
			h.register(client)

		case client := <-h.UnRegisterCh:
			h.unregister(client)

		case msg := <-h.BroadcastCh:
			h.publishMessage(msg)
		}
	}
}

func (h *Hub) SubscribeMessages() {
	ch := h.pubsub.Subscribe(context.TODO(), broadCastChannel)

	for msg := range ch {
		h.broadCastToAllClient([]byte(msg.Payload))
	}
}

func (h *Hub) publishMessage(msg []byte) {
	h.pubsub.Publish(context.TODO(), broadCastChannel, msg)
}

func (h *Hub) register(c *Client) {
	h.Clients[c] = true
}

func (h *Hub) unregister(c *Client) {
	delete(h.Clients, c)
}

func (h *Hub) broadCastToAllClient(msg []byte) {
	for c := range h.Clients {
		c.sendCh <- msg
	}
}
