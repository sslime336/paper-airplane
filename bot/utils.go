package bot

import (
	"encoding/json"

	"github.com/sslime336/paper-airplane/ws"
	"go.uber.org/zap"
)

func ExtractPrivateChatMessage(data []byte) *ws.PrivateChatMessage {
	var msg ws.PrivateChatMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		log.Error("failed to unmarshal PrivateChatMessage", zap.Error(err))
		return nil
	}
	return &msg
}

func ExtractGroupMessage(data []byte) *ws.GroupAtMessage {
	var msg ws.GroupAtMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		log.Error("failed to unmarshal GroupAtMessage", zap.Error(err))
		return nil
	}
	return &msg
}
