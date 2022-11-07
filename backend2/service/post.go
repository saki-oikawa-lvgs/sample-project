package service

import (
	"github.com/saki-oikawa-lvgs/sample-project/backend2/graph/common"
	"github.com/saki-oikawa-lvgs/sample-project/backend2/graph/customTypes"
)

func PostsGet() ([]*customTypes.Post, error) {

	db := common.InitDb()
	// defer db.Close()

	var posts []*customTypes.Post
	err := db.Table("posts").Find(&posts).Error

	return posts, err
}

func PostCreate(input customTypes.NewPost) (*customTypes.Post, error) {
	
	db := common.InitDb()
	// defer db.Close()

	newPost := customTypes.Post{
		Text: input.Text,
		Done: input.Done,
	}

	err := db.Table("posts").Create(&newPost).Error

	return &newPost, err
}

func PostGetByUserID(userID int) ([]*customTypes.Post, error) {

	db := common.InitDb()
	// defer db.Close()

	var userPosts []*customTypes.Post
	err := db.Table("posts").Where("user_id = ?", userID).Find(&userPosts).Error

	return userPosts, err
}

func PostFindByID(postID int) (*customTypes.Post, error) {

	db := common.InitDb()
	// defer db.Close()

	var post customTypes.Post
	err := db.Table("posts").Where("id = ?",postID).First(&post).Error

	return &post, err
}