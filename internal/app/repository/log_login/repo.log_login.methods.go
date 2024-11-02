package loglogin

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	loglogin "siap_app/internal/app/entity/log_login"
	"siap_app/internal/app/helpers"
	"time"
)

func (r *repository) CreateLogLogin(ctx context.Context, logLogin loglogin.LogloginRequest) error {
	tableName, err := helpers.GetTableNameMonth("log_login", nil)
	if err != nil {
		return fmt.Errorf("failed to generate tableName: %w", err)
	}

	now := time.Now()
	logLogin.LoginTime = now
	logLogin.ProcessTime = &now
	expiration := now.Add(2 * time.Hour)
	logLogin.ExpiredTime = &expiration

	query := fmt.Sprintf(queryCreateLoglogin, tableName)
	_, err = r.db.ExecContext(ctx, query,
		logLogin.Email,
		logLogin.FullName,
		logLogin.Role,
		logLogin.IPAddress,
		logLogin.LoginTime,
		logLogin.ProcessTime,
		logLogin.ExpiredTime,
	)
	if err != nil {
		return fmt.Errorf("failed to create log login: %w", err)
	}

	return nil
}

func (r *repository) CheckTableAlreadyExist(ctx context.Context) bool {
	tableName, err := helpers.GetTableNameMonth("log_login", nil)
	if err != nil {
		return false
	}

	query := fmt.Sprintf("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_schema = 'public' AND table_name = '%s')", tableName)

	var exists bool
	err = r.db.QueryRowContext(ctx, query).Scan(&exists)
	if err != nil {
		log.Printf("Error checking if table exists: %v", err)
		return false
	}

	return exists
}

func (r *repository) CreateTableLogLogin(ctx context.Context) error {
	tableName, err := helpers.GetTableNameMonth("log_login", nil)
	if err != nil {
		return fmt.Errorf("failed to generate tableName: %w", err)
	}

	qry := fmt.Sprintf(queryCreateTable, tableName)

	_, err = r.db.ExecContext(ctx, qry)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	indexQueries := []string{
		fmt.Sprintf("CREATE INDEX IF NOT EXISTS idx_%s_email ON %s (email)", tableName, tableName),
		fmt.Sprintf("CREATE INDEX IF NOT EXISTS idx_%s_ipaddress ON %s (ip_address)", tableName, tableName),
		fmt.Sprintf("CREATE INDEX IF NOT EXISTS idx_%s_role ON %s (role)", tableName, tableName),
		fmt.Sprintf("CREATE INDEX IF NOT EXISTS idx_%s_full_name ON %s (full_name)", tableName, tableName),
	}

	for _, query := range indexQueries {
		_, err := r.db.ExecContext(ctx, query)
		if err != nil {
			return fmt.Errorf("failed to create index: %w", err)
		}
	}

	return nil
}

func (r *repository) GetLastLogLoginByEmail(ctx context.Context, email string) (*loglogin.LogloginResponse, error) {
	var logLogin loglogin.LogloginResponse

	tableName, err := helpers.GetTableNameMonth("log_login", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tableName: %w", err)
	}

	query := fmt.Sprintf(queryGetLoglogin, tableName)
	err = r.db.GetContext(ctx, &logLogin, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with email %d not found", email)
		}
		return nil, err
	}

	return &logLogin, nil
}

func (r *repository) UpdateLogLogin(ctx context.Context, logLogin loglogin.LogloginRequest) error {
	tableName, err := helpers.GetTableNameMonth("log_login", nil)
	if err != nil {
		return fmt.Errorf("failed to generate tableName: %w", err)
	}

	query := fmt.Sprintf(queryUpdateLogLgoin, tableName)
	_, err = r.db.ExecContext(ctx, query,
		logLogin.FullName,
		logLogin.Role,
		logLogin.IPAddress,
		logLogin.LogoutTime,
		logLogin.ProcessTime,
		logLogin.ExpiredTime,
		logLogin.Email,
	)

	if err != nil {
		return fmt.Errorf("failed to update log login: %w", err)
	}

	return nil
}
