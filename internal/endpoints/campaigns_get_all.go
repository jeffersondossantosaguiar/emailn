package endpoints

import (
	"net/http"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	campaigns, err := h.CampaignService.Get()
	return campaigns, http.StatusOK, err
}
