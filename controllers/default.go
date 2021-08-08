package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
)

// MainController ...
type MainController struct {
	beego.Controller
}

// ActiveContent MainController
func (c *MainController) ActiveContent(view string) {
	c.Layout = "Basic-layout.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "Header.tpl"
	c.LayoutSections["Footer"] = "Footer.tpl"
	c.TplName = view + ".tpl"

	sess := c.GetSession("pqweb")
	if sess != nil {
		c.Data["InSession"] = 1
		m := sess.(map[string]interface{})
		c.Data["Firstname"] = m["first"]
	}
}

// Get MaincPage
func (c *MainController) Get() {
	c.ActiveContent("index")

	// sess := c.GetSession("pqweb")
	// if sess == nil {
	// 	c.Redirect("/home", 302)
	// 	return
	// }

	// m := sess.(map[string]interface{})
	// fmt.Println("username is", m["username"])
	// fmt.Println("logged in at", m["timestamp"])

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
}

// Notice ...
func (c *MainController) Notice() {
	c.ActiveContent("notice")

	flash := beego.ReadFromRequest(&c.Controller)
	if n, ok := flash.Data["notice"]; ok {
		c.Data["notice"] = n
	} else if n, ok = flash.Data["error"]; ok {
		c.Data["notice"] = n
	}
}

// HashPassword is a encryption function from bcrypt lib
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordAndHash is checks password and hash, output bool
func CheckPasswordAndHash(password, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
