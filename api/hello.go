package api

type HelloResponse struct {
	Data Hello `json:"data"`
}

type Hello struct {
	Message *string `json:"message"`
}

// Hello test message 用户认证
// curl -H "X-Auth-Token: gCmkIlgAtuc3vFwpLfeM1w==" https://www.yuque.com/api/v2/hello
func (c *Client) Hello() (*HelloResponse, error) {
	var (
		response HelloResponse
		apiErr   YuqueError
	)
	resp, err := c.Client.R().
		SetError(&apiErr).
		SetResult(&response).
		Get("hello")
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() {
		return &response, nil
	}

	return nil, &apiErr
}
