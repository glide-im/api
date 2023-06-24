package im

import (
	"github.com/glide-im/api/internal/dao/msgdao"
	"github.com/glide-im/glide/pkg/logger"
	"github.com/glide-im/glide/pkg/messages"
	"github.com/glide-im/glide/pkg/store"
)

func MustInitMessageStorage(kafkaAddr []string) {
	kafkaConsumer, err := store.NewKafkaConsumer(kafkaAddr)
	if err != nil {
		panic(err)
	}
	kafkaConsumer.ConsumeOfflineMessage(onOfflineMessage)
	kafkaConsumer.ConsumeChatMessage(onChatMessage)
	kafkaConsumer.ConsumeChannelMessage(onChannelMessage)
}

func onChatMessage(m *messages.ChatMessage) {
	// TODO 2023年6月24日13:46:10  没有 mid
	_, err := msgdao.AddChatMessage(&msgdao.ChatMessage{
		MID:       m.Mid,
		SessionID: m.To,
		CliSeq:    0,
		From:      m.From,
		To:        m.To,
		Type:      m.Type,
		SendAt:    m.SendAt,
		Content:   m.Content,
	})
	if err != nil {
		logger.E("error on store chat message %v", err)
	}
}

func onOfflineMessage(m *messages.ChatMessage) {
	_, err := msgdao.AddChatMessage(&msgdao.ChatMessage{
		MID:       m.Mid,
		SessionID: m.To,
		CliSeq:    0,
		From:      m.From,
		To:        m.To,
		Type:      m.Type,
		SendAt:    m.SendAt,
		Content:   m.Content,
	})
	if err != nil {
		logger.E("error on store offline message %v", err)
	}
}

func onChannelMessage(m *messages.ChatMessage) {

}
