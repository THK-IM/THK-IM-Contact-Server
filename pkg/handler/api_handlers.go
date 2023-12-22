package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/thk-im/thk-im-base-server/conf"
	"github.com/thk-im/thk-im-base-server/middleware"
	"github.com/thk-im/thk-im-contact-server/pkg/app"
	userSdk "github.com/thk-im/thk-im-user-server/pkg/sdk"
)

func RegisterContactApiHandlers(appCtx *app.Context) {
	httpEngine := appCtx.HttpEngine()
	ipAuth := middleware.WhiteIpAuth(appCtx.Config().IpWhiteList, appCtx.Logger())
	userApi := appCtx.UserApi()
	userTokenAuth := userSdk.UserTokenAuth(userApi, appCtx.Logger())

	var authMiddleware gin.HandlerFunc
	if appCtx.Config().DeployMode == conf.DeployExposed {
		authMiddleware = userTokenAuth
	} else if appCtx.Config().DeployMode == conf.DeployBackend {
		authMiddleware = ipAuth
	} else {
		panic(errors.New("check deployMode conf"))
	}

	contactGroup := httpEngine.Group("/contact")
	contactGroup.Use(authMiddleware)
	{
		contactGroup.GET("", queryContactList(appCtx))
		contactGroup.POST("/friend/apply", appFriendApply(appCtx))
		contactGroup.POST("/friend/apply/review", reviewFriendApply(appCtx))
		contactGroup.POST("/follow", followUser(appCtx))
		contactGroup.DELETE("/follow", unFollowUser(appCtx))
		contactGroup.POST("/black", addBlack(appCtx))
		contactGroup.DELETE("/black", removeBlack(appCtx))
	}
}