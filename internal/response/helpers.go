package response

import (
	"io"
	"net/http"
)

func InternalError(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func Ok(w http.ResponseWriter, result string, contentType string) {
	w.Header().Add("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, result)
}
