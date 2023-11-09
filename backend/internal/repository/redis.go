package repository

import (
	"context"
	"encoding/json"
	"mgtest/internal/models"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	rc *redis.Client
}

const (
	sessionTime = time.Hour * 24 * 2
)

func NewRedis(r *redis.Client) *Redis {
	return &Redis{
		rc: r,
	}
}

func (r *Redis) InsertProfile(ctx context.Context, data models.Data) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = r.rc.Set(ctx, data.ID, bytes, sessionTime).Err()
	if err != nil {
		return err
	}
	return nil

}

func (r *Redis) GetProfile(ctx context.Context, idString string) (models.Data, error) {
	var data models.Data
	result, err := r.rc.Get(ctx, idString).Result()
	if err == redis.Nil {
		return data, err
	}
	r.rc.ExpireGT(ctx, idString, sessionTime)
	if err := json.Unmarshal([]byte(result), &data); err != nil {
		return data, err
	}
	return data, nil
}
