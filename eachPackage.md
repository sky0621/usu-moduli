# パッケージ別の各リポジトリでの使用バージョン一覧({{.Datetime}} 時点)

#### ※ツール（ https://github.com/sky0621/usu-moduli ）による自動生成

| Package | {{range .Project2s}}{{.Name}} |{{end}}
| :--- | {{range .Project2s}}:--- |{{end}}
{{range .Packages2}}| {{.Name}} | {{range .Project2s}}{{.Version}} |{{end}}
{{end}}
