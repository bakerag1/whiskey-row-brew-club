magick \
	-size 1400 \
	-background none \
	-font Helldorado \
	-fill fuchsia \
	-gravity north \
	-pointsize 240 label:"{{.Name}}" \
	-font Arial \
	-fill white \
	-pointsize 60 \
	-gravity north caption:"{{.Overall_impression}}" \
	-append temp1.png
magick \
	-size 600 \
	-background none \
	-font Arial \
	-fill white \
	-pointsize 120 label:"OG: {{.Original_gravity.Minimum.Value}}-{{.Original_gravity.Maximum.Value}}    FG: {{.Final_gravity.Minimum.Value}}-{{.Final_gravity.Maximum.Value}}\nIBUs: {{.International_bitterness_units.Minimum.Value}}-{{.International_bitterness_units.Maximum.Value}}    SRM: {{.Color.Minimum.Value}}-{{.Color.Maximum.Value}}" \
	-gravity west \
	-append temp2.png
magick ../site/content/ghosts-of-the-west/post-background.png temp1.png \
	-gravity north \
	-composite temp3.png
magick temp3.png temp2.png \
	-gravity center \
	-geometry +400+200 \
	-composite posts/{{ .Style_id }}.png
rm temp*.png