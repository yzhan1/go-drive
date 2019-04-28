package handler

import (
	"fmt"
	"github.com/yzhan1/go-drive/db"
	"github.com/yzhan1/go-drive/util"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	PWD_SALT = "*#90"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signup.html")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	if len(username) < 3 || len(password) < 5 {
		w.Write([]byte("Invalid parameter"))
		return
	}
	encodedPassword := util.Sha1([]byte(password + PWD_SALT))
	if db.SignUp(username, encodedPassword) {
		w.Write([]byte("SUCCESS"))
	} else {
		w.Write([]byte("Failed"))
	}
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "/static/view/signin.html", http.StatusFound)
		return
	}
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	encodedPassword := util.Sha1([]byte(password + PWD_SALT))
	authenticated := db.SignIn(username, encodedPassword)

	if !authenticated {
		w.Write([]byte("Failed"))
		return
	}

	token := GenerateToken(username)
	if !db.UpdateToken(username, token) {
		w.Write([]byte("Failed"))
	} else {
		resp := util.Response{
			Code: 0,
			Msg:  "OK",
			Data: struct {
				Location string
				Username string
				Token    string
			}{
				Location: "http://" + r.Host + "/static/view/home.html",
				Username: username,
				Token:    token,
			},
		}
		w.Write(resp.ToJSONBytes())
	}
}

func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	user, err := db.GetUserInfo(username)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
	} else {
		resp := util.Response{
			Code: 0,
			Msg:  "OK",
			Data: user,
		}
		w.Write(resp.ToJSONBytes())
	}
}

func GenerateToken(username string) string {
	ts := fmt.Sprintf("%x", time.Now().Unix())
	prefix := util.MD5([]byte(username + ts + "TOKEN_SALT"))
	return prefix + ts[:8]
}

func isValidToken(token string) bool {
	if len(token) != 40 {
		return false
	}
	return true
}
