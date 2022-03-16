package dao

import (
	"log"
	"project/project1/models"
)

// 指定文章的数量
func CountGetAllPostByCategoryId(cID int) (count int) {

	rows := DB.QueryRow("select count(1) from blog_post where category_id = ?", cID)
	_ = rows.Scan(&count)
	return

}

// 所有文章的数量
func CountGetAllPost() (count int) {

	rows := DB.QueryRow("select count(1) from blog_post")
	_ = rows.Scan(&count)
	return

}

func GetPostPage(page, pagesize int) ([]models.Post, error) {
	//有分页操作
	page = (page - 1) * pagesize
	rows, err := DB.Query("select * from blog_post limit ?,?", page, pagesize)
	if err != nil {
		log.Println("查询出错：", err)
		return nil, err
	}

	// 没有报错
	var posts []models.Post
	// 把数据取出来 取到category上面
	for rows.Next() {
		for rows.Next() {
			var post models.Post
			err := rows.Scan(
				// z这里是数据库里面的字段
				&post.Pid,
				&post.Title,
				&post.Content,
				&post.Markdown,
				&post.CategoryId,
				&post.UserId,
				&post.ViewCount,
				&post.Type,
				&post.Slug,
				&post.CreateAt,
				&post.UpdateAt,
			)
			if err != nil {
				return nil, err
			}
			posts = append(posts, post)
		}
		// 返回

	}
	return posts, nil
}

func GetPostPageByCategoryId(cId, page, pagesize int) ([]models.Post, error) {
	//有分页操作
	page = (page - 1) * pagesize
	rows, err := DB.Query("select * from blog_post where category_id = ? limit ?,?", cId, page, pagesize)
	if err != nil {
		log.Println("查询出错：", err)
		return nil, err
	}

	// 没有报错
	var posts []models.Post
	// 把数据取出来 取到category上面
	for rows.Next() {
		for rows.Next() {
			var post models.Post
			err := rows.Scan(
				// z这里是数据库里面的字段
				&post.Pid,
				&post.Title,
				&post.Content,
				&post.Markdown,
				&post.CategoryId,
				&post.UserId,
				&post.ViewCount,
				&post.Type,
				&post.Slug,
				&post.CreateAt,
				&post.UpdateAt,
			)
			if err != nil {
				return nil, err
			}
			posts = append(posts, post)
		}
		// 返回

	}
	return posts, nil
}

func GetPostById(pid int) (models.Post, error) {
	rows := DB.QueryRow("select * from blog_post where pid = ?", pid)
	var post models.Post
	//有可能获取到的是空的 返回错误
	if rows.Err() != nil {
		return post, rows.Err()
	}

	err := rows.Scan(
		// z这里是数据库里面的字段
		&post.Pid,
		&post.Title,
		&post.Content,
		&post.Markdown,
		&post.CategoryId,
		&post.UserId,
		&post.ViewCount,
		&post.Type,
		&post.Slug,
		&post.CreateAt,
		&post.UpdateAt,
	)
	if err != nil {
		return post, rows.Err()
	}
	return post, nil

}
