// ProjectName:: IIITK-VizB-Services ver1.0
//
// ProjectTitle:: Vizualization services for big datasets
//
// ProjectDescription:: See doc folder for more information.
//
// Authors:: i) ETHAKOTA PAVANSAI
//          ii) Shajulin Benedict
//          iii) VANKUDOTH PREMKUMAR
//          iv) SUMALA SAI SIDDARDHA


package bigdata

import (
  "fmt"
  "net/http"
  "github.com/senseyeio/roger"
)
func Bigdatacompscatter(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgCompScatter()")
  if err != nil {
    fmt.Fprintf(w, "Graph generation failed with error %s", err.Error())
    return
  }

  // convert response to a byte array
  imageBytes, ok := returnVar.([]byte)
  if !ok {
    fmt.Fprint(w, "Unexpected response from R")
    return
  }

  // return the binary image data along with suitable headers
  w.Header().Set("Content-Type", "image/png")
  w.Header().Set("Content-Length", strconv.Itoa(len(imageBytes)))
  if _, err := w.Write(imageBytes); err != nil {
    fmt.Fprint(w, "Failed to write image to response")
  }
}
/// bigdata service - bigdata service for IIITK-VizB-Services
