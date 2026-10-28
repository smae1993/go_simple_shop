package handler

import (
	"net/http"

	"github.com/smae1993/go_simple_shop/internal/service"
)

type PublicHandler struct {
	postService *service.PostService
	renderer    *Renderer
}

func NewPublicHandler(ps *service.PostService, rn *Renderer) *PublicHandler {
	return &PublicHandler{postService: ps, renderer: rn}
}

// ۱. صفحه اصلی سایت (GET /) - نمایش ۳ پست آخر
func (h *PublicHandler) Home(w http.ResponseWriter, r *http.Request) {
	// در Go 1.22 الگوی "/" تمام مسیرهای ناشناخته رو هم می‌گیره، پس چک می‌کنیم دقیقاً "/" باشه
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	posts, err := h.postService.GetRecentPosts(r.Context(), 3)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.renderer.Render(w, "public/index.html", map[string]any{
		"Title": "صفحه اصلی وبلاگ",
		"Posts": posts,
	})
}

// ۲. صفحه آرشیو کل بلاگ (GET /blog)
func (h *PublicHandler) BlogIndex(w http.ResponseWriter, r *http.Request) {
	posts, err := h.postService.GetRecentPosts(r.Context(), 100)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.renderer.Render(w, "public/blog.html", map[string]any{
		"Title": "آرشیو مقالات",
		"Posts": posts,
	})
}

// ۳. صفحه نمایش یک پست (GET /blog/{slug})
func (h *PublicHandler) SinglePost(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	post, err := h.postService.GetBySlug(r.Context(), slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	h.renderer.Render(w, "public/post.html", map[string]any{
		"Title": post.Title,
		"Post":  post,
	})
}
