{{ range $.Site.Data.standings }}
    <b>{{.place}}</b>: {{ .name }} - {{ .points }} points<br>
{{ end }}