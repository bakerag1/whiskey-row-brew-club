---
title: {{.PostTitle}}
date: {{.Date}}
author: bakerag
---

![image](event.png)

Whiskey Row Brew Club  
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