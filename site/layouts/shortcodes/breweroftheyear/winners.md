{{ range $.Site.Data.competitions.winners }}
  <b>{{ .year }}</b>
  {{ if .post }}
     - <a href="{{ .post }}">{{.name}}</a><br>
  {{ else }}
     - {{ .name }}<br>
  {{ end }}
{{ end }}