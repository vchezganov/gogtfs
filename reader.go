package gogtfs

import (
	"github.com/vchezganov/gocsv"
	"github.com/vchezganov/gogtfs/model"
	"io"
	"os"
	"path/filepath"
)

type Reader[T model.Modelable] struct {
	f         io.Closer
	csvReader *gocsv.Reader[T]
}

func (r *Reader[T]) Next() (*T, error) {
	return r.csvReader.Next()
}

func (r *Reader[T]) Close() error {
	return r.f.Close()
}

func NewReader[T model.Modelable](ioReader io.ReadCloser) (*Reader[T], error) {
	csvReader, err := gocsv.NewReader[T](ioReader)
	if err != nil {
		return nil, err
	}

	return &Reader[T]{
		ioReader,
		csvReader,
	}, nil
}

func NewReaderFromFolder[T model.Modelable](folder string) (*Reader[T], error) {
	var m T

	p := filepath.Join(folder, m.TableFileName())
	ioReader, err := os.Open(p)
	if err != nil {
		return nil, err
	}

	return NewReader[T](ioReader)
}
