	package presenters

	import (
		"github.com/gofiber/fiber/v2"
	)

	type User struct {
		ID    string    `json:"id" bson:"_id"`
		Name  string    `json:"name" bson:"name"`
		Email string    `json:"email" bson:"email"`
    HasBeenSoftDeleted bool `json:"hasBeenSoftDeleted" bson:"hasBeenSoftDeleted"`
	}

	func UserSuccessResponse(user *User) *fiber.Map {
		return &fiber.Map {
			"messages": []string{
				"success",
			},
			"data" : []*User{
				user,
			},
		}
	}

	func UserMessageResponse(message string) *fiber.Map {
		return &fiber.Map {
			"messages": []string{
				message,
			},
			"data" : []string{},
	}
}

func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map {
		"messages": []string{
			err.Error(),
		},
		"data" : []string{},
	}
}
