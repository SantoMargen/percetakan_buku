package leveluser

import (
	"context"
	"database/sql"
	"fmt"
	levelusers "siap_app/internal/app/entity/level_users"
)

func (r *repository) GetLevelUsers(ctx context.Context) ([]levelusers.LevelUser, error) {
	var levelUsers []levelusers.LevelUser

	rows, err := r.db.QueryContext(ctx, qryLevelUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var l levelusers.LevelUser
		err := rows.Scan(
			&l.ID,
			&l.LevelUser,
			&l.Keterangan,
			&l.CreatedAt,
			&l.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		levelUsers = append(levelUsers, l)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return levelUsers, nil
}

func (r *repository) GetLevelUserByID(ctx context.Context, id int) (levelusers.LevelUser, error) {
	var levelUser levelusers.LevelUser

	row := r.db.QueryRowContext(ctx, qryLevelUserByID, id)

	err := row.Scan(
		&levelUser.ID,
		&levelUser.LevelUser,
		&levelUser.Keterangan,
		&levelUser.CreatedAt,
		&levelUser.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return levelUser, fmt.Errorf("no level user found with id %d", id)
		}
		return levelUser, err
	}

	return levelUser, nil
}
