package model

import (
	"context"
	"fmt"
	"os"
	"time"
)

type Model struct {
	path     string
	data     []byte
	fileInfo os.FileInfo
	ctx      context.Context
}

func New(ctx context.Context, path string) *Model {
	m := Model{path: path}
	go m.watch(ctx)

	return &m
}

func (m *Model) GetData() []byte {
	return m.data
}

func (m *Model) watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			m.checkFile()

			time.Sleep(500 * time.Millisecond)
		}
	}
}

func (m *Model) checkFile() {
	fileInfo, err := os.Stat(m.path)
	if err != nil {
		fmt.Printf("[model] read file error: %v\n", err)
		return
	}

	if m.fileInfo != nil && fileInfo.Size() == m.fileInfo.Size() && fileInfo.ModTime() == m.fileInfo.ModTime() {
		return
	}

	fmt.Printf("[model] file update detected at %v\n", time.Now())

	data, err := os.ReadFile(m.path)
	if err != nil {
		fmt.Printf("[model] file read error %v\n", err)
		return
	}

	m.fileInfo = fileInfo
	m.data = data

	return
}
