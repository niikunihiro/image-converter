/*
Package image is image image
 */
package image

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)


type Converter struct {
	Src  string
	Dest string
}

func NewConverter(src, dest string) *Converter {
	e := new(Converter)
	e.Src = "." + src
	e.Dest = "." + dest
	return e
}

// Convert 引数はファイルパスと変換前後の拡張子
// ファイルの出力は元画像のあるディレクトリに行う
func (c Converter) Convert(srcPath string) error {
	sf, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer sf.Close()

	destPath := strings.Replace(srcPath, c.Src, c.Dest, 1)
	df, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer df.Close()

	img, _, err := image.Decode(sf)
	if err != nil {
		return err
	}

	switch filepath.Ext(destPath) {
	case ".png":
		err = png.Encode(df, img)
	case ".jpg", ".jpeg":
		err = jpeg.Encode(df, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	}
	if err != nil {
		return err
	}

	return nil
}
