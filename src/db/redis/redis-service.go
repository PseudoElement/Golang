package redis_main

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	errors_module "github.com/pseudoelement/go-server/src/errors"
	auth_models "github.com/pseudoelement/go-server/src/modules/auth/models"
	slice_utils "github.com/pseudoelement/go-server/src/utils/slices"
	"github.com/redis/go-redis/v9"
)

func (r *RedisDB) Set(key string, value interface{}) error {
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = r.client.Set(r.ctx, key, json, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisDB) Get(key string) (string, error) {
	value, err := r.client.Get(r.ctx, key).Result()

	if err == redis.Nil {
		return "", errors.New("value equals to null")
	} else if value == "" {
		return "", errors.New("value is empty")
	} else if err != nil {
		return "", err
	}

	return value, nil
}

func (r *RedisDB) GetAllUsers() ([]auth_models.UserToClient, errors_module.ErrorWithStatus) {
	start := time.Now()

	var cursor uint64
	var wg sync.WaitGroup
	var emails []string
	for {
		newEmails, nextCursor, err := r.client.Scan(r.ctx, cursor, "*", 10).Result()

		emails = r.getUpdatedEmailsList(newEmails, emails)

		if err != nil {
			panic(err)
		}
		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}

	users := make([]auth_models.UserToClient, len(emails))

	for _, email := range emails {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var user auth_models.UserToClient
			err := r.GetStruct(email, &user)
			if err == nil {
				users = append(users, user)
			}
		}()
	}

	wg.Wait()

	gotUsers := time.Since(start)
	fmt.Printf("gotUsers took %s\n", gotUsers)

	notEmptyUsers := slice_utils.Filter(users, func(user auth_models.UserToClient, i int) bool {
		return user.Name != "" && user.Email != ""
	})

	return notEmptyUsers, nil
}

func (r *RedisDB) SetStruct(key string, object interface{}) error {
	jsonBytes, err := json.Marshal(object)
	if err != nil {
		return err
	}

	err = r.client.Set(r.ctx, key, jsonBytes, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisDB) GetStruct(key string, structToAppendValue interface{}) error {
	json_string, err := r.client.Get(r.ctx, key).Result()

	if err == redis.Nil {
		return errors.New("value equals to null")
	} else if json_string == "" {
		return errors.New("value is empty")
	} else if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(json_string), &structToAppendValue); err != nil {
		return err
	}
	return nil
}

func (r *RedisDB) getUpdatedEmailsList(newEmails []string, oldEmails []string) []string {
	var updatedEmails []string
	for _, newEmail := range newEmails {
		if !slice_utils.Contains(oldEmails, newEmail) {
			updatedEmails = append(updatedEmails, newEmail)
		}
	}
	updatedEmails = append(updatedEmails, oldEmails...)

	return updatedEmails
}
