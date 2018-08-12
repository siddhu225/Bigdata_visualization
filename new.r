install.packages("corrplot")
install.packages("Rserve")
library("corrplot")
library("Rserve")
library(ggplot2)
library(hexbin)

generateCorrelationPlot <- function () {
  # create temporary png file
  filename <- tempfile("plot", fileext = ".png")
  png(filename)

  # draw graph to png image
  M <- cor(mtcars)
  corrplot(M)
  dev.off()

  # obtain image binary array
  image <- readBin(filename, "raw", .Machine$integer.max)
  unlink(filename)

  # return binary array to Go
  return (image)
}
bgCompScatter <- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)

  rawData <- read.csv(file="C:/Users/sumala/Desktop/PROJECT/consolidatedfile-new.csv", header=T)
  plot(rawData$Node,rawData$Exclusive.PAPI_NATIVE_rapl..RAPL_ENERGY_PKG.cpu.0)
  dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}
run.Rserve()
