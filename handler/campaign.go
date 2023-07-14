package handler

import (
	"net/http"
	"startupfunding/campaign"
	"startupfunding/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHander struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHander {
	return &campaignHander{service}
}

func (h *campaignHander) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaign, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helpers.APIResponse("Error to get campaigns.", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("Success to get campaigns.", http.StatusOK, "success", campaign)
	c.JSON(http.StatusOK, response)
}
