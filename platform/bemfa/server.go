package bemfa

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Server struct {
	conn net.Conn
	uid  string
}

func NewServer(uid string) (*Server, error) {
	conn, err := net.Dial("tcp", "bemfa.com:8344")
	if err != nil {
		return nil, err
	}
	return &Server{conn, uid}, nil
}

func (s *Server) send(command string) error {
	_, err := s.conn.Write([]byte(fmt.Sprintf("%s\r\n", command)))
	return err
}

func (s *Server) HeartBeat() error {
	return s.send("ping")
}

func (s *Server) Subscription(topics ...string) error {
	// cmd=1&uid=xxxxxxxxxxxxxxxxxxxxxxx&topic=xxx1,xxx2,xxx3,xxx4
	err := s.send(fmt.Sprintf("cmd=1&uid=%s&topic=%s", s.uid, strings.Join(topics, ",")))

	buffer := make([]byte, 1024)
	n, err := s.conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return err
	}

	// 打印服务器的响应
	fmt.Printf("Server response: %s\n", string(buffer[:n]))
	return nil
}

func (s *Server) Publish(topic string, msg string) error {
	// cmd=2&uid=xxxxxxxxxxxxxxxxxxxxxxx&topic=xxx&msg=xxx
	return s.send(fmt.Sprintf("cmd=2&uid=%s&topic=%s&msg=%s", s.uid, topic, msg))
}

func (s *Server) Close() error {
	return s.conn.Close()
}

func (s *Server) Read() {

	buffer := make([]byte, 1024)
	for {
		// 读取服务器的响应
		n, err := s.conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}

		// 打印服务器的响应
		fmt.Printf("Server response: %s\n", string(buffer[:n]))

		// 可选：添加延迟以避免CPU占用过高
		time.Sleep(1 * time.Second)
	}
}
