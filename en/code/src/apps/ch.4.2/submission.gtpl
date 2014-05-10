{{define "submission"}}<!DOCTYPE html>
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
	Profile successfully submitted.
	{{end}}
	</body>
</html>
{{end}}