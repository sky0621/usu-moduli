package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetseb4e4ea6bbc5b2ab38edf0740602254e99328232 = "# パッケージ別の各リポジトリでの使用バージョン一覧({{.Datetime}} 時点)\n\n#### ※ツール（ https://github.com/sky0621/usu-moduli ）による自動生成\n\n#### ★参考★\n\n##### https://github.com/Masterminds/glide\n\n##### https://glide.readthedocs.io/en/latest/versions/\n\n| Package {{range .ProjectNames}}| {{.}} {{end}}|\n| :--- {{range .ProjectNames}}| :--- {{end}}|\n{{range .Packages2}}| {{.Name}} {{range .Project2s}}| {{.Version}} {{end}}|\n{{end}}\n"
var _Assets9170159bf07815e0219401506fbf413c7c7c26a9 = "# リポジトリ別の使用モジュール一覧({{.Datetime}} 時点)\n\n#### ※ツール（ https://github.com/sky0621/usu-moduli ）による自動生成\n\n{{range .Projects}}## {{.Name}}\n\n| Name | Version |\n| :--- | :--- |\n{{range .Packages}}| {{.Name}} | {{.Version}} |\n{{end}}\n{{end}}\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"template"}, "/template": []string{"eachPackage.md", "eachProject.md"}}, map[string]*assets.File{
	"/template": &assets.File{
		Path:     "/template",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1523184785, 1523184785052891348),
		Data:     nil,
	}, "/template/eachPackage.md": &assets.File{
		Path:     "/template/eachPackage.md",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1523127472, 1523127472829919842),
		Data:     []byte(_Assetseb4e4ea6bbc5b2ab38edf0740602254e99328232),
	}, "/template/eachProject.md": &assets.File{
		Path:     "/template/eachProject.md",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1522947428, 1522947428151337076),
		Data:     []byte(_Assets9170159bf07815e0219401506fbf413c7c7c26a9),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1523184900, 1523184900209814268),
		Data:     nil,
	}}, "")
