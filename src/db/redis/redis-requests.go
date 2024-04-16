package redis_main

import (
	"encoding/json"
	"errors"
	"fmt"

	errors_module "github.com/pseudoelement/go-server/src/errors"
	auth_models "github.com/pseudoelement/go-server/src/modules/auth/models"
	"github.com/redis/go-redis/v9"
)

func Set[T comparable](key string, value T) error {
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}

	return nil
}

func Get(key string) (string, error) {
	value, err := client.Get(ctx, key).Result()

	if err == redis.Nil {
		return "", errors.New("value equals to null")
	} else if value == "" {
		return "", errors.New("value is empty")
	} else if err != nil {
		return "", err
	}

	return value, nil
}

func GetAll() ([]auth_models.UserRegister, errors_module.ErrorWithStatus) {
	users_chan := make(chan []auth_models.UserRegister, 1)
	var cursor uint64
	var keys []string
    for {
        keys, nextCursor, err := client.Scan(ctx, cursor, "*", 10).Result()
        if err != nil {
            panic(err)
        }

        fmt.Println("Keys:", keys)

        cursor = nextCursor

        if cursor == 0 {
            break
        }
    }
	return []auth_models.UserRegister{}, nil;
}

func SetStruct[T any](key string, object T) error {
	json, err := json.Marshal(object)
	if err != nil {
		return err;
	}

	err = client.Set(ctx, key, json, 0).Err()
	if err != nil {
		return err;
	}

	return nil
}

func GetStruct[T any](key string) (T, error) {
	json_string, err := client.Get(ctx, key).Result()
	res_generic := new(T)

	if err == redis.Nil {
		return *res_generic, errors.New("value equals to null")
	} else if json_string == "" {
		return *res_generic, errors.New("value is empty")
	} else if err != nil {
		return *res_generic, err
	}

	if err := json.Unmarshal([]byte(json_string), &res_generic); err != nil {
		return *res_generic, err
	}
	return *res_generic, nil
}