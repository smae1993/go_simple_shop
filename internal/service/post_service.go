package service

import (
	"context"
	"errors"
	"strings"

	"github.com/smae1993/go_simple_shop/internal/database"
)

type PostService struct {
	queries *database.Queries
}

func NewPostService(q *database.Queries) *PostService {
	return &PostService{queries: q}
}

// متدهای عمومی (Public)
func (s *PostService) GetRecentPosts(ctx context.Context, limit int64) ([]database.Post, error) {
	return s.queries.ListPublishedPosts(ctx, limit)
}

func (s *PostService) GetBySlug(ctx context.Context, slug string) (database.Post, error) {
	return s.queries.GetPostBySlug(ctx, slug)
}

// متدهای پنل ادمین (Admin)
func (s *PostService) GetAllForAdmin(ctx context.Context) ([]database.Post, error) {
	return s.queries.ListAllPostsAdmin(ctx)
}

func (s *PostService) GetByID(ctx context.Context, id int64) (database.Post, error) {
	return s.queries.GetPostByID(ctx, id)
}

// ذخیره یا بروزرسانی پست (اگر ID صفر باشد می‌سازد، وگرنه ویرایش می‌کند)
func (s *PostService) SavePost(ctx context.Context, id int64, title, slug, summary, content string, isPublished bool) error {
	title = strings.TrimSpace(title)
	slug = strings.TrimSpace(strings.ToLower(strings.ReplaceAll(slug, " ", "-")))

	if title == "" || slug == "" {
		return errors.New("عنوان و اسلاگ (Slug) الزامی هستند")
	}

	if id == 0 {
		_, err := s.queries.CreatePost(ctx, database.CreatePostParams{
			Title:       title,
			Slug:        slug,
			Summary:     summary,
			Content:     content,
			IsPublished: isPublished,
		})
		return err
	}

	_, err := s.queries.UpdatePost(ctx, database.UpdatePostParams{
		Title:       title,
		Slug:        slug,
		Summary:     summary,
		Content:     content,
		IsPublished: isPublished,
		ID:          id,
	})
	return err
}

func (s *PostService) Delete(ctx context.Context, id int64) error {
	return s.queries.DeletePost(ctx, id)
}
