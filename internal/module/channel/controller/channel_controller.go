package controller

import (
	"github.com/GGuuse-Streams/chatbot-back/internal/db/queries"
	"github.com/GGuuse-Streams/chatbot-back/internal/module/channel/service"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type ChannelController struct {
	s *service.ChannelService
}

func NewChannelController(s *service.ChannelService) *ChannelController {
	return &ChannelController{
		s: s,
	}
}

func (cc *ChannelController) GetChannels(c *fiber.Ctx) error {
	channels, err := cc.s.GetChannels(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(channels)
}

func (cc *ChannelController) GetChannel(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	channel, err := cc.s.GetChannel(c.Context(), int32(id))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(channel)
}

func (cc *ChannelController) CreateChannel(c *fiber.Ctx) error {
	var params queries.CreateChannelParams
	if err := c.BodyParser(&params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if params.TwitchName == "" || params.TwitchID == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "invalid twitch name or id")
	}

	channel, err := cc.s.CreateChannel(c.Context(), params)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(channel)
}

func (cc *ChannelController) DeleteChannel(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	err = cc.s.DeleteChannel(c.Context(), int32(id))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (cc *ChannelController) UpdateChannel(c *fiber.Ctx) error {
	var params queries.UpdateChannelParams
	if err := c.BodyParser(&params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	channel, err := cc.s.UpdateChannel(c.Context(), params)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(channel)
}
