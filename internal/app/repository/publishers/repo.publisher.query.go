package publishers

const (
	columnSelectPublisher string = `
			publisher_id,
			name, 
			address, 
			phone,
			email, 
			website, 
			founded_year,
			country,
			contact_person_1,
			contact_person_2,
			fax,
			fb_link,
			twitter_link,
			web_link,
			join_date,
			entry_user,
			entry_name,
			entry_time
		`
	queryCreatePublisher string = `
		INSERT INTO publishers(
			name, 
			address, 
			phone,
			email, 
			website, 
			founded_year,
			country,
			contact_person_1,
			contact_person_2,
			fax,
			fb_link,
			twitter_link,
			web_link,
			join_date,
			entry_time
				)
		VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, now())
		`

	queryPublishersById string = `
		SELECT 
			` + columnSelectPublisher +
		`
		FROM publishers where publisher_id = $1`

	queryPublishersByName string = `SELECT count(*) FROM publishers where name = $1`

	queryDeletePublisher string = `DELETE FROM publishers where publisher_id = $1`

	queryUpdatePublisher string = `
		UPDATE publishers 
			SET
				name = $1, 
				address = $2, 
				phone = $3,
				email = $4, 
				website = $5, 
				founded_year = $6,
				country = $7,
				contact_person_1 = $8,
				contact_person_2 = $9,
				fax = $10,
				fb_link = $11,
				twitter_link = $12,
				web_link = $13,
				join_date = $14,
				entry_time = now()
		WHERE publisher_id = $15`
)
