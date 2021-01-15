package tcp_server

import (
	"bufio"
	"log"
	"net"

	"github.com/goforbroke1006/lock-free-research/internal"
)

func New(address string) *tcpServer {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	return &tcpServer{
		listener: listener,
	}
}

var _ internal.TCPListener = &tcpServer{}

type tcpCallback func(message string, reply *string)

type tcpServer struct {
	listener  net.Listener
	callbacks []tcpCallback
}

func (srv *tcpServer) Run() {
	for {
		conn, err := srv.listener.Accept()
		if err != nil {
			log.Printf("error accepting connection %v", err)
			continue
		}

		go func() {
			defer func() {
				log.Printf("closing connection from %v", conn.RemoteAddr())
				conn.Close()
			}()
			r := bufio.NewReader(conn)
			w := bufio.NewWriter(conn)
			scanner := bufio.NewScanner(r)

			for scanner.Scan() {
				payload := scanner.Text()
				for _, cb := range srv.callbacks {
					go func(cb tcpCallback) {
						var replay string
						cb(payload, &replay)
						if len(replay) > 0 {
							_, _ = w.WriteString(replay + "\n")
							_ = w.Flush()
						}
					}(cb)
				}
			}
		}()

	}
}

func (srv *tcpServer) OnMessage(cb func(message string, reply *string)) {
	srv.callbacks = append(srv.callbacks, cb)
}

func (srv *tcpServer) Stop() error {
	return srv.listener.Close()
}
