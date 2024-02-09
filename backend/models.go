package main

type CampaignCategory struct {
	Category string `json:"category" bind:"required"`
}

type CampaignsResponse struct {
	Campaigns []Campaign `json:"campaigns"`
}

type Campaign struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type CampaignDetail struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Plans []struct {
		PlanTitle string `json:"planTitle"`
		PlanDesc  string `json:"planDesc"`
	} `json:"plans"`
	Goals []struct {
		GoalTitle string `json:"goalTitle"`
		GoalDesc  string `json:"goalDesc"`
	} `json:"goals"`
	Benefit string `json:"benefit"`
	Period  string `json:"period"`
}

type CampaignDetailResponse struct {
	CampaignDetail CampaignDetail `json:"campaignDetail"`
}
