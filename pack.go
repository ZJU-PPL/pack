package main

import (
	"log"
	"os"
	"path/filepath"
	"zjuppl/stu/pack/util"
	"bytes"
)

func contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func getFiles(ignore_dirs []string, ignore_files []string) []string {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatalln(err)
	}
	var ret []string
	for _, f := range files {
		name := f.Name()
		if f.IsDir() {
			if !contains(ignore_dirs, name) {
				filepath.Walk(name, func(path string, info os.FileInfo, err error) error {
					if err != nil {
						log.Fatalln(err)
					}
					if !info.IsDir() {
						ret = append(ret, path)
					}
					return nil
				})
			}
		} else {
			if !contains(ignore_files, name) {
				ret = append(ret, name)
			}
		}
	}
	return ret
}

func main() {

	ignore_dirs := []string{"_build", ".git"}
	myname := filepath.Base(os.Args[0])
	ignore_files := []string{myname}

	files := getFiles(ignore_dirs, ignore_files)

	var b bytes.Buffer

	err := util.CreateArchive(files, &b)
	if err != nil {
		log.Fatalln("Error creating archive:", err)
	}

	res, err := util.Encrypt(b.Bytes())
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create("stu-code.pgp.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	_, err = f.WriteString(res)
	if err != nil {
		log.Fatalln(err)
	}

}
