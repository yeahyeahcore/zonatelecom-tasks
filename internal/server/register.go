package server

import (
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/initialize"
)

func (receiver *HTTP) Register(controllers *initialize.Controllers) *HTTP {
	groupHealth := receiver.echo.Group("/_health")
	{
		groupHealth.GET("/check", controllers.HealthCheckController.Health)
		groupHealth.GET("/ready", controllers.HealthCheckController.Ready)
	}
	groupVotes := receiver.echo.Group("/voting")
	{
		groupVotes.POST("", controllers.VoteController.CreateVote)
	}

	return receiver
}
