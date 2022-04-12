package main

type commandID int

const (
	CMD_NICK  = 1
	CMD_JOIN  = 2
	CMD_ROOMS = 3
	CMD_MSG   = 4
	CMD_QUIT  = 5
)

type command struct {
	id     commandID
	client *client
	args   []string
}
