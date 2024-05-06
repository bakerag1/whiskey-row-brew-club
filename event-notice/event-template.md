---
title: {{.PostTitle}}
date: {{.Date}}
startDate: {{ .Start }}
endDate:  {{ .End }}
location: {{ .Location }}
event: true
---

![image](event.png)
 
{{.Title}}  
{{.Time}}  
{{.Location}}  
Gauntlet: {{.Gauntlet}}  
Ed Topic: {{.Ed}}  
{{with .Notes}}  
Notes:  
  {{ range .}}
  * {{.}}  
  {{end}}
{{end}}