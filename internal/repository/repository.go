package repository

import (
	"fmt"
	"slices"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/jeenyuhs/Goose/internal/models"
)

type Server struct {
	mu			sync.Mutex
	calls 		map[string]*models.Thread
}

func NewServer() *Server {
	s := &Server{
		calls: make(map[string]*models.Thread),
	}

	return s
}

func (s *Server) GetExistingCall(id string) (*models.Thread, error) {
	s.mu.Lock()
	call, ok := s.calls[id]
	s.mu.Unlock()

	if !ok {
		return nil, fmt.Errorf("requested call doesn't exist")
	}

	return call, nil
}

func (s *Server) GetAvailableCall(ignore ...string) (*models.Thread, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, call := range s.calls {
		if slices.Contains(ignore, call.ID) {
			continue
		}
		
		if call.Status == models.THREADOPEN {
			return call, nil
		}
	}

	return nil, fmt.Errorf("no available call")
}

func (s *Server) AddCall(thread *models.Thread) error {
	s.mu.Lock()
	_, ok := s.calls[thread.ID]
	s.mu.Unlock()

	if ok {
		return fmt.Errorf("call already exists on server")
	}

	s.mu.Lock()
	s.calls[thread.ID] = thread
	s.mu.Unlock()

	return nil
}

func (s *Server) DeleteCall(thread *models.Thread) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.calls[thread.ID]
	if !ok {
		panic("thread id doesn't exist")
	}

	delete(s.calls, thread.ID)
}

func (s *Server) DebugCalls() {
	for k := range s.calls {
		fmt.Println(k)
	}
}

func (s *Server) WaitForCall(session *discordgo.Session, thread *models.Thread) {
	if len(s.calls) < 2 {
		return
	}

	for {
		recipient, err := s.GetAvailableCall(thread.ID)
		if err != nil {
			continue
		}

		s.mu.Lock()
		defer s.mu.Unlock()

		thread.Status = models.THREADCONNECTED
		recipient.Status = models.THREADCONNECTED
		
		thread.ConnectedTo = recipient
		recipient.ConnectedTo = thread

		thread.SendMessage(session, ":telephone:  **Someone picked up!**")
		recipient.SendMessage(session, ":telephone:  **Someone picked up!**")

		return
	}
}