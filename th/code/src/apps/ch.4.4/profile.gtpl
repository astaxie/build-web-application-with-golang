{{define "profile"}}<!DOCTYPE html>
<html>
	<style>
	.row{
		display: table-row;
	}
	.cell{
		display: table-cell;
	}
	.required{
		color: red
	}
	</style>
	<body>				
		<h2>Profile Setup:</h2>
		<form action="/checkprofile" method="POST">
				<div class="row">
					<div class="cell"><span class="required">*</span>User Name:</div>
					<div class="cell"><input type="text" name="username" id="username" required/></div>
				</div>
				<div class="row">
					<div class="cell"><span class="required">*</span>Age:</div>
					<div class="cell"><input type="number" min="13" max="130" name="age" id="age" size="3" required/></div>
				</div>
				<div class="row">
					<div class="cell"><span class="required">*</span>Email:</div>
					<div class="cell"><input type="email" name="email" id="email" placeholder="john@example.com" required/></div>
				</div>
				<div class="row">
					<div class="cell"><span class="required">*</span>Birth day:</div>
					<div class="cell">
						<input type="date" name="birthday" id="birthday" placeholder="MM/DD/YYYY" required/>
					</div>
				</div>
				<div class="row">
					<div class="cell">Gender:</div>
					<div class="cell">
						<label for="gender_male">
							Male: <input type="radio" name="gender" value="m" id="gender_male"/>
						</label>
						<label for="gender_female">
							Female: <input type="radio" name="gender" value="f" id="gender_female"/>
						</label>
						<label for="gender_na">
							N/A: <input type="radio" name="gender" value="na" id="gender_na"/>
						</label>
					</div>
				</div>
				<div class="row">
					<div class="cell">Siblings:</div>
					<div class="cell">
						<label for="sibling_male">
							Brother: <input type="checkbox" name="sibling" value="m" id="sibling_male"/>
						</label>
						<label for="sibling_female">
							Sister: <input type="checkbox" name="sibling" value="f" id="sibling_female"/>
						</label>
					</div>
				</div>
			<div class="row">
				<div class="cell">Shirt Size:</div>
				<div class="cell">
					<select id="shirt_size" >
						<option></option>
						<option value="s">Small</option>
						<option value="m">Medium</option>
						<option value="l">Large</option>
						<option value="xl">X-Large</option>
						<option value="xxl">XX-Large</option>
					</select>
				</div>
			</div>
			<div class="row">
				<div class="cell">Chinese Name:</div>
				<div class="cell"><input type="text" name="chineseName" id="chineseName"/></div>
			</div>
			<br/>
			<span class="required">*</span>Required
			<br/>
			<input type="hidden" name="token" value="{{.Token}}"/>
			<input type="submit" value="Submit" id="submitBtn"/>
		</form>
	</body>
</html>
{{end}}
