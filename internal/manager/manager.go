package manager

import (
	"fmt"

	"github.com/onewesong/goforward/internal/models"
	"github.com/onewesong/goforward/internal/pkg/forward"
)

type Manager struct {
	ForwardMap map[string]*forward.Forward
}

func (m Manager) AddForward(ForwardLinks models.ForwardLinks, override bool) error {
	for _, i := range ForwardLinks {
		key := i.ListenAddr.String()
		forward := forward.NewForward(i)
		if f, ok := m.ForwardMap[key]; ok {
			if override {
				err := f.Stop()
				if err != nil {
					return fmt.Errorf("forward %s stop err: %s", key, err)
				}
			} else {
				return fmt.Errorf("forward %s already exists, if want to override, set param override to true", key)
			}
		}
		err := forward.Start()
		if err != nil {
			return fmt.Errorf("forward %s start err: %s", key, err)
		}
		m.ForwardMap[key] = forward
	}
	return nil
}

func (m Manager) DelForward(listenAddr string) error {
	if _, ok := m.ForwardMap[listenAddr]; !ok {
		return fmt.Errorf("forward %s not found", listenAddr)
	}
	err := m.ForwardMap[listenAddr].Stop()
	if err != nil {
		return fmt.Errorf("forward %s stop err: %s", listenAddr, err)
	}
	delete(m.ForwardMap, listenAddr)
	return nil
}

var manager *Manager

func init() {
	manager = &Manager{
		ForwardMap: make(map[string]*forward.Forward),
	}
}

func GetInstance() *Manager {
	return manager
}
