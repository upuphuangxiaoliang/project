package service

import (
	"html/template"
	"project/project1/config"
	"project/project1/dao"
	"project/project1/models"
)

func GetPostDetal(pid int) (*models.PostRes, error) {
	// 从数据库里面导出
	post, err := dao.GetPostById(pid)
	if err != nil {
		return nil, err

	}
	categoryName := dao.GetCategoryNameById(post.CategoryId)
	userName := dao.GetUserNameById(post.UserId)
	postMore := models.PostMore{
		post.Pid,
		post.Title,
		post.Slug,
		template.HTML(post.Content),
		post.CategoryId,
		categoryName,
		post.UserId,
		userName,
		post.ViewCount,
		post.Type,
		models.DateDay(post.CreateAt),
		models.DateDay(post.UpdateAt),
	}
	var postRes = &models.PostRes{
		// 数据大部分为config里面的
		config.Cfg.Viewer,
		config.Cfg.System,
		postMore,
	}
	return postRes, nil
}
