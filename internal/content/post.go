package content

import (
	"bytes"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	highlighting "github.com/yuin/goldmark-highlighting/v2"

	"github.com/adrg/frontmatter"
	"github.com/yuin/goldmark"
)

type Post struct {
	Title string    `yaml:"title"`
	Slug  string    `yaml:"slug"`
	Date  time.Time `yaml:"date"`
	Draft bool      `yaml:"draft"`
	Body  string
}

func GetBlogPosts() []Post {
	ext := ".md"
	var posts []Post

	filepath.WalkDir("./content/blog", func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ext {
			var post = Post{}

			content, err := os.ReadFile(path)

			rest, err := frontmatter.Parse(bytes.NewReader(content), &post)

			if err != nil {
				log.Println("Could not read file!")
			}

			post.Body = MdToHtml(rest)
			posts = append(posts, post)
		}
		return nil
	})

	sort.Slice(posts, func(i, j int) bool {
		return posts[j].Date.Before(posts[i].Date)
	})

	for _, post := range posts {
		log.Println(post.Date)
	}

	return posts
}

func MdToHtml(content []byte) string {
	var htmlOutput bytes.Buffer

	markdown := goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithStyle("gruvbox"),
				highlighting.WithFormatOptions(),
			),
		),
	)

	if err := markdown.Convert(content, &htmlOutput); err != nil {
		log.Println("Converting from markdown to html failed!")
	}

	return htmlOutput.String()
}
