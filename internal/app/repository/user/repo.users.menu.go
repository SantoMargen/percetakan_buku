package user

const (
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
		id, 
		full_name,
		email, 
		password, 
		phone_number, 
		date_of_birth, 
		profile_picture, 
		gender, 
		address, 
		city, 
		country, 
		role, 
		created_by, 
		updated_by, 
		created_at, 
		updated_at
	FROM users 
	WHERE email = $1
	`
)
