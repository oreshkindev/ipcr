package converter

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/chai2010/webp"
)

type Converter interface {
	Convert(i, o string, quality int) error
}

func New() Converter {
	return &converterImpl{}
}

type converterImpl struct{}

func (c *converterImpl) Convert(i, o string, quality int) error {
	switch filepath.Ext(i) {
	case ".jpg", ".jpeg":
		// decode .jpg / .jpeg
		img, err := DecodeJPEG(i)
		if err != nil {
			return err
		}
		// compress .jpg
		if err := EncodeJPG(img, o, quality); err != nil {
			return fmt.Errorf("failed to EncodeJPG %s; %w", o, err)
		}
		// convert to .webp
		if err := EncodeWEBP(img, o, quality); err != nil {
			return fmt.Errorf("failed to EncodeWEBP %s; %w", o, err)
		}
		// removing the original image
		if err := Remove(i); err != nil {
			return fmt.Errorf("failed to Remove %s; %w", i, err)
		}
	case ".png":
		// decode .png
		img, err := DecodePNG(i)
		if err != nil {
			return err
		}
		// convert to .jpg и вызываем новое событие
		if err := EncodeJPG(img, o, quality); err != nil {
			return fmt.Errorf("failed to EncodeJPG %s; %w", o, err)
		}
		// convert to .webp
		if err := EncodeWEBP(img, o, quality); err != nil {
			return fmt.Errorf("failed to EncodeWEBP %s; %w", o, err)
		}
		// removing the original image
		if err := Remove(i); err != nil {
			return fmt.Errorf("failed to Remove %s; %w", i, err)
		}
	default:
		return fmt.Errorf("unsupported file type %s", i)
	}

	return nil
}

// convert to .jpg
func DecodeJPEG(i string) (image.Image, error) {
	inf, err := os.Open(i)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s file to read; %s", i, err)
	}
	defer inf.Close()

	// decode .jpeg
	img, err := jpeg.Decode(inf)
	if err != nil {
		return nil, fmt.Errorf("failed to decode %s file; %w", i, err)
	}

	return img, nil
}

// decode .png
func DecodePNG(i string) (image.Image, error) {
	inf, err := os.Open(i)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s file to read; %w", i, err)
	}
	defer inf.Close()

	// decode .png
	img, err := png.Decode(inf)
	if err != nil {
		return nil, fmt.Errorf("failed to decode %s file; %w", i, err)
	}

	return img, nil
}

// convert to .webp
func EncodeWEBP(img image.Image, o string, q int) error {
	var buf bytes.Buffer
	// Encode lossless webp
	if err := webp.Encode(&buf, img, &webp.Options{Lossless: false, Quality: float32(q)}); err != nil {
		return fmt.Errorf("failed to decode %s file; %w", o, err)
	}
	if err := os.WriteFile(o, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to open %s file to write; %w", o, err)
	}
	return nil
}

// convert to .jpg
func EncodeJPG(img image.Image, o string, q int) error {
	t, err := os.OpenFile(strings.ReplaceAll(o, filepath.Ext(o), ".jpg"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open %s file to write; %w", o, err)
	}
	defer t.Close()

	if err := jpeg.Encode(t, img, &jpeg.Options{Quality: q}); err != nil {
		return fmt.Errorf("failed to encode %s file; %w", o, err)
	}
	return nil
}

// removing the original image
func Remove(i string) error {
	if err := os.Remove(i); err != nil {
		return fmt.Errorf("failed to remove %s file; %w", i, err)
	}
	return nil
}
