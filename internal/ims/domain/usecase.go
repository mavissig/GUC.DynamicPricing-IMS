package domain

import "github.com/google/uuid"

type Repo interface {
	Set(key string, data []byte) error
	GetByKey(key string) ([]byte, error)
}

type UseCase struct {
	repo Repo
}

func New(
	repo Repo,
) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (uc *UseCase) DataSet(data *Data) error {
	err := uc.repo.Set(data.ID.String(), data.Data)
	return err
}

func (uc *UseCase) DataGetByKey(key uuid.UUID) (*Data, error) {
	res, err := uc.repo.GetByKey(key.String())
	if err != nil {
		return nil, err
	}

	data := &Data{
		ID:   key,
		Data: res,
	}

	return data, nil
}
