package service

import (
	"context"
	"log"
	"mgtest/internal/models"
	"reflect"
)

type ProfileService interface {
	InsertProfile(ctx context.Context, data models.Data) (models.Data, error)
	UpdateProfile(ctx context.Context, data models.Data, prof models.Data) (int, error)
	GetProfile(ctx context.Context, idString string) (models.Data, error)
	Close()
}
type CachingService interface {
	InsertProfile(ctx context.Context, data models.Data) error
	GetProfile(ctx context.Context, idString string) (models.Data, error)
}

type Service struct {
	PS ProfileService
	CS CachingService
}

func New(ps ProfileService, cs CachingService) *Service {
	return &Service{
		PS: ps,
		CS: cs,
	}
}

// InsertProfile создает новый профиль
func (s *Service) InsertProfile(ctx context.Context, data models.Data) (models.Data, error) {
	//Создаем профиль в БД
	new, err := s.PS.InsertProfile(ctx, data)
	//Если ошибка, то выходим и не сохраняем в кеш
	if err != nil {
		return data, err
	}
	//Создаем профиль в кеше
	if err := s.CS.InsertProfile(ctx, new); err != nil {
		log.Println("ошибка сохранения профиля в кеш:", err)
	}

	return new, nil
}

// UpdateProfile обновляет профиль в кеше и БД
func (s *Service) UpdateProfile(ctx context.Context, data models.Data) (int, error) {
	//Сначала проверяем кэш, чтобы лишний раз не обращаться к БД
	prof, err := s.CS.GetProfile(ctx, data.ID)
	if err != nil {
		//Если нет, то ищем профиль в БД
		prof, err = s.PS.GetProfile(ctx, data.ID)
		if err != nil {
			return 00, err
		}
	}
	//Делаем проверку на идентичность объектов. Если одинаковы, то выходим и ничего не делаем
	if reflect.DeepEqual(prof, data) {
		return 0, nil
	}
	//Иначе обновляем профиль
	count, err := s.PS.UpdateProfile(ctx, data, prof)
	if err != nil {
		return count, err
	}
	//И записываем его в кеш
	s.CS.InsertProfile(ctx, data)
	return count, err
}
