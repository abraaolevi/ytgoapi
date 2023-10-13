package e2e

import "net/http"

type PostSuccessfulSuite struct {
	api      ApiClient
	username string
	body     string
}

func NewPostSuccessfulSuite() PostSuccessfulSuite {
	return PostSuccessfulSuite{
		api: ApiClient{
			baseUrl: "http://localhost:8080",
		},
		username: "my_username_test",
		body:     "My Body Description Test",
	}
}

func (suite *PostSuccessfulSuite) Run() {
	id := suite.create()
	suite.read(id)
	suite.delete(id)
}

func (suit *PostSuccessfulSuite) create() string {
	payload := map[string]string{
		"username": suit.username,
		"body":     suit.body,
	}

	resp := suit.api.Post("/posts", payload)
	if resp.StatusCode != http.StatusCreated {
		panic("Invalid Status Code for creation")
	}

	res := suit.api.ParseBody(resp)
	id := res["id"].(string)

	if id == "" {
		panic("Invalid ID")
	}

	if res["username"].(string) != payload["username"] {
		panic("Invalid Username")
	}

	if res["body"].(string) != payload["body"] {
		panic("Invalid Body")
	}

	if res["created_at"].(string) == "0001-01-01T00:00:00Z" {
		panic("Invalid CreatedAt")
	}

	return id
}

func (suit *PostSuccessfulSuite) read(id string) {
	resp := suit.api.Get("/posts/" + id)
	if resp.StatusCode != http.StatusOK {
		panic("Invalid Status Code")
	}

	res := suit.api.ParseBody(resp)
	if res["id"].(string) != id {
		panic("Invalid ID")
	}

	if res["username"].(string) != suit.username {
		panic("Invalid Username")
	}

	if res["body"].(string) != suit.body {
		panic("Invalid Body")
	}
}

func (suit *PostSuccessfulSuite) delete(id string) {
	resp := suit.api.Delete("/posts/" + id)
	if resp.StatusCode != http.StatusNoContent {
		panic("Invalid Status Code")
	}
}
