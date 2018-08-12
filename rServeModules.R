library("Rserve")
library(ggplot2)
library(hexbin)
library(ggthemes)
library(tabplot)
library(ggcorrplot)
library(ggExtra)
library(ggalt)
library(hexbin)
library(RColorBrewer)
library(magrittr)
library(leaflet)


rawData <- read.csv(file="C:/Users/sumala/Desktop/PROJECT/consolidatedfile-new.csv", header=TRUE,sep="\t")

bgCompScatter <- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)


  plot(rawData$Node,rawData$Exclusive.PAPI_NATIVE_rapl..RAPL_ENERGY_PKG.cpu.0)
  dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}

bgDistBox <- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)

x=rnorm(100)
y=rnorm(100,5,1)
g <- ggplot(rawData, aes(rawData$Node,rawData$Exclusive.PAPI_NATIVE_rapl..RAPL_ENERGY_PKG.cpu.0))
plot(g + geom_boxplot(varwidth=T, fill="plum") +labs(title="Box plot"))
dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}

bgCompArea <- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)

x=rnorm(100)
y=rnorm(100,5,1)
plot(ggplot(rawData, aes(Node,NumSubr)) + geom_area( fill="red", alpha=.2)+ geom_line())
dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}

bgDistDotBox <- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)

x=rnorm(100)
y=rnorm(100,5,1)
plot(ggplot(rawData,aes(Node,Exclusive.PAPI_NATIVE_rapl..RAPL_ENERGY_PKG.cpu.0))+ geom_boxplot() + geom_dotplot(binaxis='y', stackdir='center', dotsize = .5, fill="red") +theme(axis.text.x = element_text(angle=65, vjust=0.6)) + labs(title="Box plot + Dot plot",x="Node",y="Exclusive.PAPI_NATIVE_rapl..RAPL_ENERGY_PKG.cpu.0"))
dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}

bgDistKernalDensity <- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)

plot(density(rawData$Node))
dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}

bgDistNotchedBox<- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)

p<-ggplot(rawData, aes(Node, Exclusive.PAPI_NATIVE_rapl..RAPL_ENERGY_PKG.cpu.0))
plot(p + geom_boxplot())
dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}

bgDistTufteBox<- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)

g <- ggplot(rawData, aes(rawData$Node,rawData$Exclusive.PAPI_NATIVE_rapl..RAPL_ENERGY_PKG.cpu.0))
plot(g + geom_tufteboxplot() + theme(axis.text.x = element_text(angle=65, vjust=0.6)) + labs(title="Tufte Styled Boxplot"))
dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}


bgDistViolin<- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)

plot(ggplot(rawData,aes(rawData$Node,rawData$Exclusive.PAPI_NATIVE_rapl..RAPL_ENERGY_PKG.cpu.0))+geom_violin())
dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}


bgRelBubble<- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)


set.seed(200)

x_var <- rnorm( n = 15, mean = 5, sd = 2)
y_var <- x_var + rnorm(n = 15, mean = 5, sd =4)
size_var <- runif(15, 1,10)

df.test_data <- data.frame(x_var, y_var, size_var)

# PLOT THE DATA USING GGPLOT2
plot(ggplot(data=df.test_data, aes(x=x_var, y=y_var)) +
  geom_point(aes(size=size_var)) +
  scale_size_continuous(range=c(2,15)) +
  theme(legend.position = "none"))
dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}

bgRelTable<- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)

is.na(rawData$Node) <- rawData$Thread == "Ideal"
is.na(rawData$Thread) <- (runif(nrow(rawData)) > 1)
plot(tableplot(rawData))
dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}

bgRelCount<- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)

plot(ggplot(rawData, aes(rawData$Node,rawData$Exclusive.PAPI_NATIVE_rapl..RAPL_ENERGY_PKG.cpu.0)) + geom_count() + labs(title="Scatterplot Count"))
dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}

bgRelJitter<- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)

plot(ggplot(rawData, aes(rawData$Node,rawData$Exclusive.PAPI_NATIVE_rapl..RAPL_ENERGY_PKG.cpu.0)) + geom_jitter() + labs(title="Scatterplot(Jitter)"))
dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}

bgRelCorrelogram<- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)

data(mtcars)
corr <- round(cor(mtcars), 1)

# Plot
plot(ggcorrplot(corr, hc.order = TRUE,
           type = "lower",
           lab = TRUE,
           lab_size = 3,
           method="circle",
           colors = c("tomato2", "white", "springgreen3"),
           title="Correlogram of mtcars",
           ggtheme=theme_bw))
dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}


bgRelMarginalHistogram<- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)

data_select <- rawData[rawData$Node>=20 & rawData$Thread<=10 ]
g<-ggplot(rawData, aes(Node,Thread))+ geom_count()+ labs(y="Thread",x="Node",title="marginal histogrm and box plot")
plot(ggMarginal(g,type="histogram",fill="transparent"))
plot(ggMarginal(g,type="boxplot",fill="transparent"))
dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}

bgRelScatterEncircling<- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)

plot(ggplot(rawData, aes(rawData$Node,rawData$NumSubr)) + geom_encircle() + labs(title="Scatterplot"))
dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}

bgCompPie<- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)

pie <- ggplot(rawData, aes(x = "", fill = factor(Node))) +  geom_bar(width = 1) +theme(axis.line = element_blank(), plot.title = element_text(hjust=0.5)) +
labs(fill="class", x=NULL, y=NULL, title="Pie Chart of class")
plot(pie + coord_polar(theta = "y", start=0))
dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}

bgCompStackedArea<- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)

set.seed(345)
Sector <- rep(c("S01","S02","S03","S04","S05","S06","S07"),times=7)
Year <- as.numeric(rep(c("1950","1960","1970","1980","1990","2000","2010"),each=7))
Value <- runif(49, 10, 100)
data <- data.frame(Sector,Year,Value)
plot(ggplot(data, aes(x=Year, y=Value, fill=Sector)) +
    geom_area())
dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}

bgCompWaffle<- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)

var<-rawData$Node
nrows <- 10
df <- expand.grid(y = 1:nrows, x = 1:nrows)
categ_table <- round(table(var) * ((nrows*nrows)/(length(var))))
categ_table
df$category <- factor(rep(names(categ_table), categ_table))
plot(ggplot(df, aes(x = x, y = y, fill = category)) + geom_tile(color = "black", size = 0.5) +
      scale_x_continuous(expand = c(0, 0)) +
      scale_y_continuous(expand = c(0, 0), trans = 'reverse') +
      scale_fill_brewer(palette = "Set3") +
      labs(title="Waffle Chart"))
dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}

bgCompCorrelation <- function () {
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

bgCmpHistogram<- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)

g <- ggplot(rawData, aes(rawData$Node)) + scale_fill_brewer(palette = "Spectral")
plot(g + geom_histogram(aes(fill=Thread),
                      bins=5,
                      col="black",
                      size=.1))
dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}

bgCmpBDBar<- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)

plot(ggplot(rawData, aes(x = Thread, y = Node ,fill = Node)) + geom_bar(width = 0.85, stat="identity"))
dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}



bgCmpHeat<- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)

data=as.matrix(mtcars)
head(data)
heatmap(data)
dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}

bgCmpHexbin<- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)


a=hexbin(rawData$Node,diamonds$NumSubr,xbins=40)
plot(a)
dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}

bgCmpMosaic<- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)


data(HairEyeColor)
mosaicplot(HairEyeColor)
dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}

bgCmpMap<- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)


m <- leaflet() %>%
addTiles() %>%  # Add default OpenStreetMap map tiles
addMarkers(lng=77.2310, lat=28.6560, popup="The delicious food of chandni chowk")
m
dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}

bgRelLollipop<- function() {

# create temporary png file
scatter <- tempfile("plot", fileext = ".png")
png(scatter)

plot(ggplot(rawData, aes(x=Node, y=NumSubr)) +
  geom_point(size=3) +
  geom_segment(aes(x=Node,
                   xend=Node,
                   y=0,
                   yend=NumCalls)))
dev.off()

  # obtain image binary array
  image <- readBin(scatter, "raw", .Machine$integer.max)
  unlink(scatter)

  # return binary array to Go
  return (image)
}


#Start the server now
run.Rserve()
