package service

import (
	"context"
	"github.com/GGuuse-Streams/chatbot-back/internal/db/queries"
	"github.com/GGuuse-Streams/chatbot-back/internal/module/channel/repository"
)

type ChannelService struct {
	r *repository.ChannelRepository
}

func NewChannelService(r *repository.ChannelRepository) *ChannelService {
	return &ChannelService{r: r}
}

func (cs *ChannelService) GetChannels(ctx context.Context) ([]queries.Channel, error) {
	return cs.r.GetChannels(ctx)
}

func (cs *ChannelService) GetChannel(ctx context.Context, id int32) (queries.Channel, error) {
	return cs.r.GetChannel(ctx, id)
}

func (cs *ChannelService) CreateChannel(ctx context.Context, arg queries.CreateChannelParams) (queries.Channel, error) {
	return cs.r.CreateChannel(ctx, arg)
}

func (cs *ChannelService) DeleteChannel(ctx context.Context, id int32) error {
	return cs.r.DeleteChannel(ctx, id)
}

func (cs *ChannelService) UpdateChannel(ctx context.Context, arg queries.UpdateChannelParams) (queries.Channel, error) {
	return cs.r.UpdateChannel(ctx, arg)
}
