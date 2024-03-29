package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.listRoom(cmd.client)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client)
		}
	}
}

func (s *server) newClient(conn net.Conn) *client {
	fmt.Printf("New client has joined: %s", conn.RemoteAddr().String())
	return &client{
		conn:     conn,
		nick:     "anonymous",
		commands: s.commands,
	}
}

func (s *server) nick(c *client, args []string) {
	if len(args) < 2 {
		c.msg("nick is required. usage /nick NAME")
		return
	}
	c.nick = args[1]
	c.msg(fmt.Sprintf("Welcome %s", c.nick))
}

func (s *server) join(c *client, args []string) {
	if len(args) < 2 {
		c.msg("room name is required. usage /nick NAME")
		return
	}
	roomName := args[1]

	r, ok := s.rooms[roomName]

	if !ok {
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		s.rooms[roomName] = r
	}
	r.members[c.conn.RemoteAddr()] = c

	s.quitCurrentRoom(c)
	c.room = r
	r.broadcast(c, fmt.Sprintf("%s joined the room", c.nick))
	c.msg(fmt.Sprintf("Welcome to %s", roomName))
}

func (s *server) listRoom(c *client) {
	var rooms []string
	for name := range s.rooms {
		rooms = append(rooms, name)
	}
	c.msg(fmt.Sprintf("available rooms: %s", strings.Join(rooms, ", ")))
}

func (s *server) msg(c *client, args []string) {
	if len(args) < 2 {
		c.msg("msg is required. usage /msg MESSAGE")
		return
	}
	msg := strings.Join(args[1:], " ")
	c.room.broadcast(c, c.nick+": "+msg)
}
func (s *server) quit(c *client) {
	log.Printf("Client has left the chat: %s", c.conn.RemoteAddr().String())
	s.quitCurrentRoom(c)
	c.msg("Bye!!")
	c.conn.Close()
}

func (s *server) quitCurrentRoom(c *client) {
	if c.room != nil {
		oldRoom := s.rooms[c.room.name]
		delete(s.rooms[c.room.name].members, c.conn.RemoteAddr())
		oldRoom.broadcast(c, fmt.Sprintf("%s has left the room", c.nick))
	}
}
