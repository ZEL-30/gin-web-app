package assembler

import (
	"fmt"
	"net/http"

	"github.com/ZEL-30/gin-web-app/entity"
	"github.com/ZEL-30/gin-web-app/representation"
)

type UserAssembler struct{}

func NewUserAssembler() UserAssembler {
	return UserAssembler{}
}

// ToData 将 representation.User 转换为 entity.User
func (s *UserAssembler) ToData(rep representation.User) *entity.User {
	return &entity.User{
		Name:     rep.Name,
		Password: rep.Password,
		Email:    rep.Email,
	}
}

// ToRepresentation 将 entity.User 转换为 representation.User
func (s *UserAssembler) ToRepresentation(data entity.User) *representation.User {
	return &representation.User{
		Base: representation.Base{
			ID:        data.ID,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,

			Links: []representation.ResourceLink{
				{
					Rel:    "self",
					Method: http.MethodGet,
					Href:   fmt.Sprintf("/users/%d", data.ID),
				},
				{
					Rel:    "add-user",
					Method: http.MethodPost,
					Href:   "/users",
				},
				{
					Rel:    "edit-user",
					Method: http.MethodPut,
					Href:   fmt.Sprintf("/users/%d", data.ID),
				},
				{
					Rel:    "delete-user",
					Method: http.MethodDelete,
					Href:   fmt.Sprintf("/users/%d", data.ID),
				},
			},
		},

		Name:     data.Name,
		Password: data.Password,
		Email:    data.Email,
	}
}
