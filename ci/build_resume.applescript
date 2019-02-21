set resume to POSIX file "/Users/atec/www/resume.pdf"
tell application "Pages"
	set doc to open "/Users/atec/www/resume.pages"
	export doc to file resume as PDF
	close doc saving no
	quit saving no
end tell