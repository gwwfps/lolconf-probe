package main

import (
  "bufio"
  "encoding/json"
  "errors"
  "fmt"
  "github.com/gwwfps/lolconf-probe/display"
  "os"
  "strings"
)

type dispatcher struct {
  handlers map[string]inner
}
type inner func() (interface{}, error)
type HanderError struct {
  Message string `json:errorMsg`
}

func (d *dispatcher) dispatch(command string) string {
  inf := d.handlers[command]
  if inf == nil {
    return errorToString(errors.New("Invalid command: " + command))
  }

  result, innerError := inf()
  if innerError != nil {
    return errorToString(innerError)
  }

  serialized, marshalError := json.Marshal(result)
  if marshalError != nil {
    return errorToString(marshalError)
  }

  return string(serialized)
}

func (d *dispatcher) register(command string, f inner) {
  if d.handlers == nil {
    d.handlers = map[string]inner{}
  }
  d.handlers[command] = f
}

func errorToString(err error) string {
  serialized, marshalError := json.Marshal(HanderError{err.Error()})
  if marshalError != nil {
    return "{\"Message\":\"" + err.Error() + " An additional error occurred during JSON serialization.\"}"
  }
  return string(serialized)
}

func main() {
  d := new(dispatcher)
  d.register("resolutions", display.ListAvailableResolutions)

  reader := bufio.NewReader(os.Stdin)

  var line string
  var err error
  for ; err == nil; line, err = reader.ReadString('\n') {
    command := strings.Trim(line, "\r\n ")
    if command != "" {
      fmt.Println(d.dispatch(command))
    }
  }
}
