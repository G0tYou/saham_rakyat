package redis

import (
	"os"
	config "saham_rakyat/config"
	"strings"
	"time"
)

func Set(key string, data []byte) error {
	client, err := config.ConnectRedis()
	if err != nil {
		return err
	}
	exp, err := time.ParseDuration(os.Getenv("REDIS_EXP"))
	if err != nil {
		return err
	}
	err = client.Set(key, data, exp).Err()
	if err != nil {
		return err
	}

	return nil
}

func Get(key string) (string, error) {
	var value string
	client, err := config.ConnectRedis()
	if err != nil {
		return value, err
	}

	value, err = client.Get(key).Result()
	if err != nil {
		return value, err
	}

	return value, nil
}

func Delete(key string) error {
	client, err := config.ConnectRedis()
	if err != nil {
		return err
	}

	err = client.Del(key).Err()
	if err != nil {
		return err
	}

	return nil
}

func MultipleDelete(key string) error {
	client, err := config.ConnectRedis()
	if err != nil {
		return err
	}

	iter := client.Scan(0, "*"+key+"*", 0).Iterator()

	// Iterate over the keys and delete them
	for iter.Next() {
		val := iter.Val()
		if strings.Contains(val, key) {
			err := client.Del(val).Err()
			if err != nil {
				return err
			}
		}
	}

	return nil
}
