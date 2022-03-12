package main

import (
	"errors"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	file, _ := os.OpenFile(fromPath, os.O_RDONLY, 0o666)
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	if offset > fileInfo.Size() {
		return ErrOffsetExceedsFileSize
	}
	_, err = file.Seek(offset, 0)

	if err != nil {
		return err
	}

	if fileInfo.Size() == 0 {
		return ErrUnsupportedFile
	}

	if fileInfo.Size() < limit {
		limit = fileInfo.Size()
	}
	if fileInfo.Size()-offset < limit {
		limit = fileInfo.Size() - offset
	}
	if limit == 0 {
		limit = fileInfo.Size()
	}

	byteSlice := make([]byte, limit)
	bar := pb.New(int(fileInfo.Size()))
	bar.Start()
	rd := bar.NewProxyReader(file)
	_, err = rd.Read(byteSlice)
	bar.Finish()

	if err != nil {
		return err
	}

	destFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	wr := bar.NewProxyWriter(destFile)
	_, err = wr.Write(byteSlice)

	if err != nil {
		return err
	}

	return nil
}
