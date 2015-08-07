{{define "upload"}}<!DOCTYPE html>
<html>
	<body>
		{{if .Errors}}
		<h2>Errors:</h2>
			<ol>
				{{range .Errors}}	
				<li>{{.}}</li>
				{{end}}
			</ol>	
	{{else}}
	File uploaded successfully.<br/>
	Note: Refreshing the page will produce a duplicate entry.
	{{end}}
	</body>
</html>
{{end}}
