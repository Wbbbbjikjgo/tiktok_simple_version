package domain

type User struct {
	// id、密码、随机盐字段在返回给用户时应屏蔽
	IsFollow      bool   `json:"is_follow,omitempty" gorm:"-"`
	Id            int64  `json:"id,omitempty" gorm:"primaryKey"`
	FollowCount   int64  `json:"follow_count,omitempty" gorm:"-"`
	FollowerCount int64  `json:"follower_count,omitempty" gorm:"-"`
	TotalLikes    int64  `json:"total_likes,omitempty"`
	LikesCount    int64  `json:"likes-count,omitempty"`
	Salt          string `json:"salt,omitempty" gorm:"type:char(4)"`
	Name          string `json:"name,omitempty" gorm:"type:varchar(32); index"`
	Pwd           string `json:"pwd,omitempty" gorm:"type:char(60)"`
}
