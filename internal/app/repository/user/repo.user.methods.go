package user

import (
	"context"
	"database/sql"
	"fmt"
	"siap_app/internal/app/entity/user"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

func (r *repository) GetListUserAll(ctx context.Context, input user.PaginationUser) ([]user.ResponseUser, int64, error) {

	var (
		dataUserList    []user.ResponseUser
		offset          int
		query           string
		countQuery      string
		total           int64
		totalUser       int
		totalActiveUser int
		totalNewUser    int
	)

	errTotalUser := r.db.QueryRowContext(ctx, queryCountTotalUser).Scan(&totalUser)
	if errTotalUser != nil {
		if errTotalUser == sql.ErrNoRows {
			totalUser = 0
		}
	}

	errActiveUser := r.db.QueryRowContext(ctx, queryCountUserActive).Scan(&totalActiveUser)
	if errActiveUser != nil {
		if errActiveUser == sql.ErrNoRows {
			totalActiveUser = 0
		}
	}

	errNewUser := r.db.QueryRowContext(ctx, queryCountNewUserLastWeek).Scan(&totalNewUser)
	if errNewUser != nil {
		if errNewUser == sql.ErrNoRows {
			totalNewUser = 0
		}
	}

	offset = (input.Page - 1) * input.Size
	query = "SELECT " + columnSelectUser + " FROM users WHERE 1=1"
	countQuery = "SELECT COUNT(*) FROM users WHERE 1=1"

	var args []interface{}
	var nextLimit int

	if input.Filter != nil {
		if input.Filter.Email != "" {
			query += " AND email = $" + strconv.Itoa(len(args)+1)
			countQuery += " AND email = $" + strconv.Itoa(len(args)+1)
			args = append(args, input.Filter.Email)
			nextLimit++
		}
		if input.Filter.Role != "" {
			query += " AND role LIKE $" + strconv.Itoa(len(args)+1)
			countQuery += " AND role LIKE $" + strconv.Itoa(len(args)+1)
			args = append(args, "%"+input.Filter.Role+"%")
			nextLimit++
		}
	}

	query += " ORDER BY id ASC LIMIT $" + strconv.Itoa(len(args)+1) + " OFFSET $" + strconv.Itoa(len(args)+2)
	args = append(args, input.Size, offset)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var dataUser user.ResponseUser
		dataUser.TotalUsers = totalUser
		dataUser.ActiveUsers = totalActiveUser
		dataUser.NewUsersLastWeek = totalNewUser

		if err := rows.Scan(
			&dataUser.ID,
			&dataUser.FullName,
			&dataUser.Email,
			&dataUser.Password,
			&dataUser.Role,
			&dataUser.Gender,
			&dataUser.PhoneNumber,
			&dataUser.TanngalLahir,
			&dataUser.ImageProfile,
			&dataUser.Address,
			&dataUser.City,
			&dataUser.Country,
			&dataUser.CreatedBy,
			&dataUser.UpdatedBy,
			&dataUser.CreatedAt,
			&dataUser.UpdatedAt,
			&dataUser.Status,
			&dataUser.Biography,
			&dataUser.BornPlace,
			&dataUser.Experience,
			&dataUser.Achievement,
			&dataUser.LastEduaction,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row: %w", err)
		}
		dataUserList = append(dataUserList, dataUser)
	}

	err = r.db.QueryRowContext(ctx, countQuery, args[:len(args)-2]...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count items: %w", err)
	}

	return dataUserList, total, nil
}

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
		&user.Role,
		&user.Gender,
		&user.PhoneNumber,
		&user.TanngalLahir,
		&user.ImageProfile,
		&user.Address,
		&user.City,
		&user.Country,
		&user.CreatedBy,
		&user.UpdatedBy,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Status,
		&user.BornPlace,
		&user.Biography,
		&user.Experience,
		&user.Achievement,
		&user.LastEduaction,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.Wrap(err, "user not found")
		}
		return user, errors.Wrap(err, "failed to get user by email")
	}

	return user, nil
}

func (r *repository) GetUserById(ctx context.Context, userId string) (user.ResponseUser, error) {
	var user user.ResponseUser
	getUserID, _ := strconv.Atoi(userId)
	err := r.db.QueryRowContext(ctx, queryGetUserByUserId, getUserID).Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.Gender,
		&user.PhoneNumber,
		&user.TanngalLahir,
		&user.ImageProfile,
		&user.Address,
		&user.City,
		&user.Country,
		&user.CreatedBy,
		&user.UpdatedBy,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Status,
		&user.BornPlace,
		&user.Biography,
		&user.Experience,
		&user.Achievement,
		&user.LastEduaction,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.Wrap(err, "user not found")
		}
		return user, errors.Wrap(err, "failed to get user by id")
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

func (r *repository) UpdateUser(ctx context.Context, userId int, input user.RequestUpdateUser) error {

	result, err := r.db.ExecContext(ctx, queryUpdateUser,
		input.Fullname,
		input.SetBirthDate,
		input.BornPlace,
		input.Graduated,
		input.Gender,
		input.Biografi,
		input.Experience,
		input.Achievement,
		input.Address,
		input.PhoneNumber,
		input.Country,
		input.City,
		userId,
	)

	if err != nil {
		return fmt.Errorf("error execute: %w", err)
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

func (r *repository) GetLogLogin(ctx context.Context, input user.PaginationLog) ([]user.ResponseLog, int64, error) {

	var (
		dataLogLoginUser []user.ResponseLog
		offset           int
		query            string
		countQuery       string
		total            int64
		tableName        string
	)

	getYear := strconv.Itoa(time.Now().Year())
	tableName = "log_login_" + getYear

	offset = (input.Page - 1) * input.Size
	query = "SELECT id,email,ip_address,user_agent,login_time,logout_time FROM " + tableName + " WHERE 1=1"

	countQuery = "SELECT COUNT(*) FROM " + tableName + "  WHERE 1=1"

	var args []interface{}
	var nextLimit int

	if input.Filter != nil {
		if input.Filter.LastActivity != "" {
			query += " AND login_time = $" + strconv.Itoa(len(args)+1)
			countQuery += " AND login_time = $" + strconv.Itoa(len(args)+1)
			args = append(args, input.Filter.LastActivity)
			nextLimit++
		}
		if input.Filter.Email != "" {
			query += " AND email LIKE $" + strconv.Itoa(len(args)+1)
			countQuery += " AND email LIKE $" + strconv.Itoa(len(args)+1)
			args = append(args, "%"+input.Filter.Email+"%")
			nextLimit++
		}
	}

	query += " ORDER BY id ASC LIMIT $" + strconv.Itoa(len(args)+1) + " OFFSET $" + strconv.Itoa(len(args)+2)
	args = append(args, input.Size, offset)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var dataLog user.ResponseLog

		if err := rows.Scan(
			&dataLog.ID,
			&dataLog.Email,
			&dataLog.IPAddress,
			&dataLog.UserAgent,
			&dataLog.LoginTime,
			&dataLog.LogoutTime,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row: %w", err)
		}
		dataLogLoginUser = append(dataLogLoginUser, dataLog)
	}

	err = r.db.QueryRowContext(ctx, countQuery, args[:len(args)-2]...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count items: %w", err)
	}

	return dataLogLoginUser, total, nil
}
