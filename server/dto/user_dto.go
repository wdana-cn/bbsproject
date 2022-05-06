package dto

import "github.com/wdana-cn/ginpro/model"

type UserDto struct {
	UserId    uint   `json:"UserId"`
	UserName  string `json:"name"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
}

type PostDto struct{
	PostTitle	string	`json:"PostTitle"`
	PostUID		uint 	`json:PostUID`
	PostContent	string	`json:PostContent`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		UserId:    user.ID,
		UserName:  user.Name,
		Telephone: user.Telephone,
		Email:     user.Email,
	}
}
func ToPostDto(post model.Postdata) PostDto{
	return PostDto{
		PostTitle:		post.Title,
		PostUID:		post.UserId,
		PostContent:	post.Content,
	}
	
}
