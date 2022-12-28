package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input_0.txt
var input0 string

//go:embed input_1.txt
var input1 string

//go:embed input_2.txt
var input2 string

func assertLen[T any, S ~[]T](s S, length int) S {
	if len(s) != length {
		log.Panicf("expected length %v got %v", length, s)
	}
	return s
}

func mustInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Panic(err)
	}
	return i
}

type Filename string

type MyFile struct {
	Name     Filename
	Size     int64
	IsDir    bool
	Children map[Filename]*MyFile
	Parent   *MyFile
}

func (mf *MyFile) String() string {
	//return fmt.Sprintf("%d:%s", mf.Size, string(mf.Name))
	return string(mf.Name)
}

func (mf *MyFile) getOrAddChildDir(name Filename) *MyFile {
	if _, ok := mf.Children[name]; !ok {
		mf.Children[name] = &MyFile{
			Name:     name,
			IsDir:    true,
			Children: make(map[Filename]*MyFile),
			Parent:   mf,
		}
	}
	return mf.Children[name]
}

func (mf *MyFile) getOrAddChildFile(name Filename, size int64) *MyFile {
	if _, ok := mf.Children[name]; !ok {
		mf.Children[name] = &MyFile{
			Name:   name,
			IsDir:  false,
			Parent: mf,
			Size:   size,
		}
	}
	return mf.Children[name]
}

func (mf *MyFile) calculateSize() int64 {
	if !mf.IsDir {
		return mf.Size
	}
	var s int64
	for _, f := range mf.Children {
		s += f.calculateSize()
	}
	return s
}

func (mf *MyFile) getAllSubdirs() []*MyFile {
	var subdirs []*MyFile
	for _, f := range mf.Children {
		if f.IsDir {
			subdirs = append(subdirs, f)
			subdirs = append(subdirs, f.getAllSubdirs()...)
		}
	}
	return subdirs
}

func getFS(in string) (root *MyFile) {
	rows := strings.Split(in, "\n")

	root = &MyFile{
		Name:     "/",
		IsDir:    true,
		Children: make(map[Filename]*MyFile),
		Parent:   nil,
	}

	pwd := root

	for i := 0; i < len(rows); i++ {
		r := rows[i]
		//log.Printf("processing row `%v`", r)
		switch {
		case r == "":
			continue
		case r == "$ cd /":
			pwd = root
		case r == "$ cd ..":
			pwd = pwd.Parent
		case strings.HasPrefix(r, "$ cd "):
			subdirName := strings.TrimPrefix(r, "$ cd ")
			subdir := pwd.getOrAddChildDir(Filename(subdirName))
			pwd = subdir
		case r == "$ ls":
			for i++; i < len(rows) && !strings.HasPrefix(rows[i], "$"); i++ {
				r = rows[i]
				switch {
				case r == "":
					break
				case strings.HasPrefix(r, "dir "):
					subdirName := strings.TrimPrefix(r, "dir ")
					_ = pwd.getOrAddChildDir(Filename(subdirName))
				default:
					splits := strings.Split(r, " ")
					_ = pwd.getOrAddChildFile(Filename(splits[1]), mustInt(splits[0]))
				}
			}
			i--
		}
	}
	return root
}

func part1(in string) int64 {
	root := getFS(in)
	allDirs := []*MyFile{root}
	allDirs = append(allDirs, root.getAllSubdirs()...)

	var sumSize int64
	for _, d := range allDirs {
		if d.calculateSize() <= 100000 {
			log.Printf("got %v %v", d, d.calculateSize())
			sumSize += d.calculateSize()
		}
	}

	return sumSize
}

func part2(in string) int64 {
	diskAvail := int64(70000000)
	desiredFree := int64(30000000)

	root := getFS(in)
	allDirs := []*MyFile{root}
	allDirs = append(allDirs, root.getAllSubdirs()...)

	diskUsed := root.calculateSize()
	log.Printf("root dir has size %v", diskUsed)

	bestDirSizeToRM := diskUsed
	for _, d := range allDirs {
		size := d.calculateSize()
		if size < bestDirSizeToRM && diskUsed-size <= diskAvail-desiredFree {
			log.Printf("dir %v is new best, brings usage to %v", d, diskUsed-size)
			bestDirSizeToRM = size
		}
	}

	return bestDirSizeToRM
}

func main() {

}
