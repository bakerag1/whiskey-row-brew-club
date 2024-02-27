{{ range $year, $winner := $.Site.Data.competitions.winners }}
  <h2>{{ $year }}</h2>
  {{ $w := index $winner 0 }}
  {{ if $w.post }}
    <a href="{{ $w.post }}">{{$w.name}}
  {{ else }}
    <p>{{ $w.name }} <p>
  {{ end }}
{{ end }}