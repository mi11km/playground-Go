package main

import (
	"app/web-app/chat"
	"app/web-app/chat/trace"
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("web-app/chat/templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	addr := flag.String("addr", ":8080", "appのアドレス")
	flag.Parse()

	r := chat.NewRoom()
	r.Tracer = trace.New(os.Stdout)

	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	go r.Run()

	log.Println("webサーバーを開始。ポート：", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
