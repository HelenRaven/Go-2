package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

type File struct {
	Path    string
	Name    string
	Size    int64
	Deleted bool
}

func (f File) Eql(s File) bool {
	return f.Name == s.Name && f.Size == s.Size
}

type Filez []File

func (f Filez) PrintMe() {
	for _, v := range f {
		fmt.Println(v.Path)
	}
}

func (f Filez) PrintDeleted() {
	for _, v := range f {
		if v.Deleted {
			fmt.Println(v.Path, "Deleted")
		} else {
			fmt.Println(v.Path)
		}
	}
}

func (f Filez) Has(file File) bool {
	for _, v := range f {
		if v.Eql(file) {
			return true
		}
	}
	return false
}

func (f Filez) Len() int {
	return len(f)
}

func (f Filez) Less(i, j int) bool {
	return f[i].Size < f[j].Size
}

func (f Filez) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

type Configuration struct {
	Path             string
	DeleteDuplicates bool
}

func LoadConfig() (*Configuration, error) {
	var conf Configuration

	dir, err := os.Getwd()
	if err != nil {
		return &conf, err
	}

	flag.StringVar(&conf.Path, "path", dir, "path to find duplicates")
	flag.BoolVar(&conf.DeleteDuplicates, "r", false, "flag: delete duplicate files")
	flag.Parse()

	return &conf, nil
}

func LoadFilesToArr(path string) ([]File, error) {
	files := []File{}

	var wg = sync.WaitGroup{}
	var mu sync.Mutex

	err := filepath.Walk(path,
		func(filepath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				wg.Add(1)
				go func() {
					defer wg.Done()
					f := File{
						Name:    info.Name(),
						Size:    info.Size(),
						Path:    filepath,
						Deleted: false,
					}
					mu.Lock()
					files = append(files, f)
					mu.Unlock()
				}()
			}
			return nil
		})
	if err != nil {
		return files, err
	}
	wg.Wait()
	return files, nil
}

func OnlyDuplicates(files []File) []File {
	if len(files) == 0 {
		return files
	}
	var wg = sync.WaitGroup{}
	var mu sync.Mutex

	result := make([]File, 0, len(files))

	for i := 0; i < len(files); i++ {
		mu.Lock()
		has := Filez(result).Has(files[i])
		mu.Unlock()
		if has {
			continue
		}
		get := false
		for j := i + 1; j < len(files); j++ {
			if files[i].Eql(files[j]) {
				wg.Add(1)
				go func(i, j int) {
					defer wg.Done()
					mu.Lock()
					result = append(result, files[j])
					mu.Unlock()
				}(i, j)
				get = true
			}
		}
		if get {
			mu.Lock()
			result = append(result, files[i])
			mu.Unlock()
		}
	}
	wg.Wait()
	sort.Sort(Filez(result))
	return result
}

func RemoveDuplicates(files []File) {
	for i := 0; i < len(files); i++ {
		for j := i + 1; j < len(files); j++ {
			if files[i].Eql(files[j]) && !files[j].Deleted {
				e := os.Remove(files[j].Path)
				if e != nil {
					log.Println(e)
				} else {
					files[j].Deleted = true
				}

			}
		}
	}
	sort.Sort(Filez(files))
}

func main() {
	conf, err := LoadConfig()
	if err != nil {
		log.Println(err)
		return
	}

	files, err := LoadFilesToArr(conf.Path)
	if err != nil {
		log.Println(err)
		return
	}

	dupl := OnlyDuplicates(files)
	log.Println("----------------------DUPLICATES-------------------")
	Filez(dupl).PrintMe()

	var a string
	if conf.DeleteDuplicates && len(dupl) > 0 {
		fmt.Print("Duplicates will be removed. Are u sure? [y/n]:")
		fmt.Scanln(&a)
	}

	if strings.ToLower(a) == "y" {
		RemoveDuplicates(dupl)
		log.Println("---------------------RESULT---------------------")
		Filez(dupl).PrintDeleted()
	}
}
