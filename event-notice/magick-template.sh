magick -size 1080x1080 gradient:lime-green temp1.png
magick \
	-background none \
	-font Gentleman-Caller \
	-pointsize 120 \
	-gravity center label:"{{.Title}}\n" \
	-pointsize 90 label:"{{.Time}} at\n{{.Location}}\n" \
	-pointsize 50 label:"\n{{.Gauntlet}} Comp\n{{.Ed}} Ed. Topic" \
	-append temp2.png
magick temp1.png temp2.png \
	-gravity center \
	-composite out/event.png
rm temp1.png temp2.png