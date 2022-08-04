package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	formmaterCampaign := CampaignFormatter{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageURL:         "",
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
	}

	if len(campaign.CampaignImages) > 0 {
		formmaterCampaign.ImageURL = campaign.CampaignImages[0].FileName
	}

	return formmaterCampaign
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {

	// if len(campaigns) == 0 {
	// 	return []CampaignFormatter{}
	// }

	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		CampaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, CampaignFormatter)
	}

	return campaignsFormatter

}

type CampaignDetailFormmater struct {
	ID               int                       `json:"id"`
	Name             string                    `json:"name"`
	ShortDescription string                    `json:"short_description"`
	Description      string                    `json:"description"`
	ImageURL         string                    `json:"image_url"`
	GoalAmount       int                       `json:"goal_amount"`
	CurrentAmount    int                       `json:"current_amount"`
	UserID           int                       `json:"user_id"`
	Slug             string                    `json:"slug"`
	Perks            []string                  `json:"perks"`
	User             CampaignUserFormmater     `json:"user"`
	Images           []CampaignImagesFormmater `json:"images"`
}

type CampaignUserFormmater struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImagesFormmater struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormmater {

	campaignDetailFormmater := CampaignDetailFormmater{}

	campaignDetailFormmater.ID = campaign.ID
	campaignDetailFormmater.Name = campaign.Name
	campaignDetailFormmater.ShortDescription = campaign.ShortDescription
	campaignDetailFormmater.Description = campaign.Description
	campaignDetailFormmater.ImageURL = ""
	campaignDetailFormmater.GoalAmount = campaign.GoalAmount
	campaignDetailFormmater.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormmater.UserID = campaign.UserID
	campaignDetailFormmater.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormmater.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	campaignDetailFormmater.Perks = perks

	//Json formatter relasi get user
	user := campaign.User

	campaignUserFormmater := CampaignUserFormmater{}
	campaignUserFormmater.Name = user.Name
	campaignUserFormmater.ImageURL = user.AvatarFileName

	//Next formmater
	campaignDetailFormmater.User = campaignUserFormmater

	//Json formmater relasi get image
	images := []CampaignImagesFormmater{}
	for _, image := range campaign.CampaignImages {

		campaignImageFormmater := CampaignImagesFormmater{}
		campaignImageFormmater.ImageURL = image.FileName

		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		campaignImageFormmater.IsPrimary = isPrimary

		images = append(images, campaignImageFormmater)

	}

	//Next formmater
	campaignDetailFormmater.Images = images

	return campaignDetailFormmater

}
