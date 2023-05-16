package api

import (
	"github.com/gin-gonic/gin"
)

type Inner struct {
	Id        string `json:"id,omitempty"`
	CompanyId string `json:"companyId,omitempty"`
}
type GetUserOutput struct {
	Id         string  `json:"id,omitempty"`
	FirstName  string  `json:"firstName,omitempty"`
	LastName   string  `json:"lastName,omitempty"`
	CompanyId  string  `json:"companyId,omitempty"`
	SSN        string  `json:"ssn,omitempty"`
	Inner      Inner   `json:"inner,omitempty"`
	InnerSlice []Inner `json:"inners,omitempty"`
}

var homeRoutes = RouterGroup{
	Get("/users", GetHandler),
	Get("/users/:id", GetUsersHandler),
}

func GetHandler(context *gin.Context) {
	output := GetUserOutput{
		Id:        "1",
		FirstName: "Breno",
		LastName:  "Gerude",
		CompanyId: "1",
		SSN:       "123",
		Inner: Inner{
			Id:        "2",
			CompanyId: "2",
		},
		InnerSlice: []Inner{
			{
				Id:        "1",
				CompanyId: "2",
			},
			{
				Id:        "2",
				CompanyId: "3",
			},
			{
				Id:        "3",
				CompanyId: "4",
			},
			{
				Id:        "4",
				CompanyId: "5",
			},
		},
	}
	context.JSON(200, &output)
}

func GetUsersHandler(context *gin.Context) {
	output := []GetUserOutput{
		{
			Id:        "1",
			FirstName: "Breno",
			LastName:  "Gerude",
			CompanyId: "1",
			SSN:       "123",
			Inner: Inner{
				Id:        "2",
				CompanyId: "2",
			},
		},
		{
			Id:        "2",
			FirstName: "Wanny",
			LastName:  "Gerude",
			CompanyId: "1",
			SSN:       "124",
			Inner: Inner{
				Id:        "2",
				CompanyId: "2",
			},
		},
		{
			Id:        "3",
			FirstName: "Maya",
			LastName:  "Gerude",
			CompanyId: "1",
			SSN:       "125",
			Inner: Inner{
				Id:        "2",
				CompanyId: "2",
			},
		},
		{
			Id:        "4",
			FirstName: "Vitor",
			LastName:  "Gerude",
			CompanyId: "1",
			SSN:       "126",
			Inner: Inner{
				Id:        "2",
				CompanyId: "2",
			},
		},
	}
	context.JSON(200, &output)
}
