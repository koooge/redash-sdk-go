package redash

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"strconv"
)

// GET /api/users
type ListUsersInput struct {
	Page     int
	PageSize int
	Q        string
	Disabled bool
	Pending  bool
}

type ListUsersOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) ListUsers(input *ListUsersInput) *ListUsersOutput {
	path := "/api/users"

	if input != nil {
		values := url.Values{}
		if input.Page != 0 {
			values.Add("page", strconv.Itoa(input.Page))
		}
		if input.PageSize != 0 {
			values.Add("page_size", strconv.Itoa(input.PageSize))
		}
		if input.Q != "" {
			values.Add("q", input.Q)
		}
		if input.Disabled != false {
			values.Add("disabled", strconv.FormatBool(input.Disabled))
		}
		if input.Pending != false {
			values.Add("pending", strconv.FormatBool(input.Pending))
		}
		if len(values) > 0 {
			path = path + "?" + values.Encode()
		}
	}

	resp, err := c.get(path)
	if err != nil {
		return &ListUsersOutput{StatusCode: resp.StatusCode, Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &ListUsersOutput{
		StatusCode: resp.StatusCode,
		Body:       string(b),
	}
}

// POST /api/users
type CreateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserOutput struct {
	Body       string `json:"-"`
	StatusCode int    `json:"-"`
	Id         int    `json:"id"`
}

func (c *Client) CreateUser(input *CreateUserInput) *CreateUserOutput {
	path := "/api/users"

	body, err := json.Marshal(input)
	if err != nil {
		return &CreateUserOutput{Body: `{"error":"` + err.Error() + `"}`}
	}

	resp, err := c.post(path, string(body))
	if err != nil {
		return &CreateUserOutput{StatusCode: resp.StatusCode, Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	resBody := &CreateUserOutput{}
	if err := json.Unmarshal(b, &resBody); err != nil {
		return &CreateUserOutput{StatusCode: resp.StatusCode, Body: `{"error":"` + err.Error() + `"}`}
	}

	return &CreateUserOutput{
		StatusCode: resp.StatusCode,
		Body:       string(b),
		Id:         resBody.Id,
	}
}

// GET /api/users/{userId}
type GetUserInput struct {
	UserId int
}

type GetUserOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetUser(input *GetUserInput) *GetUserOutput {
	path := "/api/users/" + strconv.Itoa(input.UserId)

	resp, err := c.get(path)
	if err != nil {
		return &GetUserOutput{StatusCode: resp.StatusCode, Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetUserOutput{
		StatusCode: resp.StatusCode,
		Body:       string(b),
	}
}

// POST /api/users/{userId}
type UpdateUserInput struct {
	UserId int    `json:"-"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
}

type UpdateUserOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) UpdateUser(input *UpdateUserInput) *UpdateUserOutput {
	path := "/api/users/" + strconv.Itoa(input.UserId)

	body, err := json.Marshal(input)
	if err != nil {
		return &UpdateUserOutput{Body: `{"error":"` + err.Error() + `"}`}
	}

	resp, err := c.post(path, string(body))
	if err != nil {
		return &UpdateUserOutput{StatusCode: resp.StatusCode, Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	resBody := &UpdateUserOutput{}
	if err := json.Unmarshal(b, &resBody); err != nil {
		return &UpdateUserOutput{StatusCode: resp.StatusCode, Body: `{"error":"` + err.Error() + `"}`}
	}

	return &UpdateUserOutput{
		StatusCode: resp.StatusCode,
		Body:       string(b),
	}
}

// DELETE /api/users/{userId}
// POST /api/users/{userId}/invite
// POST /api/users/{userId}/reset_password
// POST /api/users/{userId}/generate_api_key
// POST /api/users/{userId}/disable
// DELETE /api/users/{userId}/disable
