package types

import "time"

type RoleInfo struct {
	Id         int       `json:"Id"`
	RoleName   string    `json:"RoleName"`
	Password   string    `json:"Password"`
	RealName   string    `json:"RealName"`
	Gender     int       `json:"Gender"`
	Mobile     string    `json:"Mobile"`
	Email      string    `json:"Email"`
	Birthday   time.Time `json:"Birthday"`
	PositionId int       `json:"PositionId"`
	Intro      string    `json:"Intro "`
	Status     int       `json:"Status"`
	LoginNum   int       `json:"LoginNum"`
	LoginIp    string    `json:"LoginIp"`
	LoginTime  time.Time `json:"LoginTime"`
	CreateUser int       `json:"CreateUser"`
	CreateTime time.Time `json:"CreateTime"`
	UpdateUser int       `json:"UpdateUser"`
	UpdateTime time.Time `json:"UpdateTime"`
	Mark       int       `json:"Mark"`
}
