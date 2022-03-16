package service

import (
	"html/template"
	"project/project1/config"
	"project/project1/dao"
	"project/project1/models"
)

func GetPostByCategoryId(cId, page, pagesize int) (*models.CategoryResponse, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	// 由于前端需要的比数据库里面的东西还要多 所以不满足 还需要增加缺少的
	posts, err := dao.GetPostPageByCategoryId(cId, page, pagesize)
	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)

		// rune 是因为要按照中文字符来
		content := []rune(post.Content)
		if len(content) > 100 {
			// 如果内容长度大于100  切一下
			content = content[0:100]
		}
		postMore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(content),
			post.CategoryId,
			categoryName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			models.DateDay(post.CreateAt),
			models.DateDay(post.UpdateAt),
		}
		// 放进去
		postMores = append(postMores, postMore)
	}
	// 假数据
	total := dao.CountGetAllPostByCategoryId(cId)
	pagesCount := (total-1)/10 + 1
	var pages []int
	for i := 0; i < pagesCount; i++ {
		pages = append(pages, i+1)
	}
	var hr = &models.HomeResponse{
		// 这里的 在config.文件已经弄好了 直接拿到
		config.Cfg.Viewer,
		categorys,
		postMores,
		total,
		page,
		pages,
		page != pagesCount,
	}
	categoryName := dao.GetCategoryNameById(cId)
	categoryResponse := &models.CategoryResponse{
		hr,
		categoryName,
	}
	return categoryResponse, nil
}
