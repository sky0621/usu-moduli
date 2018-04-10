# リポジトリ別の使用モジュール一覧({{.Datetime}} 時点)

#### ※ツール（ https://github.com/sky0621/usu-moduli ）による自動生成

{{range .Projects}}## {{.Name}}

| Name | Version |
| :--- | :--- |
{{range .Packages}}| {{.Name}} | {{.Version}} |
{{end}}
{{end}}
