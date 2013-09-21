package main

import (
  "bufio"
  "encoding/json"
  "errors"
  "fmt"
  "github.com/gwwfps/lolconf-probe/display"
  "log"
  "os"
)

type dispatcher struct {
  handlers map[string]inner
}
type inner func() (interface{}, error)
type HanderError struct {
  Message string `json:"errorMsg"`
}
type Query struct {
  Command string `json:"command"`
  SeqNo   string `json:"seqNo"`
}
type ResultWrapper struct {
  Result interface{} `json:"result"`
  SeqNo  string      `json:"seqNo"`
}

func (d *dispatcher) dispatch(line string) string {
  query := new(Query)
  unmarshalError := json.Unmarshal([]byte(line), query)
  if unmarshalError != nil {
    return errorToString(unmarshalError)
  }

  inf := d.handlers[query.Command]
  if inf == nil {
    return errorToString(errors.New("Invalid command: " + query.Command))
  }

  result, innerError := inf()
  if innerError != nil {
    return errorToString(innerError)
  }

  resultWrapper := ResultWrapper{result, query.SeqNo}

  serialized, marshalError := json.Marshal(resultWrapper)
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

  logFile, _ := os.Create("probe.log")
  defer logFile.Close()
  log.SetOutput(logFile)

  log.Println("Started probe.")

  var line string
  var err error
  for ; err == nil; line, err = reader.ReadString('\n') {
    log.Println("Received query:", line)
    if line != "" {
      result := d.dispatch(line)
      log.Println("Returning:", result)
      fmt.Println(result)
    }
  }
}
