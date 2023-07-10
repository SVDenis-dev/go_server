package mongodb

import (
	"GoNews/pkg/storage"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestStorage_Posts(t *testing.T) {
	type fields struct {
		db *mongo.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    []storage.Post
		wantErr bool
	}{
		{
			name: "Тест с пустой базой данных",
			fields: fields{
				db: nil, // Установим базу данных в nil, чтобы имитировать пустую базу данных.
			},
			want:    []storage.Post{},
			wantErr: false,
		},
		{
			name: "Тест с базой данных, содержащей один пост",
			fields: fields{
				db: &mongo.Client{}, // Установим базу данных, содержащую один пост.
			},
			want: []storage.Post{
				{
					ID:         1,
					Title:      "Первый пост",
					Content:    "Это первый пост.",
					AuthorName: "Автор",
					CreatedAt:  0,
				},
			},
			wantErr: false,
		},
		{
			name: "Тест с базой данных, содержащей несколько постов",
			fields: fields{
				db: &mongo.Client{}, // Установим базу данных, содержащую несколько постов.
			},
			want: []storage.Post{
				{
					ID:         1,
					Title:      "Первый пост",
					Content:    "Это первый пост.",
					AuthorName: "Автор",
					CreatedAt:  0,
				},
				{
					ID:         2,
					Title:      "Второй пост",
					Content:    "Это второй пост.",
					AuthorName: "Автор",
					CreatedAt:  0,
				},
			},
			wantErr: false,
		},
		// Добавьте дополнительные тест-кейсы здесь
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
