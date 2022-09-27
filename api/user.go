package api

import (
	"fmt"
	"time"

	"github.com/imroc/req/v3"
	"github.com/zhaolion/yuque-tools/api/pretty"
)

type UserResponse struct {
	Data UserDetail `json:"data"`
}

func (resp *UserResponse) String() string {
	return pretty.Struct(&resp.Data)
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

// User 获取认证的用户的个人信息
func (c *Client) User(query ...string) (*UserResponse, error) {
	var (
		response UserResponse
		apiErr   YuqueError
		resp     *req.Response
		err      error
	)
	if len(query) == 0 {
		// curl --location --request GET 'https://www.yuque.com/api/v2/user' \
		//	--header 'X-Auth-Token: abcd'
		resp, err = c.Client.R().
			SetError(&apiErr).
			SetResult(&response).
			Get("user")
	} else {
		// 获取单个用户信息
		// curl --location --request GET 'https://www.yuque.com/api/v2/users/114999' \
		//--header 'X-Auth-Token: abcd'
		resp, err = c.Client.R().
			SetError(&apiErr).
			SetResult(&response).
			Get(fmt.Sprintf("users/%s", query[0]))
	}
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() {
		return &response, nil
	}

	return nil, &apiErr
}
