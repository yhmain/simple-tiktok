package controller

import (
	"github.com/yhmain/simple-tiktok/model"
	"github.com/yhmain/simple-tiktok/service"
)

var DemoVideos = []model.Video{
	{
		Id:            1,
		User:          DemoUsers[0],
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
	{
		Id:            2,
		User:          DemoUsers[1],
		PlayUrl:       service.PREFIX_VIDEOS + "bear.mp4",
		CoverUrl:      service.PREFIX_COVERS + "bear.jpg",
		FavoriteCount: 2,
		CommentCount:  2,
		IsFavorite:    false,
	},
	{
		Id:            3,
		User:          DemoUsers[2],
		PlayUrl:       service.PREFIX_VIDEOS + "maxclub.mp4",
		CoverUrl:      service.PREFIX_COVERS + "long.jpeg",
		FavoriteCount: 3,
		CommentCount:  3,
		IsFavorite:    false,
	},
}

var DemoComments = []model.Comment{
	{
		Id:         1,
		User:       DemoUsers[0],
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUsers = []model.User{
	{
		Id:            1,
		NickName:      "TestUser",
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	},
	{
		Id:            2,
		NickName:      "HaHaUser",
		FollowCount:   2,
		FollowerCount: 2,
		IsFollow:      false,
	},
	{
		Id:            3,
		NickName:      "WoDeUser",
		FollowCount:   3,
		FollowerCount: 3,
		IsFollow:      false,
	},
}
