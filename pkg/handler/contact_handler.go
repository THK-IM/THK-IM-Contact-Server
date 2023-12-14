package handler

import "github.com/thk-im/thk-im-contact-server/pkg/app"

func RegisterContactApiHandlers(appCtx *app.Context) {
	httpEngine := appCtx.HttpEngine()

	contactGroup := httpEngine.Group("/contact")
	contactGroup.GET("/friend", queryFriendList(appCtx))
	contactGroup.POST("/friend/apply", appFriendApply(appCtx))
	contactGroup.POST("/friend/apply/review", reviewFriendApply(appCtx))

	contactGroup.POST("/follow", followUser(appCtx))
	contactGroup.DELETE("/follow", unFollowUser(appCtx))
	contactGroup.POST("/black", followUser(appCtx))
	contactGroup.DELETE("/black", unFollowUser(appCtx))
}
