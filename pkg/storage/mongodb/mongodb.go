package mongodb

import (
	"GoNews/pkg/storage"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseName   = "posts" // имя БД
	collectionName = "posts" // имя коллекции в БД
)

// Хранилище данных.
type Storage struct {
	db *mongo.Client
}

// Конструктор, принимает строку подключения к БД.
func New(constr string) (*Storage, error) {
	mongoOpts := options.Client().ApplyURI(constr)
	db, err := mongo.Connect(context.Background(), mongoOpts)
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
	collection := s.db.Database(databaseName).Collection(collectionName)
	filter := bson.D{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	var data []storage.Post
	for cur.Next(context.Background()) {
		var p storage.Post
		err := cur.Decode(&p)
		if err != nil {
			return nil, err
		}
		data = append(data, p)
	}
	return data, cur.Err()
}

// создание новой публикации
func (s *Storage) AddPost(p storage.Post) error {
	collection := s.db.Database(databaseName).Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), p)
	if err != nil {
		return err
	}
	return nil
}

// обновление публикации
func (s *Storage) UpdatePost(p storage.Post) error {
	collection := s.db.Database(databaseName).Collection(collectionName)
	selector := bson.M{"id": p.ID}
	// Обновить поля документа, которые необходимо изменить
	/* update := bson.M{"$set": bson.M{
		"title":      p.Title,
		"content":    p.Content,
		"author_id":  p.AuthorID,
		"created_at": p.CreatedAt,
	}} */

	// Сохранить изменения в базе данных
	_, err := collection.UpdateByID(context.Background(), p.ID, selector)
	if err != nil {
		return err
	}
	return nil
}

// удаление публикации по ID
func (s *Storage) DeletePost(p storage.Post) error {
	collection := s.db.Database(databaseName).Collection(collectionName)
	selector := bson.M{"id": p.ID}
	_, err := collection.DeleteOne(context.Background(), selector)
	if err != nil {
		return err
	}
	return nil
}
