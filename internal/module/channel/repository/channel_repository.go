package repository

import (
	"context"
	"github.com/GGuuse-Streams/chatbot-back/internal/db/queries"
)

type ChannelRepository struct {
	q *queries.Queries
}

func NewChannelRepository(q *queries.Queries) *ChannelRepository {
	return &ChannelRepository{q: q}
}

func (cr *ChannelRepository) CreateChannel(ctx context.Context, arg queries.CreateChannelParams) (queries.Channel, error) {
	return cr.q.CreateChannel(ctx, arg)
}

func (cr *ChannelRepository) DeleteChannel(ctx context.Context, id int32) error {
	return cr.q.DeleteChannel(ctx, id)
}

func (cr *ChannelRepository) GetChannel(ctx context.Context, id int32) (queries.Channel, error) {
	return cr.q.GetChannel(ctx, id)
}

func (cr *ChannelRepository) GetChannels(ctx context.Context) ([]queries.Channel, error) {
	return cr.q.GetChannels(ctx)
}

func (cr *ChannelRepository) UpdateChannel(ctx context.Context, arg queries.UpdateChannelParams) (queries.Channel, error) {
	return cr.q.UpdateChannel(ctx, arg)
}
