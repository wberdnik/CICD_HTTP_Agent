package services

/**
 * Package  httpException

 * @author  Morph
 * @version 0.1
 */

import (
	"encoding/json"
	"net/http"
)

// respondCode - prepare response
func respondCode(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		switch t := data.(type) {
		case string:
			w.Write([]byte(t))
		case map[string]string:
			json.NewEncoder(w).Encode(t)
		}
	}
	//w.Header().Set("Connection", "close")
	//	w.finishRequest()
}

func Http400Exception(w http.ResponseWriter, err error) { // общая ошибка
	respondCode(w, http.StatusBadRequest, map[string]string{"httpException": err.Error()})
}

func Http401Exception(w http.ResponseWriter) {
	w.Header().Add("WWW-Authenticate", `Bearer realm="Need refresh token"`)
	respondCode(w, http.StatusUnauthorized, nil)
}

func Http404Exception(w http.ResponseWriter, r *http.Request) {
	http.NotFoundHandler().ServeHTTP(w, r)
}

func Http411ExceptionNeedLenght(w http.ResponseWriter) { //Для тех кто с телом
	respondCode(w, http.StatusLengthRequired, "Content-Length required.")
}

func Http413ExceptionTooLarge(w http.ResponseWriter) {
	respondCode(w, http.StatusRequestEntityTooLarge, "Payload Too Large. Limit of entity is 500000 bytes.")
}
func Http415ExceptionFormat(w http.ResponseWriter) {
	respondCode(w, http.StatusUnsupportedMediaType, "The requested response format unspecified. It is expected: application/json.")
}
func Http422ExceptionBisness(w http.ResponseWriter, reason string) {
	respondCode(w, http.StatusUnprocessableEntity, reason)
}
