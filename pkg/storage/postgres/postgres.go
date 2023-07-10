package postgres

import (
	"GoNews/pkg/storage"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Хранилище данных.
type Storage struct {
	db *pgxpool.Pool
}

// Конструктор, принимает строку подключения к БД.
func New(constr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

// получение всех публикаций
func (s *Storage) Posts() ([]storage.Post, error) {
	rows, err := s.db.Query(context.Background(),
		`SELECT 
				id,
				title, 
				content,
				author_id,
				created_at
		FROM posts
		ORDER BY id;
`)
	if err != nil {
		return nil, err
	}
	var posts []storage.Post
	for rows.Next() {
		var p storage.Post
		err = rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.AuthorID,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, rows.Err()
}

// AddPost - создание новой публикации
func (s *Storage) AddPost(p storage.Post) error {
	_, err := s.db.Exec(context.Background(), `INSERT INTO posts (id, autor_id, title, content, created_at) VALUES ($1, $2, $3, $4, $5)`,
		p.ID,
		p.AuthorID,
		p.Title,
		p.Content,
		p.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to add post: %w", err)
	}
	return nil
}

// обновление публикации
func (s *Storage) UpdatePost(post storage.Post) error {
	_, err := s.db.Exec(context.Background(), `UPDATE posts SET title=$1, content=$2 WHERE id=$3`, post.Title, post.Content, post.ID)
	if err != nil {
		return fmt.Errorf("failed to update post: %w", err)
	}
	return nil
}

// удаление публикации по ID
func (s *Storage) DeletePost(post storage.Post) error {
	_, err := s.db.Exec(context.Background(), `DELETE FROM posts WHERE id = $1`, post.ID)
	if err != nil {
		return fmt.Errorf("failed to delete post: %w", err)
	}
	return nil
}
