package dao

import (
	"log"
	"project/project1/models"
)

// 这一条是根据cid返回名字
func GetCategoryNameById(cId int) string {
	row := DB.QueryRow("select name from blog_category where cid=?", cId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	// 查询到之后取出来
	var categoryName string
	_ = row.Scan(&categoryName)
	return categoryName
}

// 返回的是 category结构体中的信息
func GetAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Println("查询出错：", err)
		return nil, err
	}

	// 没有报错
	var categorys []models.Category
	// 把数据取出来 取到category上面
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Println("GetAllCategory 取值出错:", err)
			return nil, err
		}
		// 取出来之后 加入到category上面
		categorys = append(categorys, category)
	}
	// 返回
	return categorys, nil
}
