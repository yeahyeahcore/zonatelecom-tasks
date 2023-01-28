package server

import (
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/initialize"
)

func (receiver *HTTP) Register(controllers *initialize.Controllers) *HTTP {
	// groupDebt := receiver.echo.Group("/debt")
	// {
	// 	groupDebt.GET("/all", controllers.DebtController.GetAllDebts)
	// 	groupDebt.GET("/:id", controllers.DebtController.GetDebtByID)
	// 	groupDebt.POST("/search", controllers.DebtController.GetDebts)
	// }

	return receiver
}
