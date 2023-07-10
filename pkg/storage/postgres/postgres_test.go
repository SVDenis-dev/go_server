package postgres

import (
	"GoNews/pkg/storage"
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

func TestStorage_Posts(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	testDB, err := setupTestDB()
	if err != nil {
		fmt.Println(err)
	}
	emptyDB, err := setupEmptyDB()
	if err != nil {
		fmt.Println(err)
	}
	errorDB, err := setupErrorDB()
	if err != nil {
		fmt.Println(err)
	}
	tests := []struct {
		name    string
		fields  fields
		want    []storage.Post
		wantErr bool
	}{
		{
			name:   "Проверка получения всех записей",
			fields: fields{db: testDB},
			want: []storage.Post{
				{ID: 1, Title: "Заголовок 1", Content: "Содержание 1"},
				{ID: 2, Title: "Заголовок 2", Content: "Содержание 2"},
				{ID: 3, Title: "Заголовок 3", Content: "Содержание 3"},
			},
			wantErr: false,
		},
		{
			name:    "Проверка получения записей, когда база данных пуста",
			fields:  fields{db: emptyDB},
			want:    []storage.Post{},
			wantErr: false,
		},
		{
			name:    "Проверка получения записей, когда произошла ошибка",
			fields:  fields{db: errorDB},
			want:    []storage.Post{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.Posts()
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Posts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.Posts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func setupTestDB() (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, err
	}
	config.ConnConfig.Host = "localhost"
	config.ConnConfig.User = "testuser"
	config.ConnConfig.Password = "testpass"
	config.ConnConfig.Database = "testdb"

	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(context.Background(), `
		CREATE TABLE posts (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			content TEXT NOT NULL
		);
	`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(context.Background(), `
		INSERT INTO posts (title, content) VALUES
			('Заголовок 1', 'Содержание 1'),
			('Заголовок 2', 'Содержание 2'),
			('Заголовок 3', 'Содержание 3');
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func setupEmptyDB() (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, err
	}
	config.ConnConfig.Host = "localhost"
	config.ConnConfig.User = "testuser"
	config.ConnConfig.Password = "testpass"
	config.ConnConfig.Database = "emptydb"

	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func setupErrorDB() (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, err
	}
	config.ConnConfig.Host = "localhost"
	config.ConnConfig.User = "testuser"
	config.ConnConfig.Password = "testpass"
	config.ConnConfig.Database = "errordb"

	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(context.Background(), `
		CREATE TABLE posts (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			content TEXT NOT NULL
		);
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
