package logic

import (
	"chat/protocol"
	"go-lib/log"

	"google.golang.org/grpc"
)

func (s *ChatLogic) AddNode(addr string) {
	s.Lock()
	defer s.Unlock()
	var conn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return
	}
	if _, ok := s.roomClients[addr]; ok {
		delete(s.roomClients, addr)
	}
	var client = protocol.NewRoomServiceClient(conn)
	s.roomClients[addr] = client
}
func (s *ChatLogic) DelNode(addr string) {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.roomClients[addr]; ok {
		delete(s.roomClients, addr)
	}
}

func (s *ChatLogic) UpdateNode(addr string) {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.roomClients[addr]; ok {
		return
	}
	var conn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return
	}
	var client = protocol.NewRoomServiceClient(conn)
	s.roomClients[addr] = client
}

func (s *ChatLogic) Watch() {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("watcher paniced, recover:", err)
		}
	}()

	services, err := s.Registry.GetService("live-chat.voip")
	if err == nil {
		for _, srv := range services {
			for _, node := range srv.Nodes {
				s.AddNode(node.Address)
				log.Infof("add node %s :%s", srv.Name, node.Address)
			}
		}
	} else {
		log.Error(err)
	}

	for {
		select {
		case <-s.stop:
			return
		default:
			res, err := s.watcher.Next()
			if err != nil {
				log.Errorf("watch error:%v", err)
				continue
			}

			var name = res.Service.Name

			if name == "live-chat.voip" {
				switch res.Action {
				case "create":
					for _, node := range res.Service.Nodes {
						s.AddNode(node.Address)
						log.Infof("new node %s :%s", name, node.Address)
					}

				case "delete":
					for _, node := range res.Service.Nodes {
						s.DelNode(node.Address)
						log.Infof("del node %s :%s", name, node.Address)
					}
				case "update":
					for _, node := range res.Service.Nodes {
						s.UpdateNode(node.Address)
						log.Infof("update node %s :%s", name, node.Address)
					}
				}
			} else {
				for _, node := range res.Service.Nodes {
					log.Infof("%s node %s :%s", res.Action, name, node.Address)
				}
			}
		}
	}
}
