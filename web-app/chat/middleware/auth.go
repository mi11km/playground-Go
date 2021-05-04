package middleware

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type authHandler struct {
	next http.Handler
}

func (a *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("auth"); err == http.ErrNoCookie {
		// 未認証
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	} else if err != nil {
		// 何らかの別のエラー発生
		panic(err.Error())
	} else {
		a.next.ServeHTTP(w, r)
	}
}

func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

//--------------

const oauthStateString = "aaefa3234AEAFweafw9r923favn[@;:"

var (
	oauthConfigs      = make(map[string]*oauth2.Config)
	userInfoEndpoints = make(map[string]string)
)

func init() {
	oauthConfigs["google"] = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/callback/google",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	userInfoEndpoints["google"] = "https://www.googleapis.com/oauth2/v2/userinfo"
}

// LoginHandler サードパーティーへのログイン処理の受け持ち
// path: /auth/{action}/{provider}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	if len(segs) < 4 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "url: %sには非対応です", r.URL.Path)
	} else {
		action := segs[2]
		provider := segs[3]
		switch action {
		case "login":
			url := oauthConfigs[provider].AuthCodeURL(oauthStateString)
			http.Redirect(w, r, url, http.StatusTemporaryRedirect)
		case "callback":
			contents, err := getUserInfo(provider, r.FormValue("state"), r.FormValue("code"))
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			} else {
				// todo 
				fmt.Fprintf(w, "Content: %s\n", contents)
			}
		default:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "アクション%sには非対応です", action)
		}
	}
}

func getUserInfo(provider, state, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := oauthConfigs[provider].Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, userInfoEndpoints[provider], nil)
	if err != nil {
		return nil, fmt.Errorf("generating new requeset failed: %s", err.Error())
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("new requeset failed: %s", err.Error())
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body failed: %s", err.Error())
	}
	return contents, nil
}
