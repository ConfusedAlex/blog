package content

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/feeds"
)

func CreateFeed(posts []Post) *feeds.Feed {
	feed := &feeds.Feed{
		Title:       "confusedalex's blog",
		Link:        &feeds.Link{Href: "https://confusealex.dev/blog"},
		Description: "my personal blog",
		Author:      &feeds.Author{Name: "confusedalex", Email: "hello@confusedalex.dev"},
	}

	var items []*feeds.Item

	for _, post := range posts {
		if !post.Draft {

			// link := templ.SafeURL(path.Join(post.Slug, "/"))

			items = append(items, &feeds.Item{
				Title:   post.Title,
				Created: post.Date,
				// Link:    "https://confusealex.dev",
			})
		}
	}

	feed.Items = items

	return feed
}

func ServeFeed(w http.ResponseWriter, r *http.Request) {
	rss, err := CreateFeed(GetBlogPosts()).ToRss()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, rss)
}
