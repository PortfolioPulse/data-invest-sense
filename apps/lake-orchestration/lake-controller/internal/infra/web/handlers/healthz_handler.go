package handlers

import (
     "fmt"
     "net/http"
     "time"
)

type WebHealthzHandler struct {
     startedAt time.Time
}

func NewWebHealthzHandler() *WebHealthzHandler {
     return &WebHealthzHandler{
          startedAt: time.Now(),
     }
}

func (h *WebHealthzHandler) Healthz(w http.ResponseWriter, r *http.Request) {
     duration := time.Since(h.startedAt)
     if duration.Seconds() < 30 {
          w.WriteHeader(http.StatusInternalServerError)
          w.Write([]byte(fmt.Sprintf("Healthz check failed after %v seconds", duration.Seconds())))
     } else {
          w.WriteHeader(http.StatusOK)
          w.Write([]byte("Healthz check passed"))
     }
}
