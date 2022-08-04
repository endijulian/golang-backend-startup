package campaign

type GetCampaignDetailInput struct {
	ID int `uri:"id" bindfing:"required"`
}
