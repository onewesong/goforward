package forward

import (
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/onewesong/goforward/internal/models"
	"github.com/onewesong/goforward/internal/pkg/log"
)

type Status string

const (
	Idle    Status = "idle"
	Running Status = "running"
	Fail    Status = "fail"
	Quit    Status = "quit"
)

type Forward struct {
	Link models.ForwardLink

	listener net.Listener
	Status   Status
	quit     chan bool
}

func (f *Forward) CreateListener() (err error) {
	f.listener, err = net.Listen(f.Link.ListenAddr.Network(), f.Link.ListenAddr.String())
	return
}

func (f *Forward) Run() {
	for {
		select {
		case <-f.quit:
			log.Info("forward quit")
			f.Status = Quit
			return
		default:
			f.Status = Running
			conn, err := f.listener.Accept()
			if err != nil {
				select {
				case <-f.quit:
					f.quit <- true
				default:
					log.Error("conn accept err: %s", err)
				}
				continue
			}

			log.Info("Accepted connection from %s", conn.RemoteAddr())
			go func() {
				targetConn, err := net.Dial(f.Link.TargetAddr.Network(), f.Link.TargetAddr.String())
				if err != nil {
					conn.Write([]byte(err.Error()))
					return
				}
				defer targetConn.Close()

				go func() {
					_, err := io.Copy(targetConn, conn)
					if err != nil {
						log.Error("forward copy err: %s", err)
					}
				}()

				_, err = io.Copy(conn, targetConn)
				if err != nil {
					log.Error("forward copy err: %s", err)
				}

				log.Info("Forwarded from client %s to %s done", conn.RemoteAddr(), targetConn.RemoteAddr())
			}()
		}
	}
}

func (f *Forward) Start() error {
	f.Status = Idle
	err := f.CreateListener()
	if err != nil {
		f.Status = Fail
		return err
	}
	log.Info("Forward listening on %s to %s", f.listener.Addr(), f.Link.TargetAddr)
	go f.Run()
	return nil
}

func (f *Forward) Stop() error {
	f.quit <- true
	if f.listener != nil {
		if err := f.listener.Close(); err != nil {
			return fmt.Errorf("forward stop err: %s", err)
		}
	}
	return nil
}

func (f *Forward) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		TargetAddr string `json:"target_addr"`
		Status     string `json:"status"`
	}{
		TargetAddr: f.Link.TargetAddr.String(),
		Status:     string(f.Status),
	})
}

func NewForward(forwardLink models.ForwardLink) *Forward {
	return &Forward{
		Link: forwardLink,
		quit: make(chan bool, 1),
	}
}
