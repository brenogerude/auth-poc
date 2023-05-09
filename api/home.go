package api

import (
	"net/http"
	"oauth-poc/model"
)

type GetUserOutput struct {
	Id        string `permissions:"support" json:"id,omitempty"`
	FirstName string `permissions:"support,client" json:"firstName,omitempty"`
	LastName  string `permissions:"support,client" json:"lastName,omitempty"`
	CompanyId string `permissions:"support" json:"companyId,omitempty"`
	SSN       string `permissions:"support" json:"ssn,omitempty"`
}

var homeRoutes = RouterGroup{
	Get("/", GetHandler, AuthGuard(model.AdminRole, model.SupportRole)),
	Get("/users/:id", GetUsersHandler, AuthGuard(model.AdminRole, model.SupportRole)),
}

func GetHandler(request *http.Request, response *Response) {
	output := GetUserOutput{
		Id:        "1",
		FirstName: "Breno",
		LastName:  "Gerude",
		CompanyId: "1",
		SSN:       "123",
	}
	body := BuildBody(output)
	response.Success(body)
}

func GetUsersHandler(request *http.Request, response *Response) {
	output := GetUserOutput{
		Id:        "1",
		FirstName: "Breno",
		LastName:  "Gerude",
		CompanyId: "1",
		SSN:       "123",
	}
	body := BuildBody(output)
	response.Success(body)
}
