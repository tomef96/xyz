package event

import (
	"context"

	"github.com/tomef96/coop/mastodon/config"
	"github.com/tomef96/coop/mastodon/domain"
)

type PostsPublisher struct {
	KafkaPublisher
}

func NewPostsPublisher() *PostsPublisher {
	return &PostsPublisher{
		*NewKafkaPublisher(config.KAFKA_TOPIC_POSTS),
	}
}

func (p *PostsPublisher) Publish(ctx context.Context, post domain.Post) error {
	return p.KafkaPublisher.Publish(ctx, Event{
		Version: post.Version(),
		Schema:  "post",
		Payload: post,
	})
}

func (p *PostsPublisher) Close() {
	p.KafkaPublisher.Close()
}
