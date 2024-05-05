package domain

import (
	"context"
	"testing"
)

type SpyPostsStorage struct {
	Calls int
}

func (s *SpyPostsStorage) Store(ctx context.Context, post NewPost) (Post, error) {
	s.Calls++
	return Post{}, nil
}

func (s *SpyPostsStorage) List(ctx context.Context) ([]Post, error) {
	s.Calls++
	return []Post{}, nil
}

type SpyPostsPublisher struct {
	Calls int
}

func (s *SpyPostsPublisher) Publish(ctx context.Context, post Post) error {
	s.Calls++
	return nil
}

func TestPostsService(t *testing.T) {
	t.Run("HandleNewPost should store the post and publish it", func(t *testing.T) {
		publisher := &SpyPostsPublisher{}
		storage := &SpyPostsStorage{}
		service := NewPostsService(storage, publisher)
		service.HandleNewPost(context.Background(), NewPost{})

		if publisher.Calls != 1 {
			t.Errorf("expected publisher.Publish to be called 1 times, got: %v", publisher.Calls)
		}

		if storage.Calls != 1 {
			t.Errorf("expected storage.Store to be called 1 times, got: %v", storage.Calls)
		}
	})
}
