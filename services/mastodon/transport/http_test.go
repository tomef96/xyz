package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tomef96/coop/mastodon/domain"
)

type PostsStorage struct{}

func (p *PostsStorage) Store(ctx context.Context, post domain.NewPost) (domain.Post, error) {
	return domain.Post{}, nil
}

func (p *PostsStorage) List(ctx context.Context) ([]domain.Post, error) {
	return []domain.Post{
		{ID: "test", Title: "title", Body: "body"},
	}, nil
}

type PostsPublisher struct{}

func (p *PostsPublisher) Publish(ctx context.Context, post domain.Post) error {
	return nil
}

func TestHTTPServer_Serve(t *testing.T) {
	t.Run("it responds to get calls towards /posts", func(t *testing.T) {
		postsService := domain.NewPostsService(&PostsStorage{}, &PostsPublisher{})

		router := Router(postsService)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/posts", nil)

		router.ServeHTTP(w, req)

		if w.Code != 200 {
			t.Errorf("want status code %d, got %d", 200, w.Code)
		}

		res := []domain.Post{}
		if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
			t.Errorf("unexpected response: %v", err)
		}

		if res[0].ID != "test" {
			t.Errorf("want id %s, got %s", "test", res[0].ID)
		}
	})

	t.Run("it responds to post calls towards /posts", func(t *testing.T) {
		postsService := domain.NewPostsService(&PostsStorage{}, &PostsPublisher{})

		router := Router(postsService)

		w := httptest.NewRecorder()
		data, _ := json.Marshal(domain.NewPost{
			Title: "hello",
			Body:  "world",
		})
		req, _ := http.NewRequest(http.MethodPost, "/posts", bytes.NewReader(data))

		router.ServeHTTP(w, req)

		if w.Code != 200 {
			t.Errorf("want status code %d, got %d", 200, w.Code)
		}
	})
}
