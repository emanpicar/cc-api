package routes

import (
	"encoding/json"
	"net/http"

	"github.com/emanpicar/cc-api/auth"
	"github.com/emanpicar/cc-api/logger"

	"github.com/emanpicar/cc-api/card"
	"github.com/gorilla/mux"
)

type (
	routeHandler struct {
		cardManager card.Manager
		authManager auth.Manager
		router      *mux.Router
	}

	Router interface {
		ServeHTTP(http.ResponseWriter, *http.Request)
	}

	JsonMessage struct {
		Message string `json:"message"`
	}
)

func New(cardManager card.Manager, authManager auth.Manager) Router {
	rh := &routeHandler{
		cardManager: cardManager,
		authManager: authManager,
	}

	return rh.newRouter()
}

func (rh *routeHandler) newRouter() *mux.Router {
	router := mux.NewRouter()
	rh.registerRoutes(router)

	return router
}

func (rh *routeHandler) registerRoutes(router *mux.Router) {
	router.HandleFunc("/api/authenticate", rh.authenticate).Methods("POST")
	router.HandleFunc("/api/validateCards", rh.authMiddleware(rh.validateCards)).Methods("POST")

	rh.router = router
}

func (rh *routeHandler) authenticate(w http.ResponseWriter, r *http.Request) {
	logger.Log.Infoln("Authenticating user")

	w.Header().Set("Content-Type", "application/json")
	data, err := rh.authManager.Authenticate(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		rh.encodeError(json.NewEncoder(w).Encode(&JsonMessage{err.Error()}), w)
		return
	}

	rh.encodeError(json.NewEncoder(w).Encode(data), w)
}

func (rh *routeHandler) validateCards(w http.ResponseWriter, r *http.Request) {
	logger.Log.Infoln("Starting cards validation")

	w.Header().Set("Content-Type", "application/json")
	data, err := rh.cardManager.ValidateCardList(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		rh.encodeError(json.NewEncoder(w).Encode(&JsonMessage{err.Error()}), w)
		return
	}

	rh.encodeError(json.NewEncoder(w).Encode(data), w)
}

func (rh *routeHandler) authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := rh.authManager.ValidateRequest(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			rh.encodeError(json.NewEncoder(w).Encode(&JsonMessage{err.Error()}), w)
			return
		}

		next(w, r)
	})
}

func (rh *routeHandler) encodeError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
