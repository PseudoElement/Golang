package redis_main

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	errors_module "github.com/pseudoelement/go-server/src/errors"
	auth_models "github.com/pseudoelement/go-server/src/modules/auth/models"
	"github.com/pseudoelement/go-server/src/utils"
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
	start := time.Now()

	var cursor uint64
	var wg sync.WaitGroup;
	var emails []string;
    for {
        keys, nextCursor, err := client.Scan(ctx, cursor, "*", 10).Result()

		if(len(keys) > len(emails)){
			emails = keys;
		}

        if err != nil {
            panic(err)
        }
        cursor = nextCursor
        if cursor == 0 {
            break
        }
    }

	users := make([]auth_models.UserRegister, len(emails))
	
	for _, email := range emails {
		wg.Add(1);
		go func (){
			defer wg.Done();
			user, _ := GetStruct[auth_models.UserRegister](email);
			users = append(users, user)
		}()
	}

	wg.Wait()

	gotUsers := time.Since(start);
	fmt.Printf("gotUsers took %s\n", gotUsers);

	notEmptyUsers := utils.Filter(users, func(user auth_models.UserRegister, i int) bool {
		return user.Name != "" && user.Email != ""
	})

	return notEmptyUsers, nil;
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