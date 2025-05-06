package models

import "github.com/bwmarrin/discordgo"

type ThreadStatus int

const (
	THREADOPEN	ThreadStatus	= iota
	THREADCONNECTED
)

type Thread struct {
	ID		string
	Status	ThreadStatus


	ConnectedTo *Thread
}

func NewThread(id string) *Thread {
	return &Thread{
		ID: id,
		Status: THREADOPEN,
	}
}

func (thread Thread) SendMessage(session *discordgo.Session, message string) {
	session.ChannelMessageSend(thread.ID, message)
}