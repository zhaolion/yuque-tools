package api

import (
	"fmt"
	"time"
)

type UserCurrentResponse struct {
	Data UserDetail `json:"data"`
}

func (resp *UserCurrentResponse) String() string {
	return fmt.Sprintf("%+v", resp.Data)
}

type UserDetail struct {
	ID               int       `json:"id"`
	Type             string    `json:"type"`
	SpaceID          int       `json:"space_id"`
	AccountID        int       `json:"account_id"`
	Login            string    `json:"login"`
	Name             string    `json:"name"`
	AvatarURL        string    `json:"avatar_url"`
	BooksCount       int       `json:"books_count"`
	PublicBooksCount int       `json:"public_books_count"`
	FollowersCount   int       `json:"followers_count"`
	FollowingCount   int       `json:"following_count"`
	Public           int       `json:"public"`
	Description      string    `json:"description"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Serializer       string    `json:"_serializer"`
}

// UserCurrent 获取认证的用户的个人信息
// curl --location --request GET 'https://www.yuque.com/api/v2/user' \
//	--header 'X-Auth-Token: abcd'
func (c *Client) UserCurrent() (*UserCurrentResponse, error) {
	var (
		response UserCurrentResponse
		apiErr   YuqueError
	)
	resp, err := c.Client.R().
		SetError(&apiErr).
		SetResult(&response).
		Get("user")
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() {
		return &response, nil
	}

	return nil, &apiErr
}
