{{ range $.Site.Data.competitions.schedules }}
  <h2>{{ .year }}</h2>
  {{ range .comps }}
    {{ if or ( eq .compType "club" ) ( eq .compType "gauntlet" ) }}
      <b>{{ time.Format "January" .date }}<b> - 
      {{ with .style }} 
        {{ .name }}
      {{ end }}<br>
    {{ end }}
  {{ end }}
{{ end }}