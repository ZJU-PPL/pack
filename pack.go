package main

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"zjuppl/stu/pack/util"
)


func topLevelJudge(topLevelAllowedReg []string, name string) bool {
	for _, r := range topLevelAllowedReg {
		b, _ := regexp.MatchString(r , name)
		if b {
			return true
		}
	}
	return false
}

func getFiles(topLevelAllowedReg []string) []string {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatalln(err)
	}
	var fir []string
	for _, f := range files {
		name := f.Name()
		if !topLevelJudge(topLevelAllowedReg, name) {
			continue
		}
		if f.IsDir() {
			filepath.Walk(name, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					log.Fatalln(err)
				}
				if !info.IsDir() {
					pathSlash := filepath.ToSlash(path)
					fir = append(fir, pathSlash)
				}
				return nil
			})
		} else {
			fir = append(fir, name)
		}
	}
	return fir
}

func main() {

	outputName := "stu-code.pgp.txt"
	topLevelAllowedReg := []string{"lib", "bin", "test", "dune-project", `.*\.opam`}
	
	files := getFiles(topLevelAllowedReg)

	var b bytes.Buffer

	err := util.CreateArchive(files, &b)
	if err != nil {
		log.Fatalln("Error creating archive:", err)
	}

	res, err := util.Encrypt(b.Bytes())
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create(outputName)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	_, err = f.WriteString(res)
	if err != nil {
		log.Fatalln(err)
	}

}
