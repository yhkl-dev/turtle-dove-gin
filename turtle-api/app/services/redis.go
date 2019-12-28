package services

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"github.com/yhkl-dev/turtle-dove-beego/turtle-api/utils"
)

type redisService struct{}

func (rs *redisService) Get(keyName string) (string, error) {
	recv, err := redis.String(r.Do("Get", keyName))
	if err != nil {
		return "", err
	}
	return recv, nil
}

func (rs *redisService) IsExist(keyName string) bool {
	recv, _ := r.Do("EXISTS", keyName)
	if recv.(int64) == 1 {
		return true
	}
	return false
}

func (rs *redisService) Set(keyName, valueName string) error {

	expireTime := beego.AppConfig.String("EXPIRE_TIME")
	if len(expireTime) == 0 {
		expireTime = string(utils.DEFAULT_EXPIRE_SECONDS)
	}

	_, err := r.Do("SET", keyName, valueName, "EX", expireTime)
	if err != nil {
		return err
	}
	return nil
}
