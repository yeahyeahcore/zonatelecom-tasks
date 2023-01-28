package server

import (
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/initialize"
)

func (receiver *HTTP) Register(controllers *initialize.Controllers) *HTTP {
	groupVotes := receiver.echo.Group("/voting")
	{
		groupVotes.POST("", controllers.VoteController.CreateVote)
	}

	return receiver
}
