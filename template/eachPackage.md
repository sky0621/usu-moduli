# パッケージ別の各リポジトリでの使用バージョン一覧({{.Datetime}} 時点)

#### ※ツール（ https://github.com/sky0621/usu-moduli ）による自動生成

#### ★参考★

##### https://github.com/Masterminds/glide

##### https://glide.readthedocs.io/en/latest/versions/

| Package {{range .ProjectNames}}| {{.}} {{end}}|
| :--- {{range .ProjectNames}}| :--- {{end}}|
{{range .Packages2}}| {{.Name}} {{range .Project2s}}| {{.Version}} {{end}}|
{{end}}
