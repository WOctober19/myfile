package model

import "time"

const (
	UserAssociation         = "Users"
	UserAuthInfoAssociation = "AuthInfos"
	GroupAssociation        = "Groups"
)

func (*User) TableName() string {
	return "users"
}

type User struct {
	ID        uint       `json:"id" gorm:"autoIncrement;primaryKey"`
	Name      string     `json:"name" gorm:"size:100;not null;unique"`
	Password  string     `json:"-" gorm:"size:256"`
	Email     string     `json:"email" gorm:"size:256"`
	Avatar    string     `json:"avatar" gorm:"size:256"`
	AuthInfos []AuthInfo `json:"authInfos" gorm:"foreignKey:UserId;references:ID"`
	Groups    []Group    `json:"groups" gorm:"many2many:user_groups;"`
	Roles     []Role     `json:"roles" gorm:"many2many:user_roles;"`

	BaseModel
}

type AuthInfo struct {
	ID           uint      `json:"id" gorm:"autoIncrement;primaryKey"`
	UserId       uint      `json:"userId" gorm:"size:256"`
	Url          string    `json:"url" gorm:"size256"`
	AuthType     string    `json:"authType" gorm:"size:256"`
	AuthId       string    `json:"authId" gorm:"size:256"`
	AccessToken  string    `json:"-" gorm:"size:256"`
	RefreshToken string    `json:"-" gorm:"size:256"`
	Expiry       time.Time `json:"-"`

	BaseModel
}

type AuthUser struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	SetCookie bool   `json:"setCookie"`
	AuthType  string `json:"authType"`
	AuthCode  string `json:"authCode"`
}

type UserRole struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}
