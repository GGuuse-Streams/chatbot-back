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

func (c *ChannelRepository) CreateChannel(ctx context.Context, arg queries.CreateChannelParams) (queries.Channel, error) {
	channel, err := c.q.CreateChannel(ctx, arg)
	if err != nil {
		return queries.Channel{}, err
	}

	return channel, nil
}

func (c *ChannelRepository) DeleteChannel(ctx context.Context, id int32) error {
	err := c.q.DeleteChannel(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (c *ChannelRepository) GetChannel(ctx context.Context, id int32) (queries.Channel, error) {
	channel, err := c.q.GetChannel(ctx, id)
	if err != nil {
		return queries.Channel{}, err
	}

	return channel, nil
}

func (c *ChannelRepository) GetChannels(ctx context.Context) ([]queries.Channel, error) {
	channels, err := c.q.GetChannels(ctx)
	if err != nil {
		return nil, err
	}

	return channels, nil
}

func (c *ChannelRepository) UpdateChannel(ctx context.Context, arg queries.UpdateChannelParams) (queries.Channel, error) {
	channel, err := c.q.UpdateChannel(ctx, arg)
	if err != nil {
		return queries.Channel{}, err
	}

	return channel, nil
}
