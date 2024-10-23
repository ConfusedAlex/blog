package cmd

import (
	"context"
	"log"
	"os"

	"codeberg.org/confusedalex/website/components"
	"codeberg.org/confusedalex/website/internal/content"

	"github.com/spf13/cobra"
)

var (
	build = &cobra.Command{
		Use:   "build",
		Short: "builds the blog",
		Run: func(cmd *cobra.Command, args []string) {
			buildSite()
		},
	}
)

func init() {
	rootCmd.AddCommand(build)
}

func buildSite() {
	os.MkdirAll("public/posts", os.ModePerm)
	os.CopyFS("public/assets", os.DirFS("assets"))

	posts := content.GetBlogPosts()

	for _, post := range posts {
		file, err := os.Create("public/posts/" + post.Slug + ".html")
		if err != nil {
			log.Fatalf("failed to create output file: %v", err)
		}

		err = components.ContentPage(post).Render(context.Background(), file)
		if err != nil {
			log.Fatalf("failed to write output file: %v", err)
		}
	}

	var files = [2]string{"now", "index"}

	for _, fileName := range files {
		file, err := os.Create("public/" + fileName + ".html")
		if err != nil {
			log.Fatalf("failed to create output file: %v", err)
		}

		err = components.Page(content.GetFile(fileName+".md")).Render(context.Background(), file)
		if err != nil {
			log.Fatalf("failed to write output file: %v", err)
		}
	}
}
