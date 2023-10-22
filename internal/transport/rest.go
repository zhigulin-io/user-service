package transport

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type RestServer struct {
	userService UserService
}

type UserService interface {
	RegisterUser(username, password string) error
}

type regAuthReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewRestServer(userService UserService) *RestServer {
	return &RestServer{
		userService: userService,
	}
}

func (serv RestServer) Start() {
	http.HandleFunc("/register", serv.handleRegister)
	http.HandleFunc("/login", serv.handleLogin)
	http.ListenAndServe(":8080", nil)
}

func (serv RestServer) handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	req := regAuthReq{}
	if err := json.Unmarshal(bytes, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = serv.userService.RegisterUser(req.Username, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (serv RestServer) handleLogin(w http.ResponseWriter, r *http.Request) {
	t := jwt.New(jwt.SigningMethodHS256)
	s, err := t.SignedString([]byte("key"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(s))
}
