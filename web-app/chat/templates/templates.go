package templates

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type TemplateHandler struct {
	once     sync.Once
	Filename string
	templ    *template.Template
}

func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("web-app/chat/templates", t.Filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		if userInfoBytes, err := base64.StdEncoding.DecodeString(authCookie.Value); err == nil {
			// userDataの形式
			//{
			//	"id": "1222222223445",
			//	"email": "example@gmail.com",
			//	"verified_email": true,
			//	"name": "example",
			//	"given_name": "example",
			//	"picture": "https://asdfj;alskjdf;kajsd;flkad",
			//	"locale": "ja"
			//}
			userData := make(map[string]interface{})
			if err := json.Unmarshal(userInfoBytes, &userData); err != nil {
				log.Printf("failed to unmarshal: %s", err.Error())
			}
			data["UserData"] = userData
		}
	}
	if err := t.templ.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}
