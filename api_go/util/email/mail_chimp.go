package email

import (
	"api/util/log"
	"fmt"
	"github.com/zeekay/gochimp3"
)

const (
	UnRegisterListID = "1cc10602dc"
	RegisterListID = "a85ad350a8"
	NewsletterListID = "3307298445"
	AdminListID = "15d9ba4089"

	ApiKey = "a1477662ff7be4731ae498613db055f8-us12"
)

type MailChimp struct {
	ChimpAPI	*gochimp3.API
	CampaignId	string
}

func MailChimpNew() (mailChimpAPI *MailChimp) {
	mailChimpAPI = &MailChimp{}
	mailChimpAPI.ChimpAPI = gochimp3.New(ApiKey)
	return mailChimpAPI
}

func (api *MailChimp)AddUserToList(email string, listId string) (id string){
	if listId == "" {
		listId = UnRegisterListID
	}

	member := &gochimp3.MemberRequest{}

	member.EmailAddress = email
	member.Status = "subscribed"

	ret, err := api.ChimpAPI.NewListResponse(listId).CreateMember(member)

	if err != nil{
		return "0"
	}

	return ret.ID
}

func (api *MailChimp)DeleteUserFormList(id string, listId string){
	if listId == "" {
		listId = UnRegisterListID
	}
	_, err := api.ChimpAPI.NewListResponse(listId).DeleteMember(id)
	log.Error(err)
}

func (api *MailChimp)AddCampaign(title string, email string, listId string) *MailChimp{

	if listId == "" {
		listId = UnRegisterListID
	}

	campaignNew := &gochimp3.CampaignCreationRequest{}

	campaignNew.Type = gochimp3.CAMPAIGN_TYPE_REGULAR

	campaignNew.Settings = gochimp3.CampaignCreationSettings{
		SubjectLine: title,
		FromName: "7SevenCoin團隊",
		ReplyTo: "7Sevencoin@7officer.com",
	}
	var paramMap []interface{}

	condition := map[string]interface{}{
		"condition_type": "EmailAddress",
		"op": "is",
		"field": "EMAIL",
		"value": email,
	}

	paramMap = append(paramMap, condition)

	campaignNew.Recipients = gochimp3.CampaignCreationRecipients{
		ListId: listId,
		SegmentOptions: gochimp3.CampaignCreationSegmentOptions{
			Match: "all",
			Conditions: paramMap,
		},
	}

	campaignNew.Tracking = gochimp3.CampaignTracking{
		Opens: true,
		HtmlClicks: true,
	}

	res, _ := api.ChimpAPI.CreateCampaign(campaignNew)

	api.CampaignId = res.ID

	return api
}

func (api *MailChimp)AddCampaignForAllPerson(title string, listId string) *MailChimp{

	if listId == "" {
		listId = UnRegisterListID
	}

	campaignNew := &gochimp3.CampaignCreationRequest{}

	campaignNew.Type = gochimp3.CAMPAIGN_TYPE_REGULAR

	campaignNew.Settings = gochimp3.CampaignCreationSettings{
		SubjectLine: title,
		FromName: "7SevenCoin團隊",
		ReplyTo: "7Sevencoin@7officer.com",
	}

	campaignNew.Recipients = gochimp3.CampaignCreationRecipients{
		ListId: listId,
		SegmentOptions: gochimp3.CampaignCreationSegmentOptions{
			SavedSegmentId: 262429,
			Match: "all",
			Conditions: []string{},
		},
	}

	campaignNew.Tracking = gochimp3.CampaignTracking{
		Opens: true,
		HtmlClicks: true,
	}

	res, _ := api.ChimpAPI.CreateCampaign(campaignNew)

	api.CampaignId = res.ID

	return api
}

func (api *MailChimp)EditCampaignContent(content string) *MailChimp{
	response := new(gochimp3.CampaignResponse)

	params := map[string]interface{}{
		"html": content,
	}

	log.Error(api.ChimpAPI.Request("PUT", "/campaigns/" + api.CampaignId +"/content", nil, params, response))

	return api
}

func (api *MailChimp)SendCampaign() bool{
	sendId := &gochimp3.SendCampaignRequest{
		CampaignId : api.CampaignId,
	}

	res, err := api.ChimpAPI.SendCampaign(api.CampaignId, sendId)

	if err != nil{
		fmt.Println(err)
		return false
	}

	return res
}