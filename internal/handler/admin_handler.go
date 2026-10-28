package handler

import (
	"net/http"
	"strconv"

	"github.com/smae1993/go_simple_shop/internal/database"
	"github.com/smae1993/go_simple_shop/internal/service"
)

type AdminHandler struct {
	postService *service.PostService
	renderer    *Renderer
}

func NewAdminHandler(ps *service.PostService, rn *Renderer) *AdminHandler {
	return &AdminHandler{postService: ps, renderer: rn}
}

// ۱. لیست پست‌ها در ادمین (GET /admin/blog)
func (h *AdminHandler) List(w http.ResponseWriter, r *http.Request) {
	posts, err := h.postService.GetAllForAdmin(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.renderer.Render(w, "admin/list.html", map[string]any{
		"Title": "مدیریت مقالات",
		"Posts": posts,
	})
}

// ۲. صفحه ساخت یا ویرایش پست (GET /admin/blog/post?id=1)
// اگر پارامتر id در URL باشد حالت ویرایش است، وگرنه حالت ساخت پست جدید
func (h *AdminHandler) PostForm(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	var post database.Post
	pageTitle := "نوشتن مقاله جدید"

	if idStr != "" {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		existingPost, err := h.postService.GetByID(r.Context(), id)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		post = existingPost
		pageTitle = "ویرایش مقاله: " + post.Title
	} else {
		post.IsPublished = true // پیش‌فرض برای پست جدید
	}

	h.renderer.Render(w, "admin/form.html", map[string]any{
		"Title": pageTitle,
		"Post":  post,
	})
}

// ۳. ذخیره فرم ساخت/ویرایش (POST /admin/blog/post)
func (h *AdminHandler) SavePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)
	isPublished := r.FormValue("is_published") == "on"

	err := h.postService.SavePost(
		r.Context(),
		id,
		r.FormValue("title"),
		r.FormValue("slug"),
		r.FormValue("summary"),
		r.FormValue("content"),
		isPublished,
	)
	if err != nil {
		http.Error(w, "خطا در ذخیره: "+err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/admin/blog", http.StatusSeeOther)
}

// ۴. حذف پست (POST /admin/blog/post/{id}/delete)
func (h *AdminHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(r.PathValue("id"), 10, 64)
	_ = h.postService.Delete(r.Context(), id)
	http.Redirect(w, r, "/admin/blog", http.StatusSeeOther)
}
