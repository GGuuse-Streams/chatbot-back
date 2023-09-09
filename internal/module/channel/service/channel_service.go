package service

import (
	"github.com/GGuuse-Streams/chatbot-back/internal/db/queries"
	"github.com/GGuuse-Streams/chatbot-back/internal/module/channel/repository"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type ChannelService struct {
	r *repository.ChannelRepository
}

func NewChannelService(r *repository.ChannelRepository) *ChannelService {
	return &ChannelService{r: r}
}

func (cs *ChannelService) GetChannels(c *fiber.Ctx) error {
	channels, err := cs.r.GetChannels(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(channels)
}

func (cs *ChannelService) GetChannel(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	channel, err := cs.r.GetChannel(c.Context(), int32(id))
	if err != nil {
		return err
	}

	return c.JSON(channel)
}

func (cs *ChannelService) CreateChannel(c *fiber.Ctx) error {
	var params queries.CreateChannelParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	if params.TwitchName == "" || params.TwitchID == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "invalid twitch name or id")
	}

	channel, err := cs.r.CreateChannel(c.Context(), params)
	if err != nil {
		return err
	}

	return c.JSON(channel)
}

func (cs *ChannelService) DeleteChannel(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	err = cs.r.DeleteChannel(c.Context(), int32(id))
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (cs *ChannelService) UpdateChannel(c *fiber.Ctx) error {
	var params queries.UpdateChannelParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	channel, err := cs.r.UpdateChannel(c.Context(), params)
	if err != nil {
		return err
	}

	return c.JSON(channel)
}
