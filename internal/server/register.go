package server

import (
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/initialize"
)

func (receiver *HTTP) Register(controllers *initialize.Controllers) *HTTP {
	// groupVotes := receiver.echo.Group("/debt")
	// {
	// 	groupVotes.GET("/all", controllers.VotesController.GetAllVotess)
	// 	groupVotes.GET("/:id", controllers.VotesController.GetVotesByID)
	// 	groupVotes.POST("/search", controllers.VotesController.GetVotess)
	// }

	return receiver
}
