package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/niikunihiro/image-encode/image"
)

// arg ディレクトリ名
// -s source extension
// -d destination extension
func main() {
	src := flag.String("s", "jpg", "変換前拡張子")
	dest := flag.String("d", "png", "変換後拡張子")

	flag.Parse()

	dirs := flag.Args()
	if len(dirs) < 1 {
		err := errors.New("引数にディレクトリ名を指定してください")
		log.Fatal(err)
	}
	if _, err := os.Stat(dirs[0]); err != nil {
		log.Fatal(err)
	}

	c := image.NewConverter(*src, *dest)

	// ディレクトリを再帰的に処理する
	err := filepath.Walk(dirs[0], func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == c.Src {
			// パスを取得できるのでここで変換処理を呼ぶ
			return c.Convert(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
