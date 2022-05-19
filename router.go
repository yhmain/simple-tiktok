package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yhmain/simple-tiktok/controller"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)                                               //视频流接口：2022/05/16完成
	apiRouter.GET("/user/", controller.JWTAuthUserToken(), controller.UserInfo)            //用户信息：2022/05/16完成
	apiRouter.POST("/user/register/", controller.Register)                                 //用户注册：2022/05/16完成
	apiRouter.POST("/user/login/", controller.Login)                                       //用户登录：2022/05/16完成
	apiRouter.POST("/publish/action/", controller.Publish)                                 //用户发布新视频：
	apiRouter.GET("/publish/list/", controller.JWTAuthUserToken(), controller.PublishList) //用户发布的视频列表：2022/05/16完成

	// extra apis - I
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
}
