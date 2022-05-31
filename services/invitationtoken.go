package services

import (
	"context"
	"fmt"
	"time"

	"github.com/arnab333/golang-app-invite-service/helpers"
)

func CreateInviteToken(ctx context.Context, authorID string) (string, error) {
	it := "INVITE-" + helpers.GetRandomString(16)

	now := time.Now()

	expires := time.Now().Add(time.Hour * 24 * 7)

	err := redisConn.redisClient.Set(ctx, it, authorID, expires.Sub(now)).Err()
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return it, nil
}

func DeleteInviteToken(ctx context.Context, key string) (int64, error) {
	deleted, err := redisConn.redisClient.Del(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
