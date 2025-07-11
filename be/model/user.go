package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	Id       int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Role     string `gorm:"column:role;not null" json:"role"`
	Username string `gorm:"column:username;not null" form:"username" json:"username" binding:"required"`
	Password string `gorm:"column:password;not null" form:"password" json:"password" binding:"required"`
	Email    string `gorm:"column:email;not null" form:"email" json:"email"`
	Phone    string `gorm:"column:phone;not null" form:"phone" json:"phone"`
	Avatar   string `gorm:"column:avatar;not null" form:"avatar" json:"avatar"`
	// 外键。定义与Content模型的关联，表示一个用户可以有多个内容
	Contents     []Content          `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	ResumeData   []ResumeData       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	LoadedResume []LoadedResumeData `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

type Data struct {
	DataTemplate string `gorm:"column:data_template;not null" form:"data_template" json:"data_template"`
	DataStyle    string `gorm:"column:data_style;not null" form:"data_style" json:"data_style"`
	DataScript   string `gorm:"column:data_script;not null" form:"data_script" json:"data_script"`
}

type ResumeData struct {
	ResumeId     int       `gorm:"column:resume_id;primaryKey;autoIncrement" json:"resume_id"`
	Username     string    `gorm:"column:username;not null" json:"username"`
	UserID       int       `gorm:"column:user_id;not null" json:"user_id"`
	TemplateName string    `gorm:"column:template_name;not null" json:"template_name"`
	Resume       string    `gorm:"column:resume_data;not null" json:"resume_data"`
	Timestamp    time.Time `gorm:"column:timestamp;default:CURRENT_TIMESTAMP;type:datetime;not null"`
	IsShared     bool      `gorm:"column:is_shared;default:false;type:tinyint(1);not null"`
	ThumbnailUrl string    `gorm:"column:thumbnail_url;not null" json:"thumbnail_url"`
	ResumeName   string    `gorm:"column:resume_name;not null" json:"resume_name"`
	LikeCount    int       `gorm:"column:like_count;default:0;not null" json:"like_count"`
}

type LoadedResumeData struct {
	ResumeId   int       `gorm:"column:resume_id;primaryKey;autoIncrement" json:"resume_id"`
	ResumeName string    `gorm:"column:resume_name;not null" json:"resume_name"`
	Username   string    `gorm:"column:username;not null" json:"username"`
	UserID     int       `gorm:"column:user_id;not null"`
	URL        string    `gorm:"column:url;not null" json:"url"`
	Timestamp  time.Time `gorm:"column:timestamp;default:CURRENT_TIMESTAMP;type:datetime;not null"`
}

type ResumeLike struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ResumeID  int       `gorm:"column:resume_id;not null" json:"resume_id"`
	UserID    int       `gorm:"column:user_id;not null" json:"user_id"`
	Username  string    `gorm:"column:username;not null" json:"username"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;type:datetime;not null" json:"created_at"`
}

type ChangePasswordRequest struct {
	Password    string `json:"password" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

// TODO 留言功能
type Content struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"` // 标识每条留言，防止出现相同内容留言难以辨认的情况
	UserID    int       `gorm:"column:user_id;not null"`                      // 添加UserID字段作为外键,保证每个Content关联一个User
	User      User      `gorm:"-" json:"user"`                                // 不保存User到content表中
	Content   string    `gorm:"column:content;not null" form:"content" json:"content"`
	Timestamp time.Time `gorm:"column:timestamp;default:CURRENT_TIMESTAMP;type:datetime;not null"` // 打上时间戳
}

type MyClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
