package repository

import (
	"context"
	"database/sql"
	"time"

	"age-calculator/db/sqlc"
)

type UserRepository struct {
	q *sqlc.Queries
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		q: sqlc.New(db),
	}
}

func (r *UserRepository) Create(ctx context.Context, name string, dob time.Time) (int64, error) {
	res, err := r.q.CreateUser(ctx, sqlc.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (sqlc.GetUserByIDRow, error) {
	return r.q.GetUserByID(ctx, id)
}

func (r *UserRepository) GetPaginated(
	ctx context.Context,
	limit, offset int32,
) ([]sqlc.GetUsersPaginatedRow, error) {
	return r.q.GetUsersPaginated(ctx, sqlc.GetUsersPaginatedParams{
		Limit:  limit,
		Offset: offset,
	})
}


func (r *UserRepository) Update(ctx context.Context, id int64, name string, dob time.Time) error {
	_, err := r.q.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob,
	})
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	return r.q.DeleteUser(ctx, id)
}
