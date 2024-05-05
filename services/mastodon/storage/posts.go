package storage

import (
	"context"

	"github.com/tomef96/coop/mastodon/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Post struct {
	Title string             `bson:"title"`
	Body  string             `bson:"body"`
	ID    primitive.ObjectID `bson:"_id,omitempty"`
}

type PostsStorage struct {
	client *mongo.Client
}

func NewPostsStorage() *PostsStorage {
	return &PostsStorage{
		client: NewMongoClient(),
	}
}

func (s *PostsStorage) Store(ctx context.Context, post domain.NewPost) (domain.Post, error) {
	res, err := s.client.Database("db").Collection("posts").InsertOne(ctx, Post{
		Title: post.Title,
		Body:  post.Body,
	})
	if err != nil {
		return domain.Post{}, err
	}

	return domain.Post{
		ID:    res.InsertedID.(primitive.ObjectID).Hex(),
		Title: post.Title,
		Body:  post.Body,
	}, nil
}

func (s *PostsStorage) List(ctx context.Context) ([]domain.Post, error) {
	cur, err := s.client.Database("db").Collection("posts").Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	storedPosts := []Post{}
	err = cur.All(ctx, &storedPosts)

	posts := make([]domain.Post, len(storedPosts))
	for i, post := range storedPosts {
		posts[i] = domain.Post{
			ID:    post.ID.Hex(),
			Title: post.Title,
			Body:  post.Body,
		}
	}

	return posts, err
}

func (s *PostsStorage) Close() error {
	return s.client.Disconnect(context.Background())
}
