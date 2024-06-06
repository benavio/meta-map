package fsdo

import (
	"fmt"
	"os"
	"time"

	"github.com/evanoberholster/imagemeta"
	"github.com/evanoberholster/imagemeta/exif2"
)

type Dada struct {
	Name      string
	Latitude  float64
	Longitude float64
	Date      string
}

// Get exif and formating date
func GetExif() ([][]Dada, error) {
	inf := "GetExif"
	arr := make([][]Dada, 2000, 4000)
	dir, err := os.Open("./images")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		// return nil, fmt.Errorf("failed read", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", inf, err)
	}

	for i, file := range files {
		el, err := DecodeExif(file.Name(), dir)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", inf, err)
		}
		t := time.Now().Format("02-01-2006")
		if el.GPS.Date().Format("02-01-2006")[:10] == "01-01-0001" {
			arr[i] = append(arr[i], Dada{file.Name(), el.GPS.Latitude(), el.GPS.Longitude(), t})
			continue
		}
		t2 := el.GPS.Date().Format("02-01-2006")
		arr[i] = append(arr[i], Dada{file.Name(), el.GPS.Latitude(), el.GPS.Longitude(), t2})
	}
	return arr, nil
}

// Open and decode files
func DecodeExif(fileName string, dir *os.File) (exif2.Exif, error) {
	inf := "decodeExif"
	f, err := os.Open("./images/" + fileName)
	if err != nil {
		return exif2.Exif{}, fmt.Errorf("%s : %w", inf, err)
	}
	defer f.Close()

	e, err := imagemeta.Decode(f)
	if err != nil {
		return exif2.Exif{}, fmt.Errorf("%s : %w", inf, err)
	}
	return e, nil
}
