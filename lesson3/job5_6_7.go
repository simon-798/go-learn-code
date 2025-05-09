package lesson3

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

/**
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章），
Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
*/

func Job5Method() {

	err := db.AutoMigrate(&BlogsUser{}, &Post{}, &Comment{})
	if err != nil {
		fmt.Println("数据库迁移失败:", err)
	}
	fmt.Println("数据库表创建成功")

	//插入测试数据
	blogsUsers := []BlogsUser{
		{Name: "张三", Posts: []Post{
			{Title: "文章1", Content: "内容1", Comments: []Comment{
				{Content: "评论1"},
			}, CommentNum: 1},
		}},
		{Name: "李四", Posts: []Post{
			{Title: "文章2", Content: "内容2", Comments: []Comment{
				{Content: "评论2"},
				{Content: "评论3"},
			}, CommentNum: 2},
			{Title: "文章3", Content: "内容3", Comments: []Comment{
				{Content: "评论4"},
				{Content: "评论5"},
				{Content: "评论6"},
			}, CommentNum: 3},
		}},
	}

	db.Create(&blogsUsers)

}

type BlogsUser struct {
	ID         uint64 `gorm:"size:16;PRIMARY_KEY;AUTO_INCREMENT"`
	Name       string `gorm:"size:255"`
	Posts      []Post // 一对多关系
	PostsCount uint64 `gorm:"size:11,default:0"`
}

type Post struct {
	ID          uint64    `gorm:"size:16;PRIMARY_KEY;AUTO_INCREMENT"`
	Title       string    `gorm:"size:255"`
	Content     string    `gorm:"size:255"`
	BlogsUserID uint64    `gorm:"size:16"` // 外键
	Comments    []Comment // 一对多关系
	CommentNum  uint64    `gorm:"size:11"`       // 评论数量
	Status      string    `gorm:"default:'有评论'"` // 评论状态
}

type Comment struct {
	ID        uint64    `gorm:"size:16;PRIMARY_KEY;AUTO_INCREMENT"`
	Content   string    `gorm:"type:text"`
	PostID    uint64    `gorm:"size:16"` // 外键
	CreatedAt time.Time `gorm:"type:timestamp;DEFAULT:CURRENT_TIMESTAMP"`
}

/**
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
*/

func Job6Method() {

	//使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	//预加载方式,预加载的名字是外键的属性名
	var blogsUser BlogsUser
	result := db.Preload("Posts.Comments").Take(&blogsUser, 1)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println("blogsUser:", blogsUser)

	//db.Preload("Posts.Comments").Take(&blogsUser, "Title=?", "文章2")
	//db.Preload("Posts.Comments").First(&blogsUser,1)

	//使用Gorm查询评论数量最多的文章信息
	var post Post
	result = db.Select("posts.*,count(comments.id) commentNum").
		Joins("left join comments on comments.post_id = posts.id").
		Group("posts.title").
		Order("commentNum desc").First(&post)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	//SELECT posts.*,count(comments.id) commentNum FROM `posts` left join comments on comments.post_id = posts.id
	//GROUP BY `posts`.`title` ORDER BY commentNum desc,`posts`.`id` LIMIT 1

	//原生sql实现
	/*db.Raw("select a.*,count(b.id) commentNum from posts a left join comments b on a.id = b.post_id " +
	"group by a.id order by count(b.id) desc limit 1").Scan(&post)*/

	fmt.Println("post:", post)
}

/**
题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/

/*
*
在文章创建时自动更新用户的文章数量统计字段。
*/
func (post *Post) AfterCreate(tx *gorm.DB) (err error) {

	result := tx.Model(&BlogsUser{}).Where("id = ?", post.BlogsUserID).
		Update("posts_count", gorm.Expr("posts_count + ?", 1))
	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}
	fmt.Println("在文章创建时自动更新用户的文章数量统计字段。")
	return nil
}

func Job7Method() {
	//删除文章的评论
	var comment Comment
	db.Preload("&Posts").Take(&comment, 1)
	fmt.Println("comment:", comment)
	db.Delete(&comment)
}

/*
*
在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/
func (comment *Comment) AfterDelete(tx *gorm.DB) (err error) {

	var commentNum int64 //Count方法不能使用无符号整型uint接收,只能用有符号整型int接收
	tx.Model(&Comment{}).Where("post_id = ?", comment.PostID).Count(&commentNum)

	if commentNum == 0 {
		result := tx.Model(&Post{}).Where("id = ?", comment.PostID).Updates(Post{CommentNum: 0, Status: "无评论"})
		return result.Error
	}
	fmt.Println("在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为无评论")
	return nil
}
