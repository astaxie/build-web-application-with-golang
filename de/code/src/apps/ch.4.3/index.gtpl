<!doctype html>
<html>
	<body>
		<h2>Cross Site Scripting Attack Test</h2>
		{{if .}}
		Previous User Input: <br/>
		
		<code><pre>{{.}}</pre></code>
		{{end}}
		<form action="/">
			<label>
				User Input: 
				<input type="text" size=50 name="userinput" id="userinput"/>
			</label>
			<br/>
			<label>
				Escape Input: 
				<input type="checkbox" value="1" name="escape" id="escape"/>
			</label>
			<br/>
			<input type="submit" id="submitBtn" value="Submit"/>
		</form>
		<script type="text/javascript">
		var s = "<scri"+"pt>alert('pOwned by XSS.')</scri"+"pt>"
		document.getElementById("userinput").value = s;
		</script>
	</body>
</html>
