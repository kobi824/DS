package handlers

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type Serve struct {
	Node     *maelstrom.Node
	Topology map[string][]string
	Messages []int
}

func (s *Serve) Run() error {
	return s.Node.Run()
}

type Response struct {
	Type string
}

func (s *Serve) ReadHandler(req maelstrom.Message) error {
	//anonymous struct -- learned something new
	x := struct {
		Type     string `json:"type"`
		Messages []int  `json:"messages"`
	}{Type: "read_ok",
		Messages: s.Messages,
	}

	return s.Node.Reply(req, x)
}

func (s *Serve) BroadcastHandler(req maelstrom.Message) error {
	var body struct {
		Message int `json:"message"`
	}
	if err := json.Unmarshal(req.Body, &body); err != nil {
		return err
	}

	s.Messages = append(s.Messages, body.Message)

	x := Response{
		Type: "broadcast_ok",
	}

	return s.Node.Reply(req, x)
}

func (s *Serve) TopologyHandler(req maelstrom.Message) error {
	// "type": "topology",
	// "topology": {
	//   "n1": ["n2", "n3"],
	//   "n2": ["n1"],
	//   "n3": ["n1"]
	var body struct {
		Topology map[string][]string `json:"topology"`
	}
	if err := json.Unmarshal(req.Body, &body); err != nil {
		return err
	}

	s.Topology = body.Topology
	x := Response{
		Type: "topology_ok",
	}

	return s.Node.Reply(req, x)
}
