package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/jeenyuhs/Goose/internal/models"
	"github.com/jeenyuhs/Goose/internal/repository"
)

type MessageHandler struct {
    server *repository.Server
}

func NewMessageHandler(server *repository.Server) *MessageHandler {
    return &MessageHandler{
        server: server,
    }
}

func (h *MessageHandler) HandleMessages(session *discordgo.Session, message *discordgo.MessageCreate) {
	channel, err := session.State.Channel(message.ChannelID)
	if err != nil {
		panic(err)
	}

	if !channel.IsThread() && message.Content == "!call" {
		thread, err := session.MessageThreadStart(message.ChannelID, message.ID, "Phone Call", 60)

		if err != nil {
			fmt.Println("Error starting thread:", err)
			return
		}

		var t *models.Thread = models.NewThread(thread.ID)

		err = h.server.AddCall(t)
		if err != nil {
			panic(err)
		}

		session.ChannelMessageSend(thread.ID, ":telephone:  **Waiting for someone to pick up...**")

		go h.server.WaitForCall(session, t)
	}

	if channel.IsThread() && message.Author.ID != session.State.User.ID {
		call, err := h.server.GetExistingCall(message.ChannelID)
		// the call doesn't exist, which means 
		// the thread is not a call.
		if err != nil {
			return
		}

		// the thread is not connected to anybody
		// therefore it should not do anything 
		if call.Status == models.THREADOPEN {
			return
		}

		if message.Content == "!end" {
			call.SendMessage(session, ":headstone:  **You've ended the call! This phone call is now archived.**")
			call.ConnectedTo.SendMessage(session, ":headstone: **The person on the other line ended the call! This phone call is now archived.**")
			
			archived := true
			locked := true

			session.ChannelEdit(call.ID, &discordgo.ChannelEdit{
				Archived: 	&archived,
				Locked:		&locked,
			})

			session.ChannelEdit(call.ConnectedTo.ID, &discordgo.ChannelEdit{
				Archived: 	&archived,
				Locked:		&locked,
			})
			
			// delete calls from repository
			h.server.DeleteCall(call)
			h.server.DeleteCall(call.ConnectedTo)
			
			return
		}

		message := fmt.Sprintf(":speaking_head:  %s", message.Content)
		call.ConnectedTo.SendMessage(session, message)
	}
}