package database

import (
	"context"
	"go-advanced-eventsarc-cqrs/models"

	_ "github.com/lib/pq"
)

func (repo *PostgresRepository) InsertFeed(ctx context.Context, feed *models.Feed) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO feeds (id, title, description) VALUES ($1, $2, $3)", feed.Id, feed.Title, feed.Description)
	return err
}

func (repo *PostgresRepository) ListFeeds(ctx context.Context) ([]*models.Feed, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, title, description, created_at FROM feeds ORDER BY created_at DESC")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	feeds := []*models.Feed{}

	for rows.Next() {
		feed := &models.Feed{}

		if err := rows.Scan(&feed.Id, &feed.Title, &feed.Description, &feed.CreatedAt); err != nil {
			return nil, err
		}

		feeds = append(feeds, feed)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return feeds, nil
}
