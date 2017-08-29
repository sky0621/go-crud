# リポジトリ別使用テーブル一覧【{{.Branch}}ブランチ】({{.Datetime}} 時点)

{{range .Headers}}| {{.}} {{end}} |
{{range .Headers}}| :--- {{end}} |
{{range .Bodies}}{{range .}}| {{.}} {{end}} |
{{end}}
