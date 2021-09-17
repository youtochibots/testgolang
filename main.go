package main

/*

version with sessions bcrypt fmt

import (
         "github.com/gorilla/sessions"
         "golang.org/x/crypto/bcrypt"
         "html/template"
         "log"
         "net/http"
         "fmt"
		"os"
 )

 var encryptionKey = "something-very-secret"
 var loggedUserSession = sessions.NewCookieStore([]byte(encryptionKey))

 func init() {

         loggedUserSession.Options = &sessions.Options{
                 // change domain to match your machine. Can be localhost
                 // IF the Domain name doesn't match, your session will be EMPTY!
                 Domain:   "localhost",
                 Path:     "/",
                 MaxAge:   3600 * 3, // 3 hours
                 HttpOnly: true,
         }
 }

 const dashBoardPage = `<html><body>

  {{if .Username}}
          <p><b>{{.Username}}</b>, welcome to your dashboard! <a href="/logout">Logout!</a></p>
  {{else}}
          <p>Either your session has expired or you've logged out! <a href="/login">Login</a></p>
  {{end}}

  </body></html>`

 const logUserPage = `<html><body>
  {{if .LoginError}}<p style="color:red">Either username or password is not in our record! Sign Up?</p>{{end}}

  <form method="post" action="/login">
          {{if .Username}}
                   <p><b>{{.Username}}</b>, you're already logged in! <a href="/logout">Logout!</a></p>
          {{else}}
                  <label>Username:</label>
                  <input type="text" name="Username"><br>

                  <label>Password:</label>
                  <input type="password" name="Password">

                  <span style="font-style:italic"> Enter: 'mynakedpassword'</span><br>
                  <input type="submit" name="Login" value="Let me in!">
          {{end}}
  </form>
  </body></html>`

 var dashboardTemplate = template.Must(template.New("").Parse(dashBoardPage))
 var logUserTemplate = template.Must(template.New("").Parse(logUserPage))

 func DashBoardPageHandler(w http.ResponseWriter, r *http.Request) {
         conditionsMap := map[string]interface{}{}
         //read from session
         session, err := loggedUserSession.Get(r, "authenticated-user-session")

         if err != nil {
                 log.Println("Unable to retrieve session data!", err)
         }

         log.Println("Session name : ", session.Name())

         log.Println("Username : ", session.Values["username"])

         conditionsMap["Username"] = session.Values["username"]

         if err := dashboardTemplate.Execute(w, conditionsMap); err != nil {
                 log.Println(err)
         }
 }

 func LoginPageHandler(w http.ResponseWriter, r *http.Request) {

         conditionsMap := map[string]interface{}{}

         // check if session is active
         session, _ := loggedUserSession.Get(r, "authenticated-user-session")

         if session != nil {
                 conditionsMap["Username"] = session.Values["username"]
         }

         // verify username and password
         if r.FormValue("Login") != "" && r.FormValue("Username") != "" {
                 username := r.FormValue("Username")
                 password := r.FormValue("Password")

                 // NOTE: here is where you want to query your database to retrieve the hashed password
                 // for username.
                 // For this tutorial and simplicity sake, we will simulate the retrieved hashed password
                 // as $2a$10$4Yhs5bfGgp4vz7j6ScujKuhpRTA4l4OWg7oSukRbyRN7dc.C1pamu
                 // the plain password is 'mynakedpassword'
                 // see https://www.socketloop.com/tutorials/golang-bcrypting-password for more details
                 // on how to generate bcrypted password

                 hashedPasswordFromDatabase := []byte("$2a$10$4Yhs5bfGgp4vz7j6ScujKuhpRTA4l4OWg7oSukRbyRN7dc.C1pamu")

                 if err := bcrypt.CompareHashAndPassword(hashedPasswordFromDatabase, []byte(password)); err != nil {
                         log.Println("Either username or password is wrong")
                         conditionsMap["LoginError"] = true
                 } else {
                         log.Println("Logged in :", username)
                         conditionsMap["Username"] = username
                         conditionsMap["LoginError"] = false

                         // create a new session and redirect to dashboard
                         session, _ := loggedUserSession.New(r, "authenticated-user-session")

                         session.Values["username"] = username
                         err := session.Save(r, w)

                         if err != nil {
                                 log.Println(err)
                         }

                         http.Redirect(w, r, "/dashboard", http.StatusFound)
                 }

         }

         if err := logUserTemplate.Execute(w, conditionsMap); err != nil {
                 log.Println(err)
         }
 }

 func LogoutHandler(w http.ResponseWriter, r *http.Request) {
         //read from session
         session, _ := loggedUserSession.Get(r, "authenticated-user-session")

         // remove the username
         session.Values["username"] = ""
         err := session.Save(r, w)

         if err != nil {
                 log.Println(err)
         }

         w.Write([]byte("Logged out!"))
 }

 func main() {
         fmt.Println("Server starting, point your browser to localhost:8080/login to start")
         http.HandleFunc("/login", LoginPageHandler)
         http.HandleFunc("/dashboard", DashBoardPageHandler)
         http.HandleFunc("/logout", LogoutHandler)

         //http.ListenAndServe(":8080", nil)
	 //using port in heroku
	 port := os.Getenv("PORT") 
	 http.ListenAndServe(":" + port, nil)
 }
trermina version gorrila bencrypt

*/

//  original 
import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.Run(":" + port)
}
//termin ortiginall


/*
version with gorrilla  mux 


import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"net/http"
)

// cookie handling

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func getUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

// login handler

func loginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	pass := request.FormValue("password")
	redirectTarget := "/"
	if name != "" && pass != "" {
		// .. check credentials ..
		setSession(name, response)
		redirectTarget = "/internal"
	}
	http.Redirect(response, request, redirectTarget, 302)
}

// logout handler

func logoutHandler(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/", 302)
}

// index page

const indexPage = `
<h1>Login</h1>
<form method="post" action="/login">
    <label for="name">User name</label>
    <input type="text" id="name" name="name">
    <label for="password">Password</label>
    <input type="password" id="password" name="password">
    <button type="submit">Login</button>
</form>
`

func indexPageHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, indexPage)
}

// internal page

const internalPage = `
<h1>Internal</h1>
<hr>
<small>User: %s</small>
<form method="post" action="/logout">
    <button type="submit">Logout</button>
</form>
`

func internalPageHandler(response http.ResponseWriter, request *http.Request) {
	userName := getUserName(request)
	if userName != "" {
		fmt.Fprintf(response, internalPage, userName)
	} else {
		http.Redirect(response, request, "/", 302)
	}
}

// server main method

var router = mux.NewRouter()

func main() {

	router.HandleFunc("/", indexPageHandler)
	router.HandleFunc("/internal", internalPageHandler)

	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/logout", logoutHandler).Methods("POST")

	http.Handle("/", router)

         //http.ListenAndServe(":8080", nil)
	 //using port in heroku
	 port := os.Getenv("PORT") 
	 http.ListenAndServe(":" + port, nil)
}


*/
