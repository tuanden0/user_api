package models

import "log"

type Article struct {
	Id      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (a Article) All() (articles []*Article, e error) {

	if err := DB.Find(&articles).Error; err != nil {
		log.Printf("unable to get all articles! %v\n", err.Error())
		return nil, err
	}

	return articles, nil
}

func (a Article) Get(id interface{}) (article *Article, e error) {

	if err := DB.Where("id = ?", id).First(&article).Error; err != nil {
		log.Printf("unable to get article! %v\n", err.Error())
		return nil, err
	}

	return article, nil
}
