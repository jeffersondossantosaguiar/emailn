package campaign

type Repository interface {
	Create(campaign *Campaign) error
	Update(campaign *Campaign) error
	Get() ([]Campaign, error)
	GetBy(id string) (*Campaign, error)
	DeleteBy(campaign *Campaign) error
}
