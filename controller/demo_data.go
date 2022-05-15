package controller

const (
	SERVER_STATIC_PATH = "http://192.168.1.108:8081/"
)

var DemoVideos = []Video{
	{
		Id:            1,
		Author:        DemoUsers[0],
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
	{
		Id:            2,
		Author:        DemoUsers[1],
		PlayUrl:       SERVER_STATIC_PATH + "bear.mp4",
		CoverUrl:      SERVER_STATIC_PATH + "bear-1283347_1280.jpg",
		FavoriteCount: 2,
		CommentCount:  2,
		IsFavorite:    false,
	},
	{
		Id:            3,
		Author:        DemoUsers[2],
		PlayUrl:       SERVER_STATIC_PATH + "maxclub.mp4",
		CoverUrl:      SERVER_STATIC_PATH + "long.jpeg",
		FavoriteCount: 3,
		CommentCount:  3,
		IsFavorite:    false,
	},
}

var DemoComments = []Comment{
	{
		Id:         1,
		User:       DemoUsers[0],
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUsers = []User{
	{
		Id:            1,
		Name:          "TestUser",
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	},
	{
		Id:            2,
		Name:          "HaHaUser",
		FollowCount:   2,
		FollowerCount: 2,
		IsFollow:      false,
	},
	{
		Id:            3,
		Name:          "WoDeUser",
		FollowCount:   3,
		FollowerCount: 3,
		IsFollow:      false,
	},
}
