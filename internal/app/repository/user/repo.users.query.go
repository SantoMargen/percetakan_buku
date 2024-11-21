package user

const (
	columnSelectUser string = `
		"id", 
		"full_name", 
		"email", 
		"password", 
		"role", 
		COALESCE("gender", '') AS "gender", 
		COALESCE("phone_number", '') AS "phone_number", 
		"date_of_birth",
		COALESCE("profile_picture", '') AS "profile_picture", 
		COALESCE("address", '') AS "address", 
		COALESCE("city", '') AS "city", 
		COALESCE("country", '') AS "country", 
		"created_by", 
		COALESCE("updated_by", 0) AS "updated_by", 
		"created_at",
		"updated_at",
		"status",
		"born_place",
		"biography",
		"experience",
		"achievement",
		"last_education"
	`

	queryCreateUser string = `
			INSERT INTO users(
				full_name, 
				email, 
				password, 
				role, 
				gender,
				created_at
				 )
			VALUES
			($1, $2, $3, $4, $5, now())
			`

	queryGetUserByEmail string = `
		SELECT 
			` + columnSelectUser + `
		FROM users 
		WHERE email = $1
	`
	queryGetUserByUserId string = `
		SELECT 
			` + columnSelectUser + `
		FROM users 
		WHERE id = $1
		`

	queryUpdateRole string = `
	UPDATE users
	SET 
		"role" = $1,
		updated_by = $2
	WHERE id = $3`

	queryUpdatePassword string = `
	UPDATE users
	SET 
		password = $1,
		updated_by = $2
	WHERE id = $3`

	queryCountTotalUser string = `
		SELECT COUNT(*) FROM users
	`
	queryCountUserActive string = `
		SELECT COUNT(*)
			FROM log_login_$1
		WHERE logout_time IS NOT NULL
		AND process_time >= (CURRENT_DATE - EXTRACT(DOW FROM CURRENT_DATE)::INT - 7)
		AND process_time < (CURRENT_DATE - EXTRACT(DOW FROM CURRENT_DATE)::INT);
	
	`

	queryCountNewUserLastWeek string = `
	SELECT COUNT(*)
	FROM users
	WHERE created_at >= (CURRENT_DATE - INTERVAL '1 day' * EXTRACT(DOW FROM CURRENT_DATE)) - INTERVAL '7 days'
	  AND created_at < (CURRENT_DATE - INTERVAL '1 day' * EXTRACT(DOW FROM CURRENT_DATE));
	
	`

	queryUpdateUser string = `
		UPDATE users
		SET 
			full_name = $1,
			date_of_birth = $2,
			born_place = $3,
			last_education = $4,
			gender = $5,
			biography = $6,
			experience = $7,
			achievement = $8,
			address = $9,
			phone_number = $10,
			country = $11,
			city = $12

		WHERE id = $13`
)
