package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var deps = []*Dep{}

// TODO 機能実現スピード最優先での実装なので要リファクタ
func main() {
	targetDir := flag.String("d", ".", "Target Directory")
	flag.Parse()

	err := filepath.Walk(*targetDir, apply)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	tmpl := template.Must(template.ParseFiles("./tmpl.md"))
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, &Result{Datetime: time.Now().Format("2006-01-02 15:04"), Deps: deps})
	if err != nil {
		panic(err)
	}

	fmt.Println(buf.String())

	// fmt.Println(res)
}

func apply(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	if !strings.Contains(path, "glide.yaml") {
		return nil
	}

	seps := strings.Split(path, "/")
	prjName := seps[len(seps)-2]

	fp, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		if fp != nil {
			fp.Close()
		}
	}()

	dep := &Dep{Prj: prjName}

	pkgs := []Pkg{}

	nowPkg := Pkg{}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.HasPrefix(txt, "- package:") {
			if nowPkg.Name != "" {
				pkgs = append(pkgs, nowPkg)
			}
			rtxt := strings.Replace(txt, "- package: ", "", -1)
			nowPkg = Pkg{Name: strings.Trim(rtxt, " ")}
		}
		ttxt := strings.Trim(txt, " ")
		if strings.HasPrefix(ttxt, "version:") {
			vtxt := strings.Replace(ttxt, "version:", "", -1)
			tvtxt := strings.Trim(vtxt, " ")
			nowPkg.Version = tvtxt
		}
	}
	dep.Pkgs = pkgs

	deps = append(deps, dep)

	return nil
}

type Result struct {
	Datetime string
	Deps     []*Dep
}

type Dep struct {
	Prj  string
	Pkgs []Pkg
}

type Pkg struct {
	Name    string
	Version string
}
