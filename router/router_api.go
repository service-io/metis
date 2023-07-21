// Package router
// @author tabuyos
// @since 2023/7/10
// @description router
package router

import (
	"metis/api/survey"
	"metis/api/welcome"
	"metis/middleware"
)

func setApiRouter() {
	welcomeGroup := baseRouter.Group("/welcome")
	welcomeGroup.Use(middleware.CheckAuth(), middleware.CheckRBAC())
	{
		welcomeGroup.GET("/hello", welcome.Hello)
		welcomeGroup.GET("/whoami", welcome.WhoAmI)
	}

	v1Group := baseRouter.Group("/api/v1")
	{
		surveyGroup := v1Group.Group("/survey")
		{
			surveyGroup.GET("/list", survey.List)
			surveyGroup.GET("/detail/:id", survey.Detail)
		}
	}
}
