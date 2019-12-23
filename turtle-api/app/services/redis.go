package services

type redisService struct{}

func (rs *redisService) Get(keyName string) (string, error) {
	recv, err := r.Do("Get", "name")
	if err != nil {
		return "", err
	}
	str := string(recv.([]byte))
	return str, nil
}

func (rs *redisService) Set(keyName, valueName string, expireTime int) error {
	return nil
}
