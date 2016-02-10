package templates

const LoginPage = `
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8"/>
  <title> Go-Chat </title>
  <link rel="stylesheet" href="http://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css">
  <link rel="stylesheet" href="/assets/stylesheets/login.css">
  <link href='http://fonts.googleapis.com/css?family=Montserrat:400,700' rel='stylesheet' type='text/css'>
  <script src="//ajax.googleapis.com/ajax/libs/jquery/2.1.4/jquery.min.js"></script>
  <script src="http://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
</head>
<body style="background-color: #FAFAFF;">
  <div class="logo"></div>
  <div class="tagline"><h2>Don't think, Just go for it.</h2></div>
  <div class="login-block">
    <form action="/authorize_github" method="POST"><input class="btn_git socl_btn" type="submit" value="continue with github"/></form>
    <form action="/authorize_google" method="POST"><input class="btn_google socl_btn" type="submit" value="continue with google+"/></form>
  </div>
</body>
</html>
`
