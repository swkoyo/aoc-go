package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX_SIZE = 100000

type File struct {
	Key  string
	Size int
}

type Dir struct {
	Key       string
	ParentDir *Dir
	SubDir    []*Dir
	Files     []*File
}

func CreateDir(key string, parent *Dir) *Dir {
	return &Dir{
		Key:       key,
		ParentDir: parent,
		SubDir:    make([]*Dir, 0),
		Files:     make([]*File, 0),
	}
}

func (this *Dir) HasSubDir(key string) (bool, *Dir) {
	var subDir *Dir = nil
	for _, dir := range this.SubDir {
		if dir.Key == key {
			subDir = dir
			break
		}
	}
	if subDir == nil {
		return false, nil
	}
	return true, subDir
}

func (this *Dir) AddSubDir(dir *Dir) {
	this.SubDir = append(this.SubDir, dir)
}

func (this *Dir) AddFile(file *File) {
	this.Files = append(this.Files, file)
}

func (this *Dir) SumDir(sizes *[]int) int {
	size := 0

	for _, file := range this.Files {
		size += file.Size
	}

	for _, dir := range this.SubDir {
		size += dir.SumDir(sizes)
	}

    if size <= MAX_SIZE {
        *sizes = append(*sizes, size)
    }

	return size
}

func CreateFile(key string, size int) *File {
	return &File{
		Key:  key,
		Size: size,
	}
}

func IsChangeDir(s string) (bool, string) {
	if s[:4] != "$ cd" {
		return false, ""
	}
	return true, s[5:]
}

func IsDir(s string) (bool, string) {
	if s[:3] != "dir" {
		return false, ""
	}
	return true, s[4:]
}

func IsFile(s string) (bool, string, int) {
	if s[:3] == "dir" {
		return false, "", 0
	}
	file := strings.Split(s, " ")
	size, err := strconv.Atoi(file[0])
	if err != nil {
		panic(err)
	}
	key := file[1]
	return true, key, size
}

func IsCommand(s string) bool {
	return s[:1] == "$"
}

func P1() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	root := CreateDir("/", nil)
	curr := root
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	i := 1

	for i < len(lines) {
		line := lines[i]
		if ok, dirKey := IsChangeDir(line); ok {
			if dirKey == ".." {
				curr = curr.ParentDir
			} else {
				if ok, subDir := curr.HasSubDir(dirKey); ok {
					curr = subDir
				} else {
					dir := CreateDir(dirKey, curr)
					curr.AddSubDir(dir)
					curr = dir
				}
			}
			i++
		} else {
			i++
			for i < len(lines) && !IsCommand(lines[i]) {
				if ok, fileKey, fileSize := IsFile(lines[i]); ok {
					file := CreateFile(fileKey, fileSize)
					curr.AddFile(file)
				} else if ok, dirKey := IsDir(lines[i]); ok {
					dir := CreateDir(dirKey, curr)
					curr.AddSubDir(dir)
				}
				i++
			}
		}
	}

    sizes := make([]int, 0)
    result := 0

	root.SumDir(&sizes)

    for _, val := range sizes {
        result += val
    }

    fmt.Println(result)
}

