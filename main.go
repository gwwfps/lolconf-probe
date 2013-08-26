package main

import (
  "encoding/json"
  "fmt"
  "github.com/gwwfps/lolconf-probe/display"
  "log"
  "net/http"
)

type handler func(http.ResponseWriter, *http.Request)
type inner func() (interface{}, error)

func wrapHandler(h inner) handler {
  return func(w http.ResponseWriter, r *http.Request) {
    result, handlerError := h()
    if handlerError != nil {
      writeError(w, handlerError)
      return
    }

    serialized, marshalError := json.Marshal(result)
    if marshalError != nil {
      writeError(w, marshalError)
      return
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    fmt.Fprint(w, string(serialized))
  }
}

func writeError(w http.ResponseWriter, e error) {
  w.WriteHeader(http.StatusInternalServerError)
  fmt.Fprint(w, e.Error())
}

func main() {
  http.HandleFunc("/resolutions", wrapHandler(display.ListAvailableResolutions))
  e := http.ListenAndServe("127.0.0.1:5532", nil)
  if e != nil {
    log.Fatal("ListenAndServe: ", e)
  }
}
