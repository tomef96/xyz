package domain

import "context"

type NewPost struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (p NewPost) Valid() bool {
	return p.Title != "" && p.Body != ""
}

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (p *Post) Version() int {
	return 1
}

type PostsStorage interface {
	Store(ctx context.Context, post NewPost) (Post, error)
	List(ctx context.Context) ([]Post, error)
}

type PostsPublisher interface {
	Publish(ctx context.Context, post Post) error
}

type PostsService interface {
	HandleNewPost(ctx context.Context, post NewPost) error
	HandleListPosts(ctx context.Context) ([]Post, error)
}

type postsService struct {
	storage   PostsStorage
	publisher PostsPublisher
}

func NewPostsService(storage PostsStorage, publisher PostsPublisher) PostsService {
	return &postsService{
		storage,
		publisher,
	}
}

func (s *postsService) HandleNewPost(ctx context.Context, post NewPost) error {
	storedPost, err := s.storage.Store(ctx, post)
	if err != nil {
		return err
	}

	err = s.publisher.Publish(ctx, storedPost)
	if err != nil {
		return err
	}

	return nil
}

func (s *postsService) HandleListPosts(ctx context.Context) ([]Post, error) {
	return s.storage.List(ctx)
}
