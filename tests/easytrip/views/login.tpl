<form method="POST">
    <div>Enter your login</div>
    <input type = "text" name="login">
    <div>Enter your password</div>
    <input type = "password" name="password">
    <button type="submit">login</button>
</form>
{{.Errors}}
{{.Username}}