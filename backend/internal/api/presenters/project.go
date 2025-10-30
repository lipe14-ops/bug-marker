	package presenters

	import (
		"github.com/gofiber/fiber/v2"
	)

	type Project struct {
		ID    			 string    `json:"id" bson:"_id"`
		Owner        string    `json:"owner" bson:"owner"`
		Participants []string  `json:"participants" bson:"participants"`
		Title        string    `json:"title" bson:"title"`
		Description  string    `json:"description" bson:"description"`
	}

	func ProjectSuccessResponse(project *Project) *fiber.Map {
		return &fiber.Map {
			"messages": []string{
				"success",
			},
			"data" : []*Project{
				project,
			},
		}
	}

	func ProjectsSuccessResponse(projects []*Project) *fiber.Map {
		return &fiber.Map {
			"messages": []string{
				"success",
			},
			"data" : projects,
		}
	}


	func ProjectMessageResponse(message string) *fiber.Map {
		return &fiber.Map {
			"messages": []string{
				message,
			},
			"data" : []string{},
	}
}

func ProjectErrorResponse(err error) *fiber.Map {
	return &fiber.Map {
		"messages": []string{
			err.Error(),
		},
		"data" : []string{},
	}
}
