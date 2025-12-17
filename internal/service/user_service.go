package service

import (
	"context"
	"time"

	"age-calculator/internal/models"
	"age-calculator/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func formatDOB(t time.Time) string {
	return t.Format("2006-01-02")
}

func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}

func (s *UserService) Create(ctx context.Context, name string, dob time.Time) (models.UserBasicResponse, error) {
	id, err := s.repo.Create(ctx, name, dob)
	if err != nil {
		return models.UserBasicResponse{}, err
	}

	return models.UserBasicResponse{
		ID:   id,
		Name: name,
		DOB:  formatDOB(dob),
	}, nil
}


func (s *UserService) GetByID(ctx context.Context, id int64) (models.UserResponse, error) {
	u, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return models.UserResponse{}, err
	}
	return models.UserResponse{
		ID:   u.ID,
		Name: u.Name,
		DOB:  u.Dob.Format("2006-01-02"),

		Age:  calculateAge(u.Dob),
	}, nil
}

func (s *UserService) GetUsers(
	ctx context.Context,
	page, limit int,
) ([]models.UserResponse, error) {

	offset := (page - 1) * limit

	rows, err := s.repo.GetPaginated(
		ctx,
		int32(limit),
		int32(offset),
	)
	if err != nil {
		return nil, err
	}

	out := make([]models.UserResponse, 0, len(rows))
	for _, u := range rows {
		out = append(out, models.UserResponse{
			ID:   u.ID,
			Name: u.Name,
			DOB:  u.Dob.Format("2006-01-02"),
			Age:  calculateAge(u.Dob),
		})
	}
	return out, nil
}

func (s *UserService) Update(ctx context.Context, id int64, name string, dob time.Time) (models.UserBasicResponse, error) {
	if err := s.repo.Update(ctx, id, name, dob); err != nil {
		return models.UserBasicResponse{}, err
	}

	return models.UserBasicResponse{
		ID:   id,
		Name: name,
		DOB:  formatDOB(dob),
	}, nil
}


func (s *UserService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
