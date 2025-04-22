package models

// User 结构体标签
// json:"user_id,uint" 表示在进行 JSON 序列化和反序列化时，会使用小写的 user_id 作为键名，并且会将其视为无符号整数。
// db:"user_id" 标签指定了在数据库操作中使用的字段名。
// User 用户类
type User struct {
	UserID       uint64 `json:"user_id,uint" db:"user_id"` // 指定json序列化/反序列化时使用小写user_id
	UserName     string `json:"username" db:"username" binding:"required"`
	Password     string `json:"password" db:"password" binding:"required"`
	AccessToken  string
	RefreshToken string
}
