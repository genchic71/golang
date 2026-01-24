package main

import (
	"context"
	"errors"
	"fmt"
)

// zad13
type Worker interface {
	Do(ctx context.Context) error
}

type FileProcessor struct {
	FilePath string
}

func (fp *FileProcessor) Do(ctx context.Context) error {
	if fp.FilePath == "/path/to/file" {
		return nil
	} else {
		return errors.New("file not found")
	}
}

type NetworkFetcher struct {
	URL string
}

func (nf *NetworkFetcher) Do(ctx context.Context) error {
	if nf.URL == "https/ok" {
		return nil
	} else {
		return errors.New("not ok")
	}
}

func main() {

	ctx := context.Background()

	workers := []Worker{
		&FileProcessor{FilePath: "/path/to/file.txt"},
		&NetworkFetcher{URL: "https://example.com"},
	}

	for i, worker := range workers {
		fmt.Printf("Запуск воркера %d...\n", i+1)
		if err := worker.Do(ctx); err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		} else {
			fmt.Println("Успешно!")
		}
	}
}
