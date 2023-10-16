package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/v3/shared"
	"html/template"
)

const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		<h1>Login</h1>
		<form action="" method="post">
			<input type="email" name="email" placeholder="email" />
			<input type="password" name="password" placeholder="password" />
			<input type="submit" value="Login" />
		</form>
	</body>
</html>
`

func (a *Api) LoginPage(c *fiber.Ctx) error {

	// Get ticket
	ticket := shared.GetTicket(c)

	// Do if user is authenticated
	return ticket.DoIfAuthenticated(
		func() error {

			return renderLoginPage(c)

		},
		func() error {

			return renderLoginPage(c)

		},
	)

}

func renderLoginPage(c *fiber.Ctx) error {
	// set content type
	c.Set("Content-Type", "text/html")
	// create template
	t, err := template.New("webpage").Parse(tpl)
	// check err
	if err != nil {
		return err
	}
	// page data
	data := struct {
		Title string
	}{
		Title: "Login Page",
	}
	// execute template
	err = t.Execute(c, data)
	// check err
	if err != nil {
		return err
	}
	// return nil
	return nil
}
