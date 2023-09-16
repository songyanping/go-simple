package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var DB *gorm.DB

// 文章表
type Article struct {
	Id         int      `json:"id"`
	Title      string   `json:"title"`
	CategoryID int      `json:"category_id"`
	Category   Category `json:"category";gorm:"foreignkey:CategoryID"` //指定关联外键
	Tags       []Tag    `gorm:"many2many:article_tag" json:"tags"`     //多对多关系.
	//article_tag表默认article_id字段对应article表id.tag_id字段对应tag表id
	//可以把对应sql日志打印出来,便于调试
}

// 文章_标签中间表
type ArticleTag struct {
	Id        int    `json:"id" `
	ArticleId string `json:"article_id"`
	TagId     string `json:"tag_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// 标签表
type Tag struct {
	Id      int    `json:"id" `
	TagName string `json:"tag_name"`
}

// 分类表
type Category struct {
	ID           int       `json:"id"`
	CategoryName string    `json:"category_name"`
	Status       int       `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func initGormDb() *gorm.DB {
	dsn := "root:Devops2019@tcp(10.158.215.32:3306)/simple-db?charset=utf8mb4&parseTime=True&loc=Local"
	dbs, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//如果不想看到日志sql，这里可以关掉
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表明不加s
		},
	})
	if err != nil {
		panic("gorm failed to connect Mysql...")
	}
	DB = dbs
	return dbs
}

// CreateTable 创建表
func CreateTable(db *gorm.DB) {
	err := db.AutoMigrate(&Article{}, &Category{}, Tag{}, ArticleTag{})
	if err != nil {
		log.Println("failed to create table...")
		return
	}
}

// 远程一对多.一对一
func (a *Article) ListArticle(title string) (Article, error) {
	var article Article
	err := DB.Where("title like ?", "%"+title+"%").Preload("Category", "status = ?", 1).Preload("Tags").First(&article).Error
	if err != nil {
		fmt.Sprintf(err.Error())
	}
	fmt.Println(article)

	return article, nil

}

func main() {
	db := initGormDb()
	fmt.Println(db)
	//1、创建表
	CreateTable(db)
	a := Article{}
	a.ListArticle("abc")

}
