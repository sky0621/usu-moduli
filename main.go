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

var projects = []*Project{}

// TODO 機能実現スピード最優先での実装なので要リファクタ
func main() {
	targetDir := flag.String("d", ".", "Target Directory")
	flag.Parse()

	// eachProject(*targetDir)
	eachPackage(*targetDir)
}

// ------------------------------------------------------------------------------------------------
// プロジェクト別のパッケージ一覧
// ------------------------------------------------------------------------------------------------
func eachProject(targetDir string) {
	err := filepath.Walk(targetDir, applyEachProject)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	tmpl := template.Must(template.ParseFiles("./eachProject.md"))
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, &Result{Datetime: time.Now().Format("2006-01-02 15:04"), Projects: projects})
	if err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}

func applyEachProject(path string, info os.FileInfo, err error) error {
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

	project := &Project{Name: prjName}

	Packages := []Package{}

	nowPackage := Package{}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.HasPrefix(txt, "- package:") {
			if nowPackage.Name != "" {
				Packages = append(Packages, nowPackage)
			}
			rtxt := strings.Replace(txt, "- package: ", "", -1)
			nowPackage = Package{Name: strings.Trim(rtxt, " ")}
		}
		ttxt := strings.Trim(txt, " ")
		if strings.HasPrefix(ttxt, "version:") {
			vtxt := strings.Replace(ttxt, "version:", "", -1)
			tvtxt := strings.Trim(vtxt, " ")
			if strings.Contains(tvtxt, "^") {
				tvtxt = strings.Replace(tvtxt, "^", "\\^", -1)
			}
			nowPackage.Version = tvtxt
		}
	}
	project.Packages = Packages

	projects = append(projects, project)

	return nil
}

type Result struct {
	Datetime string
	Projects []*Project
}

type Project struct {
	Name     string
	Packages []Package
}

type Package struct {
	Name    string
	Version string
}

// ------------------------------------------------------------------------------------------------
// パッケージ別の各プロジェクトでの使用バージョン一覧
// ------------------------------------------------------------------------------------------------
func eachPackage(targetDir string) {
	err := filepath.Walk(targetDir, applyEachProject)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	for _, project := range projects {
		prjName := project.Name
		project2s = append(project2s, &Project2{Name: prjName})
	}

	pkgMap := make(map[string]map[string]string)

	for _, project := range projects {
		prjName := project.Name
		for _, pkg := range project.Packages {
			pkgName := pkg.Name
			pkgVer := pkg.Version

			// まだパッケージ未保存
			if _, ok := pkgMap[pkgName]; !ok {
				pkgMap[pkgName] = make(map[string]string)
			}

			prjMap := pkgMap[pkgName]
			// まだプロジェクト（とバージョン）未保存
			if _, ok := prjMap[prjName]; !ok {
				prjMap[prjName] = pkgVer
			}
		}
	}

	var eachProject2s = []*Project2{}
	for pkgName, prjMap := range pkgMap {
		pkg2 := &Package2{Name: pkgName}
		for prjName, ver := range prjMap {
			for _, project2 := range project2s {
				eachProject2 := &Project2{Name: project2.Name}
				if project2.Name == prjName && ver != "" {
					eachProject2.Version = ver
				} else {
					eachProject2.Version = "　"
				}
				eachProject2s = append(eachProject2s, eachProject2)
			}
		}
		pkg2.Project2s = eachProject2s
		packages2 = append(packages2, pkg2)
	}

	tmpl := template.Must(template.ParseFiles("./eachPackage.md"))
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, &Result2{Datetime: time.Now().Format("2006-01-02 15:04"), Project2s: project2s, Packages2: packages2})
	if err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}

var project2s = []*Project2{}

var packages2 = []*Package2{}

type Result2 struct {
	Datetime  string
	Project2s []*Project2
	Packages2 []*Package2
}

type Package2 struct {
	Name      string
	Project2s []*Project2
}

type Project2 struct {
	Name    string
	Version string
}
