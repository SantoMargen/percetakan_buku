package notification

const (
	columnSelectNotification string = `
			id_log_notif,
			key_notif,
			desc_notif,
			title_notif,
			sender, 
			receiver,
			flag_read,
			url_redirect,
			entry_time
	`

	qryInsertNotification = `
		INSERT INTO log_notif (
			key_notif,
			desc_notif,
			title_notif,
			sender, 
			receiver,
			flag_read,
			url_redirect,
			entry_time
	) 
	VALUES 
		($1, $2, $3, $4, $5, 0, $6, now())`

	queryDeleteNotification string = `
		DELETE FROM log_notif WHERE id_log_notif = $1 and receiver = $2`

	queryNotificationById string = `
		SELECT 
			` + columnSelectNotification + `
		FROM log_notif where id_log_notif = $1 and receiver = $2 ORDER BY id_log_notif ASC`
)
