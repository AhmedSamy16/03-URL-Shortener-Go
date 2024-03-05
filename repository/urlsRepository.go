package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/AhmedSamy16/03-url-shortener-Go/types"
	"github.com/google/uuid"
)

type UrlRepository struct {
	DB *sql.DB
}

func (repo *UrlRepository) GetAllUrls(ctx context.Context) (*[]types.ShortUrlModel, error) {
	stmt, err := repo.DB.PrepareContext(ctx, `SELECT * FROM urls`)
	if err != nil {
		return nil, errors.New("failed to get urls")
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, errors.New("failed to query users")
	}
	defer rows.Close()

	urls := []types.ShortUrlModel{}

	for rows.Next() {
		u, err := scanRowToShortURL(rows)
		if err != nil {
			return nil, err
		}
		urls = append(urls, u)
	}

	return &urls, nil
}

func (repo *UrlRepository) GetShortUrl(ctx context.Context, shortUrl uuid.UUID) (*types.ShortUrlModel, error) {
	stmt, err := repo.DB.PrepareContext(ctx, `UPDATE urls SET clicks = clicks + 1 WHERE short_url = $1 RETURNING *`)
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to get url")
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, shortUrl)
	if err != nil {
		return nil, errors.New("url not found")
	}
	defer rows.Close()

	u, err := scanRowToShortURL(rows)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (repo *UrlRepository) CreateShortUrl(ctx context.Context, urlToCreate *types.CreateShortUrl) (*uuid.UUID, error) {
	stmt, err := repo.DB.PrepareContext(ctx, `INSERT INTO urls (url) VALUES ($1) RETURNING short_url`)
	if err != nil {
		return nil, errors.New("failed to start creating")
	}
	defer stmt.Close()

	var url uuid.UUID
	err = stmt.QueryRowContext(ctx, urlToCreate.URL).Scan(&url)
	if err != nil {
		return nil, errors.New("failed to insert url")
	}

	return &url, nil
}

func scanRowToShortURL(row *sql.Rows) (types.ShortUrlModel, error) {
	var id uuid.UUID
	var url string
	var shortUrl uuid.UUID
	var clicks int

	if row.Next() {
		err := row.Scan(&id, &url, &shortUrl, &clicks)
		if err != nil {
			return types.ShortUrlModel{}, err
		}
	} else {
		return types.ShortUrlModel{}, errors.New("nothing exists")
	}

	return types.ShortUrlModel{
		ID:       id,
		URL:      url,
		ShortUrl: shortUrl,
		Clicks:   clicks,
	}, nil
}
