package controllers

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	_ "github.com/lib/pq"

	beego "github.com/beego/beego/v2/server/web"
)

// Profile is the same as User...
type User struct {
	Id        int       `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Username  string    `json:"username" orm:"unique"`
	Email     string    `json:"email" orm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "dbname=firstapp host=localhost user=postgres password=kaak port=5432 sslmode=disable", 30)
	orm.RegisterModel(new(User))

	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}

	// beego.BConfig.WebConfig.Session.SessionOn = true
	// beego.BConfig.WebConfig.Session.SessionProvider = "postgresql"
	// beego.BConfig.WebConfig.Session.SessionProviderConfig = "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
	// beego.BConfig.WebConfig.Session.SessionOn = true
	// beego.BConfig.WebConfig.Session.SessionProvider = "file"
	// beego.BConfig.WebConfig.Session.SessionProviderConfig = "/tmp"

}

// Login ...
func (c *MainController) Login() {
	c.ActiveContent("authorization/login")

	back := strings.Replace(c.Ctx.Input.Param(":back"), ">", "/", -1)
	fmt.Println("back is ", back)

	if c.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()

		var p User
		p.Email = c.GetString("Email")
		p.Password = c.GetString("Password")

		valid := validation.Validation{}
		valid.Email(p.Email, "Email")
		valid.MinSize(p.Password, 8, "Password")
		if valid.HasErrors() {
			errormap := []string{}
			for _, err := range valid.Errors {
				errormap = append(errormap, "Validation failed on"+err.Key+": "+err.Message+"\n")
			}
			c.Data["Errors"] = errormap
			return
		}
		fmt.Println("Authorization is ", p.Email, ":", p.Password)

		//* Read password from database *//

		o := orm.NewOrm()
		o.Using("default")

		user := User{Email: p.Email, Password: p.Password}
		err := o.Read(&user, "email")

		fmt.Println("================user.Password is: ", user.Password)
		fmt.Println("================user.Email is: ", user.Email)

		if err == nil {

			if user.Id == 0 {
				fmt.Println("Account not verified")
				flash.Notice("Account not verified")
				flash.Store(&c.Controller)
				c.Redirect("/notice", 302)
				return
			}

		}

		ok := CheckPasswordAndHash(p.Password, user.Password)
		fmt.Printf("================user.Password-%s && p.Password-%s == %t \n", user.Password, p.Password, ok)
		if !ok {
			fmt.Println("Password is not correct")
			flash.Notice("No such user/email")
			flash.Store(&c.Controller)
			return
		}

		//* Create session and go back to previous page *//
		m := make(map[string]interface{})
		m["first"] = user.Firstname
		m["username"] = user.Username
		m["timestamp"] = time.Now()
		c.SetSession("pqweb", m)
		c.Redirect("/"+back, 302)

		fmt.Println("username is ", m["username"])
		fmt.Println("logged in at ", m["timestamp"])
	}

}

func (c *MainController) Logout() {
	c.ActiveContent("logout")
	c.DelSession("pqweb")
	c.Redirect("/home", 302)
}

// SignUp register new user
func (c *MainController) SignUp() {
	c.ActiveContent("authorization/signup")

	if c.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()

		var p User
		p.Firstname = c.GetString("Firstname")
		p.Lastname = c.GetString("Lastname")
		p.Username = c.GetString("username")
		p.Email = c.GetString("Email")
		pass := c.GetString("Password")

		//* validation *//
		valid := validation.Validation{}
		valid.Required(p.Firstname, "Firstname")
		valid.Required(p.Lastname, "Lastname")
		valid.Required(p.Username, "username")
		valid.Email(p.Email, "Email")
		valid.MinSize(pass, 8, "Password")
		if valid.HasErrors() {
			errormap := []string{}
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
				errormap = append(errormap, err.Key+err.Message+"\n")
			}
			fmt.Println("errormap: ", errormap)
			return
		}

		//* hash password *//
		hash, err := HashPassword(pass)
		if err != nil {
			log.Fatal("auth -hash-", err)
		}
		p.Password = hash

		//* Inserting to database *//
		user := User{Firstname: p.Firstname, Lastname: p.Lastname, Username: p.Username,
			Email: p.Email, Password: p.Password, CreatedAt: time.Now()}

		o := orm.NewOrm()
		o.Using("default")

		_, err = o.Insert(&user)
		if err != nil {
			flash.Notice(p.Email + " already registered")
			flash.Store(&c.Controller)
			c.Redirect("/notice", 302)
			return
		}

		// db, err := orm.GetDB()
		// if err != nil {
		// 	fmt.Println("get default DataBase")
		// }
		// fmt.Println(db)

		flash.Notice("Your account has been created.")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}

}
