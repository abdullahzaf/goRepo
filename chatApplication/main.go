package main
import (
	"log"
	"net/http"
	"flag"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
)
// temp1 represent a single template
type templateHandler struct {
	once		sync.once
	filename	string
	temp1		*template.Template
}
// ServeHTTP handles the HTTP request
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.temp1 = template.Must(template.ParseFiles(filespath.Join("templates", t.filename)))
}}
t.temp1.Execute(w, r)
}
func main() {

	// root
	http.Handle("/chat", &templateHandler{filename: "chat.html"})
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	//start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	r := newRoom()
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFUnc("/auth/", loginHandler)
	http.Handle("/room", r)
	//get the room going
	go r.run()
	//start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse() //parse the flags
	//setup gomniauth
	gomniuath.SetSecurityKey("PUT YOUR AUTH KEY HERE")
	gomniauth.WithProviders(
		facebook.New("key", "secret",
			"http://localhost:8080/auth/callback/facebook"),
		github.New("key", "secret",
			"http://localhost:8080/auth/callback/github"),
		google.New("key", "secret",
			"http://localhost:8080/auth/callback/google"),
	)
	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFUnc("/auth/", loginHandler)
	http.Handle("/room", r)
	//get the room going
	go r.run()
	//start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
				<head>
					<title>Chat</title>
				</head>
				<body>
					Let's chat!
				</body>
			</html>
		))
	})
	// start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
