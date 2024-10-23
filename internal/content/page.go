package content

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func GetFile(file string) string {
	var output string

	filepath.WalkDir("./content", func(path string, d fs.DirEntry, err error) error {
		if matched, err := filepath.Match(file, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			fileContent, err := os.ReadFile(path)

			if err != nil {
				log.Println("Could not read file!")
			}

			output = MdToHtml(fileContent)
		}
		return nil
	})
	return output
}
