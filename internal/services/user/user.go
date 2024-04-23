package service

import (
	"code_snippet/internal/model"
	"code_snippet/internal/utils"
	"errors"
	"time"
)

// UserService 表示用户管理服务。
type UserService struct{}

// NewUserService 创建一个新的 UserService 实例。
func NewUserService() *UserService {
	return &UserService{}
}

// PostLogin 用户登录功能。
func (s *UserService) PostLogin(userName, password string) (int, string, error) {
	// 查询用户
	var userInfo model.User
	isExist, err := utils.XormDb.Table("user").Where("user_name = ? AND mark = 1", userName).Get(&userInfo)
	if err != nil {
		return 0, "", err
	}
	if !isExist {
		return 0, "", errors.New("用户名或密码不正确")
	}

	// 密码校验
	pwdHash, err := utils.Md5(password + userName)
	if err != nil {
		return 0, "", err
	}
	if userInfo.Password != pwdHash {
		return 0, "", errors.New("密码不正确")
	}

	// 检查账号状态
	if userInfo.Status != 1 {
		return 0, "", errors.New("您的账号已被禁用，请联系管理员")
	}

	// 更新登录时间
	now := time.Now()
	updateData := model.User{
		LoginTime: now,
	}
	_, err = utils.XormDb.ID(userInfo.Id).Update(&updateData)
	if err != nil {
		return 0, "", err
	}

	// 生成Token
	token, err := utils.GenToken(userInfo.UserName)
	if err != nil {
		return 0, "", err
	}

	return userInfo.Id, token, nil
}

// PostRegister 用户注册功能。
func (s *UserService) PostRegister(userName, password string) error {
	// 检查用户名是否已存在
	var existingUser model.User
	exist, err := utils.XormDb.Table("user").Where("user_name = ?", userName).Get(&existingUser)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("用户名已被注册")
	}

	// 对密码进行哈希处理
	pwdHash, err := utils.Md5(password + userName)
	if err != nil {
		return err
	}

	// 创建新用户
	newUser := model.User{
		UserName: userName,
		Password: pwdHash,
		Status:   1, // 设置为活跃状态
		Mark:     1,
		// 其他字段根据需要设置
	}

	// 插入数据库
	_, err = utils.XormDb.Insert(&newUser)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetInformation(username string) (interface{}, error) {
	var inform model.User
	_, err := utils.XormDb.Table("user").Where("user_name = ? ", username).Get(&inform)
	if err != nil {
		return nil, err
	}
	return inform, nil
}
