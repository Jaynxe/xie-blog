package model

import (
	"time"

	"github.com/Jaynxe/xie-blog/model/ctype"
)

type User struct {
	ID       int64  `gorm:"column:id" json:"id,omitempty"`
	Name     string `gorm:"column:name" json:"name,omitempty"`
	Password string `gorm:"column:password" json:"password,omitempty"`
	Email    string `gorm:"column:email" json:"email,omitempty"`
	NickName string `gorm:"column:nickName" json:"nick_name,omitempty"`
	IP       string `gorm:"column:ip" json:"ip,omitempty"`

	Role            string    `gorm:"column:role" json:"role,omitempty"` //权限(管理员，普通用户)
	Sex             string    `gorm:"column:sex" json:"sex,omitempty"`
	Avatar          string    `gorm:"type:varchar(256);comment:头像" json:"avatar,omitempty"` //头像地址
	Articles        []Article `gorm:"foreignKey:UserID" json:"articles,omitempty"`
	CollectArticles []Article `gorm:"many2many:user_collect_articles" json:"collect_articles,omitempty"`
	CreatedAt       time.Time `gorm:"column:createdAt" json:"created_at,omitempty"`
	UpdatedAt       time.Time `gorm:"column:updatedAt" json:"Updated_at,omitempty"`
}
type Article struct {
	ID           int64     `json:"id,omitempty"`
	Title        string    `json:"title,omitempty"`
	Content      string    `gorm:"type:text" json:"content,omitempty"`
	LookCount    uint      `json:"look_count,omitempty"`    //浏览量
	CommentCount uint      `json:"comment_count,omitempty"` //评论数
	LikeCount    uint      `json:"like_count,omitempty"`    //点赞数
	User         User      `gorm:"foreignKey:UserID" json:"-,omitempty"`
	UserID       int64     `json:"user_id,omitempty"`
	Comments     []Comment `gorm:"foreignKey:ArticleID" json:"comments,omitempty"`
	Tags         []Tag     `gorm:"many2many:article_tags" json:"tags,omitempty"`
	CategoryID   int64     `json:"category_id,omitempty"`
	Image        Image     `gorm:"foreignKey:ImageID" json:"image,omitempty"` //文章的封面
	ImageID      int64     `json:"image_id,omitempty"`
	CreatedAt    time.Time `gorm:"column:createdAt" json:"created_at,omitempty"`
	UpdatedAt    time.Time `gorm:"column:updatedAt" json:"Updated_at,omitempty"`
}
type Comment struct {
	ID        int64     `json:"id,omitempty"`
	Content   string    `gorm:"type:text" json:"content,omitempty"`
	ArticleID uint      `json:"article_id,omitempty"`
	UserID    int64     `json:"user_id,omitempty"`
	User      User      `json:"user,omitempty"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"Updated_at,omitempty"`
}

type Tag struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `gorm:"unique" json:"name,omitempty"`
	Articles  []Article `gorm:"many2many:article_tags" json:"articles,omitempty"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"Updated_at,omitempty"`
}
type Category struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `gorm:"unique" json:"name,omitempty"`
	Articles  []Article `gorm:"foreignKey:CategoryID" json:"articles,omitempty"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"Updated_at,omitempty"`
}
type Image struct {
	ID             int64                `json:"id,omitempty"`
	URL            string               `json:"url,omitempty"`
	Hash           string               `json:"hash,omitempty"` //判断图片是否重复
	Name           string               `json:"name,omitempty"`
	ImageStoreType ctype.ImageStoreType `gorm:"default:1" json:"image_store_type,omitempty"`
	CreatedAt      time.Time            `gorm:"column:createdAt" json:"created_at,omitempty"`
	UpdatedAt      time.Time            `gorm:"column:updatedAt" json:"Updated_at,omitempty"`
}

// MenuItem represents an item in the menu
type MenuItem struct {
	ID        int64      `gorm:"primaryKey" json:"id,omitempty"`
	Title     string     `json:"title,omitempty"`                               // 菜单项标题
	URL       string     `json:"url,omitempty"`                                 // 菜单项链接
	ParentID  *int64     `json:"parent_id,omitempty"`                           // 父菜单项ID，用于构建多级菜单
	Children  []MenuItem `gorm:"foreignKey:ParentID" json:"children,omitempty"` // 子菜单项列表
	Sort      int        `json:"sort,omitempty"`                                //菜单的排序
	CreatedAt time.Time  `gorm:"column:createdAt" json:"created_at,omitempty"`
	UpdatedAt time.Time  `gorm:"column:updatedAt" json:"Updated_at,omitempty"`
}
