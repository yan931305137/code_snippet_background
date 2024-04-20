package types

import "time"

type UserInfo struct {
	Id           int       `json:"id"`
	UserName     string    `json:"UserName"`
	Password     string    `json:"Password"`
	Gender       int       `json:"Gender"`
	Avatar       string    `json:"Avatar"`
	Mobile       string    `json:"Mobile"`
	Email        string    `json:"Email"`
	Birthday     time.Time `json:"Birthday"`
	ProvinceCode string    `json:"ProvinceCode"`
	CityCode     string    `json:"CityCode"`
	DistrictCode string    `json:"DistrictCode"`
	Address      string    `json:"Address"`
	CityName     string    `json:"CityName"`
	Status       int       `json:"Status"`
	LoginNum     int       `json:"LoginNum"`
	LoginIp      string    `json:"LoginIp"`
	LoginTime    time.Time `json:"LoginTime"`
	CreateTime   time.Time `json:"CreateTime"`
	Mark         int       `json:"Mark"`
}
