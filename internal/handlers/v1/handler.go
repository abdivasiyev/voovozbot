package v1

import (
	"strconv"

	"github.com/MrWebUzb/voovozbot/internal/storage"
	"go.uber.org/zap"
	tb "gopkg.in/tucnak/telebot.v2"
)

type HandlerFunc func(*tb.Message)

type HandlerV1 struct {
	log  *zap.Logger
	b    *tb.Bot
	strg storage.StorageI
}

func NewHandlerV1(b *tb.Bot, log *zap.Logger, strg storage.StorageI) *HandlerV1 {
	return &HandlerV1{
		log:  log,
		b:    b,
		strg: strg,
	}
}

func parseOffset(offset string) int {
	v, err := strconv.Atoi(offset)
	if err != nil {
		return 0
	}

	return v
}
