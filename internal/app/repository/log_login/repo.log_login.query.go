package loglogin

const (
	queryCreateLoglogin = `
		INSERT INTO %s ("email", "full_name", "role", "ip_address", "login_time", "process_time", "expired_time") 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	queryCreateTable = `
		CREATE TABLE IF NOT EXISTS %s (
			id BIGSERIAL PRIMARY KEY,
			email VARCHAR(50) NOT NULL,
			full_name VARCHAR(100) NOT NULL,
			role VARCHAR(20) NOT NULL,
			ip_address VARCHAR(20) NOT NULL,
			login_time TIMESTAMP(3) DEFAULT NULL,
			logout_time TIMESTAMP(3) DEFAULT NULL,
			process_time TIMESTAMP(3) DEFAULT NULL,
			expired_time TIMESTAMP(3) DEFAULT NULL
		)`

	queryGetLoglogin = `
		SELECT 
			"id", 
			"email", 
			"full_name", 
			"role", 
			"ip_address", 
			"login_time", 
			"logout_time", 
			"process_time", 
			"expired_time"
		FROM %s
		WHERE email = $1 
		ORDER BY login_time DESC 
		LIMIT 1`

	queryUpdateLogLgoin = `
		UPDATE %s 
		SET 
			"full_name" = $1, 
			"role" = $2, 
			"ip_address" = $3, 
			"logout_time" = $4, 
			"process_time" = $5, 
			"expired_time" = $6 
		WHERE email = $7`
)
