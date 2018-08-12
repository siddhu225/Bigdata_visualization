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
  "strconv"
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



func Bigdatadistbox(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgDistBox()")
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


func Bigdatacomparea(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgCompArea()")
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


func Bigdatadistdotbox(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgDistDotBox()")
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

func Bigdatadistkernaldensity(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgDistKernalDensity()")
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

func Bigdatadistnotchedbox(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgDistNotchedBox()")
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

func Bigdatadisttuftebox(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgDistTufteBox()")
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

func Bigdatadistviolin(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgDistViolin()")
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


func Bigdatarelbubble(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgRelBubble()")
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



func Bigdatareltable(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgRelTable()")
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


func Bigdatarelcount(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgRelCount()")
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

func Bigdatareljitter(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgRelJitter()")
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

func Bigdatarelcorrelogram(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgRelCorrelogram()")
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


func Bigdatarelmarginalhistogram(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgRelMarginalHistogram()")
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


func Bigdatarelscatterencircling(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgRelScatterEncircling()")
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

func Bigdatacomppie(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgCompPie()")
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


func Bigdatacompstackedarea(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgCompStackedArea()")
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

func Bigdatacompwaffle(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgCompWaffle()")
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

func Bigdatacmphistogram(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgCmpHistogram()")
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

func Bigdatacmpbdbar(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgCmpBDBar()")
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

func Bigdatacmpheat(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgCmpHeat()")
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


func Bigdatacmphexbin(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgCmpHexbin()")
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


func Bigdatacmpmosaic(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgCmpMosaic()")
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


func Bigdatacmpmap(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgCmpMap()")
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

func Bigdatacompcorrelation(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgCmpHexbin()")
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

func Bigdatarellollipop(w http.ResponseWriter, r *http.Request) {

  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
  }

  // call generateCorrelationPlot R function, gathering the response
  returnVar, err := rClient.Eval("bgRelLollipop()")
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
