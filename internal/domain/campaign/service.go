package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	Get() ([]contract.CampaignResponse, error)
	GetBy(id string) (*contract.CampaignResponse, error)
}

type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) Create(newCampaign contract.NewCampaign) (string, error) {

	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)

	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return campaign.ID, nil
}

func (s *ServiceImp) GetBy(id string) (*contract.CampaignResponse, error) {
	campaign, err := s.Repository.GetBy(id)

	if err != nil {
		return nil, internalerrors.ErrInternal
	}

	return &contract.CampaignResponse{
		ID:      campaign.ID,
		Name:    campaign.Name,
		Content: campaign.Content,
		Status:  campaign.Status,
	}, nil
}

func (s *ServiceImp) Get() ([]contract.CampaignResponse, error) {
	campaigns, err := s.Repository.Get()

	if err != nil {
		return nil, internalerrors.ErrInternal
	}

	var responses []contract.CampaignResponse

	for _, campaign := range campaigns {
		response := contract.CampaignResponse{
			ID:      campaign.ID,
			Name:    campaign.Name,
			Content: campaign.Content,
			Status:  campaign.Status,
		}
		responses = append(responses, response)
	}

	return responses, nil
}
