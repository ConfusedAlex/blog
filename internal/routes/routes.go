package routes

import (
	"net/http"

	"codeberg.org/confusedalex/website/components"
	"codeberg.org/confusedalex/website/internal/content"
	"github.com/a-h/templ"
)

func NewRouter() http.Handler {
	router := http.NewServeMux()

	router.Handle("/", templ.Handler(components.Page(content.GetFile("index.md"))))
	router.Handle("/now", templ.Handler(components.Page(content.GetFile("now.md"))))
	router.HandleFunc("/posts/{slug}", routePosts)
	router.Handle("/posts/", templ.Handler(components.BlogPage(content.GetBlogPosts())))
	router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	// router.HandleFunc("/feed/", serveFeed)

	return router
}

func routePosts(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	var posts = content.GetBlogPosts()
	for _, post := range posts {
		if post.Slug == slug && post.Draft != true {
			components.ContentPage(post).Render(r.Context(), w)
		}
	}
}
