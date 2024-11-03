package user

import (
	"context"
	"database/sql"
	"fmt"
	"siap_app/internal/app/entity/user"

	"github.com/pkg/errors"
)

func (r *repository) CreateUser(ctx context.Context, input user.RegisterRequest) error {
	_, err := r.db.ExecContext(ctx, queryCreateUser,
		input.FullName,
		input.Email,
		input.Password,
		input.Role,
		input.Gender)

	if err != nil {
		return errors.Wrap(err, "failed to create user")
	}

	return nil
}

func (r *repository) CreateUserByAdmin(ctx context.Context, input user.RegisterByAdminRequest) error {
	_, err := r.db.ExecContext(ctx, queryCreateUser,
		input.FullName,
		input.Email,
		input.Password,
		input.Role,
		input.Gender)

	if err != nil {
		return errors.Wrap(err, "failed to create user")
	}

	return nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (user.ResponseUser, error) {
	var user user.ResponseUser
	err := r.db.QueryRowContext(ctx, queryGetUserByEmail, email).Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.Password,
		&user.PhoneNumber,
		&user.TanngalLahir,
		&user.ImageProfile,
		&user.Gender,
		&user.Address,
		&user.City,
		&user.Country,
		&user.Role,
		&user.CreatedBy,
		&user.UpdatedBy,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.Wrap(err, "user not found")
		}
		return user, errors.Wrap(err, "failed to get user by email")
	}

	return user, nil
}

func (r *repository) UpdateRoleUser(ctx context.Context, id, userId int, role string) error {
	result, err := r.db.ExecContext(ctx, queryUpdateRole, role, userId, id)
	if err != nil {
		return fmt.Errorf("failed to update role: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to fetch affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id %d", id)
	}

	return nil
}

func (r *repository) UpdatePasswordUser(ctx context.Context, userId int, password string) error {
	result, err := r.db.ExecContext(ctx, queryUpdatePassword, password, userId, userId)
	if err != nil {
		return fmt.Errorf("failed to update role: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to fetch affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id %d", userId)
	}

	return nil
}
