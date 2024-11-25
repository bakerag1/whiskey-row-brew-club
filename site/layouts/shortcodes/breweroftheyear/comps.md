{{ range $.Site.Data.competitions.schedules }}
  <h2>{{ .year }}</h2>
  {{ range .comps }}
    <hr>
    {{ $title := .compType }}
    {{ range $.Site.Data.competitions.compTypes }}
      {{ if eq .id $title}}{{ $title = .displayName }}{{ end }}
    {{ end }}
    <b>{{ .date }}<b> - {{ $title }}<br>
    {{ if .info }}
      {{ .info }}
    {{ if .link }}<a href="{{ .link }}">details</a>{{ end }}
    <br>
    {{ end }}
    {{ with .style }}
      {{ if and (eq .guide "bjcp" ) (eq .guideYear 2021 ) .category}} 
        <a href="/styles/bjcp/2021/{{.category}}">{{ .category }}. {{ .name }}</a><br>
        {{ else }}{{if .name }}{{ .name }}<br>{{ end }}
      {{ end}}
    {{ end }}
    {{ if .entries }}&nbsp;&nbsp;{{ .entries }} entries<br>{{ end }}
    {{ with .winners }}
      {{ range . }}
        &nbsp;&nbsp;{{ .position }}: {{ .name }}
        {{ with .category }} <i>{{ . }}</i> {{ end }}
        {{ with .recipeLink}} - <a href="{{ .recipeLink }}">recipe</a>{{ end }}
        <br>
      {{ end }}
    {{ end }}
  {{ end }}
{{ end }}