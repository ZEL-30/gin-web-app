package assembler

import (
	"fmt"
	"net/http"

	"github.com/ZEL-30/gin-web-app/entity"
	rep "github.com/ZEL-30/gin-web-app/representation"
)

type UserAssembler struct{}

func NewUserAssembler() UserAssembler {
	return UserAssembler{}
}

func (s *UserAssembler) ToData(rep rep.User) *entity.User {
	return &entity.User{
		Name:     rep.Name,
		Password: rep.Password,
	}
}

func (s *UserAssembler) ToRepresentation(data entity.User) *rep.User {
	return &rep.User{
		Base: rep.Base{
			ID:        data.ID,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,

			Links: []rep.ResourceLink{
				{
					Rel:    "self",
					Method: http.MethodGet,
					Href:   fmt.Sprintf("/books/%s", data.ID),
				},
				{
					Rel:    "add-book",
					Method: http.MethodPost,
					Href:   "/books",
				},
				{
					Rel:    "edit-book",
					Method: http.MethodPut,
					Href:   fmt.Sprintf("/books/%s", data.ID),
				},
				{
					Rel:    "delete-book",
					Method: http.MethodDelete,
					Href:   fmt.Sprintf("/books/%s", data.ID),
				},
			},
		},

		Name:     data.Name,
		Password: data.Password,
	}
}
