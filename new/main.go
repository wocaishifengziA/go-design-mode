package new

import "fmt"

type Service interface {
	Start()
	Stop()
}

type server struct{}

func (s *server) Start() {
	fmt.Println("start")
}

func (s *server) Stop() {
	fmt.Println("stop")
}

func NewService() Service {
	return &server{}
}
