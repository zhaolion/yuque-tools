package api

import (
	"encoding/json"
	"time"

	"github.com/zhaolion/yuque-tools/api/pretty"
)

type GroupCreateRequest struct {
	Name        string `json:"name"`
	Login       string `json:"login"`
	Description string `json:"description"`
}

type GroupCreateResponse struct {
	Data GroupDetail `json:"data"`
}

type GroupDetail struct {
	ID                int       `json:"id"`
	SpaceID           int       `json:"space_id"`
	OrganizationID    int       `json:"organization_id"`
	Login             string    `json:"login"`
	Name              string    `json:"name"`
	AvatarURL         string    `json:"avatar_url"`
	OwnerID           int       `json:"owner_id"`
	BooksCount        int       `json:"books_count"`
	PublicBooksCount  int       `json:"public_books_count"`
	TopicsCount       int       `json:"topics_count"`
	PublicTopicsCount int       `json:"public_topics_count"`
	MembersCount      int       `json:"members_count"`
	Public            int       `json:"public"`
	Description       string    `json:"description"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	GrainsSum         int       `json:"grains_sum"`
	Serializer        string    `json:"_serializer"`
}

func (resp *GroupCreateResponse) PrettyString() string {
	return pretty.Struct(&resp.Data)
}

func (resp *GroupCreateResponse) String() string {
	data, _ := json.Marshal(resp)
	return string(data)
}

// GroupCreate 创建 Group
// curl --location --request POST 'https://www.yuque.com/api/v2/groups' \
// --header 'X-Auth-Token: abcd' \
// --header 'Content-Type: application/json' \
// --data-raw '{
//    "name": "test-lion",
//    "login": "test-lion",
//    "description": "test-lion"
// }'
func (c *Client) GroupCreate(req *GroupCreateRequest) (*GroupCreateResponse, error) {
	var (
		response GroupCreateResponse
		apiErr   YuqueError
	)
	resp, err := c.Client.R().
		SetBody(req).
		SetError(&apiErr).
		SetResult(&response).
		Post("groups")
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() {
		return &response, nil
	}

	return nil, &apiErr
}
