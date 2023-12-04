package main

import (
	"log"
	"maelstrom-broadcast/handlers"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	server := New(maelstrom.NewNode())

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

// type Server struct {
// 	Node     *maelstrom.Node
// 	Topology map[string][]string
// 	Messages []int
// }

func New(node *maelstrom.Node) *handlers.Serve {
	s := &handlers.Serve{
		Node: node,
	}

	node.Handle("broadcast", s.BroadcastHandler)
	node.Handle("read", s.ReadHandler)
	node.Handle("topology", s.TopologyHandler)

	return s
}
