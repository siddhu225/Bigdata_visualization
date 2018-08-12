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

package main

import (
  _"encoding/json"
  "log"
  "fmt"
  _"os"
  _"strconv"
  _"io/ioutil"
  "net/http"
  _"gopkg.in/mgo.v2"
  _"gopkg.in/mgo.v2/bson"
  "github.com/gorilla/mux"
  "github.com/siddhu225/data"
)


/// Main info here.
var mainInfo = `
<!DOCTYPE html>
<html>
<title>IIITK-VizB-Services</title>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="https://www.w3schools.com/w3css/4/w3.css">
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Lato">
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Montserrat">
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
<style>
body,h1,h2,h3,h4,h5,h6 {font-family: "Lato", sans-serif}
.w3-bar,h1,button {font-family: "Montserrat", sans-serif}
.fa-anchor,.fa-coffee {font-size:200px}
img {
  padding: 10x 0px;
    width:100%;
}
.navbar {
    overflow: hidden;
    background-color: #e3b048;;
    font-family: Arial, Helvetica, sans-serif;
}

.navbar a {
    float: left;
    font-size: 16px;
    color: white;
    text-align: center;
    padding: 14px 16px;
    text-decoration: none;
}

.dropdown {
    float: left;
    overflow: hidden;
}

.dropdown .dropbtn {
    font-size: 16px;
    border: none;
    outline: none;
    color: white;
    padding: 14px 16px;
    background-color: inherit;
    font-family: inherit;
    margin: 0;
}

.navbar a:hover, .dropdown:hover .dropbtn {
    background-color: red;
}

.dropdown-content {
    display: none;
    position: absolute;
    background-color: #f9f9f9;
    min-width: 160px;
    box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
    z-index: 1;
}

.dropdown-content a {
    float: none;
    color: black;
    padding: 12px 16px;
    text-decoration: none;
    display: block;
    text-align: left;
}

.dropdown-content a:hover {
    background-color: #ddd;
}

.dropdown:hover .dropdown-content {
    display: block;
}
</style>
<body>

<!-- Navbar -->
<div class="navbar">
  <a href="sa.html">Home</a>
  <a href="contact.html">contact</a>
  <a href="about.html">about</a>
  <div class="dropdown">
    <button class="dropbtn">Services
      <i class="fa fa-caret-down"></i>
    </button>
    <div class="dropdown-content">
      <a href="distribution.html">distributioncharts</a>
      <a href="comparison.html">comparisoncharts</a>
      <a href="relationship.html">Relationcharts</a>
      <a href="composition.html">Compositioncharts</a>
    </div>
  </div>
</div>




<!-- Header -->
<header class="w3-container " style="padding:0px 0px">
   <img src="p2.jpg"  width="100%" height="500">

</header>



<script>
// Used to toggle the menu on small screens when clicking on the menu button
function myFunction() {
    var x = document.getElementById("navDemo");
    if (x.className.indexOf("w3-show") == -1) {
        x.className += " w3-show";
    } else {
        x.className = x.className.replace(" w3-show", "");
    }
}
</script>

</body>
</html>

`
var contact=`<!DOCTYPE html>
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1">
<style>
body {font-family: Arial, Helvetica, sans-serif;}
<style>
body,h1,h2,h3,h4,h5,h6 {font-family: "Lato", sans-serif}
.w3-bar,h1,button {font-family: "Montserrat", sans-serif}
.fa-anchor,.fa-coffee {font-size:200px}
img {
  padding: 10x 0px;
    width:100%;
}
.navbar {
    overflow: hidden;
    background-color: #e3b048;;
    font-family: Arial, Helvetica, sans-serif;
}

.navbar a {
    float: left;
    font-size: 16px;
    color: white;
    text-align: center;
    padding: 14px 16px;
    text-decoration: none;
}

.dropdown {
    float: left;
    overflow: hidden;
}

.dropdown .dropbtn {
    font-size: 16px;
    border: none;
    outline: none;
    color: white;
    padding: 14px 16px;
    background-color: inherit;
    font-family: inherit;
    margin: 0;
}

.navbar a:hover, .dropdown:hover .dropbtn {
    background-color: red;
}

.dropdown-content {
    display: none;
    position: absolute;
    background-color: #f9f9f9;
    min-width: 160px;
    box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
    z-index: 1;
}

.dropdown-content a {
    float: none;
    color: black;
    padding: 12px 16px;
    text-decoration: none;
    display: block;
    text-align: left;
}

.dropdown-content a:hover {
    background-color: #ddd;
}

.dropdown:hover .dropdown-content {
    display: block;
}


input[type=text], select, textarea {
    width: 100%;
    padding: 12px;
    border: 1px solid #ccc;
    border-radius: 4px;
    box-sizing: border-box;
    margin-top: 6px;
    margin-bottom: 16px;
    resize: vertical;
}

input[type=submit] {
    background-color: #4559a0;
    color: white;
    padding: 12px 20px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}

input[type=submit]:hover {
    background-color: #45a049;
}

.container {
    border-radius: 5px;
    background-color: #f2f2f2;
    padding: 20px;
}
</style>
</head>
<body>
  <div class="navbar">
    <a href="sa.html">Home</a>
    <a href="contact.html">contact</a>
    <a href="about.html">about</a>
    <div class="dropdown">
      <button class="dropbtn">Services
        <i class="fa fa-caret-down"></i>
      </button>
      <div class="dropdown-content">
        <a href="distribution.html">distributioncharts</a>
        <a href="comparison.html">comparisoncharts</a>
        <a href="relationship.html">Relationcharts</a>
        <a href="composition.html">Compositioncharts</a>
      </div>
    </div>
  </div>

<h3>Contact Form</h3>

<div class="container">
  <form action="/action_page.php">
    <label for="fname">First Name</label>
    <input type="text" id="fname" name="firstname" placeholder="Your name..">

    <label for="lname">Last Name</label>
    <input type="text" id="lname" name="lastname" placeholder="Your last name..">

    <label for="country">Country</label>
    <select id="country" name="country">
      <option value="australia">Australia</option>
      <option value="canada">Canada</option>
      <option value="usa">USA</option>
    </select>

    <label for="subject">Subject</label>
    <textarea id="subject" name="subject" placeholder="Write something.." style="height:200px"></textarea>

    <input type="submit" value="Submit">
  </form>
</div>

</body>
</html>

`

var about=`
<!DOCTYPE html>
<html>
<title>IIITK-VizB-Services</title>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="https://www.w3schools.com/w3css/4/w3.css">
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Lato">
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Montserrat">
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
<style>
body,h1,h2,h3,h4,h5,h6 {font-family: "Lato", sans-serif}
.w3-bar,h1,button {font-family: "Montserrat", sans-serif}
.fa-anchor,.fa-coffee {font-size:200px}
img {
  padding: 10x 0px;
    width:100%;
}
.navbar {
    overflow: hidden;
    background-color: #e3b048;;
    font-family: Arial, Helvetica, sans-serif;
}

.navbar a {
    float: left;
    font-size: 16px;
    color: white;
    text-align: center;
    padding: 14px 16px;
    text-decoration: none;
}

.dropdown {
    float: left;
    overflow: hidden;
}

.dropdown .dropbtn {
    font-size: 16px;
    border: none;
    outline: none;
    color: white;
    padding: 14px 16px;
    background-color: inherit;
    font-family: inherit;
    margin: 0;
}

.navbar a:hover, .dropdown:hover .dropbtn {
    background-color: red;
}

.dropdown-content {
    display: none;
    position: absolute;
    background-color: #f9f9f9;
    min-width: 160px;
    box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
    z-index: 1;
}

.dropdown-content a {
    float: none;
    color: black;
    padding: 12px 16px;
    text-decoration: none;
    display: block;
    text-align: left;
}

.dropdown-content a:hover {
    background-color: #ddd;
}

.dropdown:hover .dropdown-content {
    display: block;
}
</style>
<body>

<!-- Navbar -->
<div class="navbar">
  <a href="sa.html">Home</a>
  <a href="contact.html">contact</a>
  <a href="about.html">about</a>
  <div class="dropdown">
    <button class="dropbtn">Services
      <i class="fa fa-caret-down"></i>
    </button>
    <div class="dropdown-content">
      <a href="distribution.html">distributioncharts</a>
      <a href="comparison.html">comparisoncharts</a>
      <a href="relationship.html">Relationcharts</a>
      <a href="composition.html">Compositioncharts</a>
    </div>
  </div>
</div>




<!-- Header -->
<header class="w3-container " style="padding:0px 0px">
   <h1>DATA VISUALIZATION</h1>
   <p>Data visualization is a general term that describes any effort to help people understand the significance of data by placing it in a visual context.
      Patterns, trends and correlations that might go undetected in text-based data can be exposed and recognized easier with data visualization software.
</p>
<h2>Why BigData Is important</h2>
<p>
What makes Big Data so useful to many companies is the fact that it provides answers to many questions that they didn’t even know they had in the first place, he said. In other words, it provides a point of reference. With such a massive amount of information, the data is able to be shaped or tested in any way that the company sees fit. In doing so, organizations are able to pinpoint issues in a more comprehensible form. Collecting masses of data and finding a trend within the data allows the businesses to move much more quickly, smoothly, and efficiently. It also allows them to eliminate problem areas before those previously elusive issues pull their profits or reputation through the proverbial mud.
</p>
<h3>Making Good Use Of Data</h3>
<p>
  Understanding the attributes that the data pertains to is key to using Big Data successfully. This is otherwise known as “getting to know your customer,” so that the company can market to their customers’ needs with a much higher success rate.

He gave an example of a case study where his team helped to cleanse the data of one of their clients. They removed any unrelated or outlier data from their data sets and thus narrowed it down to one key question or demographic. In doing so, they were able to identify which of their products sold more, which of their products did not sell, and thus could be eliminated. They looked at four primary data elements: revenue, frequency, value, and tenure. Mr. Abate stressed that anything more than four elements at a time, on any given visualization, become too difficult for people to follow. By eliminating products that did not sell, they were increasing their future revenue by cutting off waste. But they could not have done any of this without utilizing Data Visualization.
</p>
</header>



<script>
// Used to toggle the menu on small screens when clicking on the menu button
function myFunction() {
    var x = document.getElementById("navDemo");
    if (x.className.indexOf("w3-show") == -1) {
        x.className += " w3-show";
    } else {
        x.className = x.className.replace(" w3-show", "");
    }
}
</script>

</body>
</html>

`

var dist=`<!DOCTYPE html>
<html>
<title>IIITK-VizB-Services</title>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="https://www.w3schools.com/w3css/4/w3.css">
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Lato">
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Montserrat">
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
<style>
body,h1,h2,h3,h4,h5,h6 {font-family: "Lato", sans-serif}
.w3-bar,h1,button {font-family: "Montserrat", sans-serif}
.fa-anchor,.fa-coffee {font-size:200px}
.s1{font-size: 15px}

.navbar {
    overflow: hidden;
    background-color: #e3b048;;
    font-family: Arial, Helvetica, sans-serif;
}

.navbar a {
    float: left;
    font-size: 16px;
    color: white;
    text-align: center;
    padding: 14px 16px;
    text-decoration: none;
}

.dropdown {
    float: left;
    overflow: hidden;
}

.dropdown .dropbtn {
    font-size: 16px;
    border: none;
    outline: none;
    color: white;
    padding: 14px 16px;
    background-color: inherit;
    font-family: inherit;
    margin: 0;
}

.navbar a:hover, .dropdown:hover .dropbtn {
    background-color: red;
}

.dropdown-content {
    display: none;
    position: absolute;
    background-color: #f9f9f9;
    min-width: 160px;
    box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
    z-index: 1;
}

.dropdown-content a {
    float: none;
    color: black;
    padding: 12px 16px;
    text-decoration: none;
    display: block;
    text-align: left;
}

.dropdown-content a:hover {
    background-color: #ddd;
}

.dropdown:hover .dropdown-content {
    display: block;
}
</style>
<body>

<!-- Navbar -->
<div class="navbar">
  <a href="sa.html">Home</a>
  <a href="contact.html">contact</a>
  <a href="about.html">about</a>
  <div class="dropdown">
    <button class="dropbtn">Services
      <i class="fa fa-caret-down"></i>
    </button>
    <div class="dropdown-content">
      <a href="distribution.html">distributioncharts</a>
      <a href="comparison.html">comparisoncharts</a>
      <a href="relationship.html">Relationcharts</a>
      <a href="composition.html">Compositioncharts</a>
    </div>
  </div>
</div>




<!-- Header -->
<header class="s1" style="padding:0px 0px">
   <h1 style="font-size:15px">Distribution Charts<h1>
   <p style="font-size:15px">A distribution combines the merits and functionality of both Comparison and Composition. It can help you see the full data spectrum and visualize related or unrelated data points.
     With distribution you can see if there are any correlations, trends, patterns, shapes, clusters, averages, or outliers.


   <ul style="font-size:15px">
      Use a distribution or deviation visualizations chart to:
      <li>Visualize the distribution of ages in a group of people</li>
      <li>Visualize the correlations between cause and effect</li>
      <li>Spot bugs or constraints in production or quality control systems</li>
      <li>Visualize sales in various stores, to spot best and worst performing ones</li>
   </ul>
   <ul type="square" style="font-size:15px">
      Some of the most important distribution charts are:
        <li><a href="/2">BOXCHART </a>  </li>
        <li><a href="/4">DOT+BOX CHART</a>  </li>
        <li><a href="/5">KERNALDENSITY CHART</a>  </li>
        <li><a href="/6">NOTCHEDBOX CHART</a>  </li>
        <li><a href="/7">TUFTEBOX CHART</a>  </li>
        <li><a href="/8">VIOLIMBOX CHART</a>  </li>


    </ul>


</p>


</header>



<script>
// Used to toggle the menu on small screens when clicking on the menu button
function myFunction() {
    var x = document.getElementById("navDemo");
    if (x.className.indexOf("w3-show") == -1) {
        x.className += " w3-show";
    } else {
        x.className = x.className.replace(" w3-show", "");
    }
}
</script>

</body>
</html>
`

var compare=`<!DOCTYPE html>
<html>
<title>IIITK-VizB-Services</title>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="https://www.w3schools.com/w3css/4/w3.css">
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Lato">
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Montserrat">
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
<style>
body,h1,h2,h3,h4,h5,h6 {font-family: "Lato", sans-serif}
.w3-bar,h1,button {font-family: "Montserrat", sans-serif}
.fa-anchor,.fa-coffee {font-size:200px}
.s1{font-size: 15px}

.navbar {
    overflow: hidden;
    background-color: #e3b048;;
    font-family: Arial, Helvetica, sans-serif;
}

.navbar a {
    float: left;
    font-size: 16px;
    color: white;
    text-align: center;
    padding: 14px 16px;
    text-decoration: none;
}

.dropdown {
    float: left;
    overflow: hidden;
}

.dropdown .dropbtn {
    font-size: 16px;
    border: none;
    outline: none;
    color: white;
    padding: 14px 16px;
    background-color: inherit;
    font-family: inherit;
    margin: 0;
}

.navbar a:hover, .dropdown:hover .dropbtn {
    background-color: red;
}

.dropdown-content {
    display: none;
    position: absolute;
    background-color: #f9f9f9;
    min-width: 160px;
    box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
    z-index: 1;
}

.dropdown-content a {
    float: none;
    color: black;
    padding: 12px 16px;
    text-decoration: none;
    display: block;
    text-align: left;
}

.dropdown-content a:hover {
    background-color: #ddd;
}

.dropdown:hover .dropdown-content {
    display: block;
}
</style>
<body>

<!-- Navbar -->
<div class="navbar">
  <a href="sa.html">Home</a>
  <a href="contact.html">contact</a>
  <a href="about.html">about</a>
  <div class="dropdown">
    <button class="dropbtn">Services
      <i class="fa fa-caret-down"></i>
    </button>
    <div class="dropdown-content">
      <a href="distribution.html">distributioncharts</a>
      <a href="comparison.html">comparisoncharts</a>
      <a href="relationship.html">Relationcharts</a>
      <a href="composition.html">Compositioncharts</a>
    </div>
  </div>
</div>




<!-- Header -->
<header class="s1" style="padding:0px 0px">
   <h1 style="font-size:15px">Comparison Charts<h1>
   <p style="font-size:15px">Comparison of data points is probably the most common and easy-to-understand method for data analysis.
      As the name suggests, we use comparison to evaluate and compare values between two or more data points.
     With comparison you can also easily find the lowest and highest values in the chart.


   <ul style="font-size:15px">
      Use a comparison visualizations chart to:
      <li>To compare two or more values side by side to visually grasp the difference</li>
      <li>To rank several data categories from best to worst or the other way around</li>
      <li>To show pattern recognition by visually highlighting gaps, spikes, outliers, or trends</li>
   </ul>
   <ul type="square" style="font-size:15px">
      Some of the most important comparison charts are:
        <li><a href="/18">HISTOGRAM CHART </a>  </li>
        <li><a href="/19">BAR CHART</a>  </li>
        <li><a href="/20">HEAT MAP</a>  </li>
        <li><a href="/21">HEXBIN CHART</a>  </li>
        <li><a href="/22">MOSAIC CHART</a>  </li>
        <li><a href="/23">MAP CHART</a>  </li>


    </ul>


</p>


</header>



<script>
// Used to toggle the menu on small screens when clicking on the menu button
function myFunction() {
    var x = document.getElementById("navDemo");
    if (x.className.indexOf("w3-show") == -1) {
        x.className += " w3-show";
    } else {
        x.className = x.className.replace(" w3-show", "");
    }
}
</script>

</body>
</html>
`

var compos=`<!DOCTYPE html>
<html>
<title>IIITK-VizB-Services</title>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="https://www.w3schools.com/w3css/4/w3.css">
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Lato">
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Montserrat">
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
<style>
body,h1,h2,h3,h4,h5,h6 {font-family: "Lato", sans-serif}
.w3-bar,h1,button {font-family: "Montserrat", sans-serif}
.fa-anchor,.fa-coffee {font-size:200px}
.s1{font-size: 15px}

.navbar {
    overflow: hidden;
    background-color: #e3b048;;
    font-family: Arial, Helvetica, sans-serif;
}

.navbar a {
    float: left;
    font-size: 16px;
    color: white;
    text-align: center;
    padding: 14px 16px;
    text-decoration: none;
}

.dropdown {
    float: left;
    overflow: hidden;
}

.dropdown .dropbtn {
    font-size: 16px;
    border: none;
    outline: none;
    color: white;
    padding: 14px 16px;
    background-color: inherit;
    font-family: inherit;
    margin: 0;
}

.navbar a:hover, .dropdown:hover .dropbtn {
    background-color: red;
}

.dropdown-content {
    display: none;
    position: absolute;
    background-color: #f9f9f9;
    min-width: 160px;
    box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
    z-index: 1;
}

.dropdown-content a {
    float: none;
    color: black;
    padding: 12px 16px;
    text-decoration: none;
    display: block;
    text-align: left;
}

.dropdown-content a:hover {
    background-color: #ddd;
}

.dropdown:hover .dropdown-content {
    display: block;
}
</style>
<body>

<!-- Navbar -->
<div class="navbar">
  <a href="sa.html">Home</a>
  <a href="contact.html">contact</a>
  <a href="about.html">about</a>
  <div class="dropdown">
    <button class="dropbtn">Services
      <i class="fa fa-caret-down"></i>
    </button>
    <div class="dropdown-content">
      <a href="distribution.html">distributioncharts</a>
      <a href="comparison.html">comparisoncharts</a>
      <a href="relationship.html">Relationcharts</a>
      <a href="composition.html">Compositioncharts</a>
    </div>
  </div>
</div>




<!-- Header -->
<header class="s1" style="padding:0px 0px">
   <h1 style="font-size:15px">Composition Charts<h1>
   <p style="font-size:15px"> Composition of data is probably the most misused method in data representation endeavors.

The idea is to show how individual parts make up the whole by combining them together and displaying them as a sum.
 Composition can also be used to show how a total value can be divided into parts or to highlight the significance of each
 part relative to the total value.


   <ul style="font-size:15px">
      Use a composition visualizations chart to:
      <li>Company market share and a few key players in the market</li>
       <li>Total country population by TOP religions, languages, or ethnical groups</li>
      <li> Total revenue, by TOP product lines, divisions, or regions</li>
   </ul>
   <ul type="square" style="font-size:15px">
      Some of the most important composition charts are:
        <li><a href="/15">PIECHART </a>  </li>
        <li><a href="/16">STACKEDAREA CHART</a>  </li>
        <li><a href="/1">SCATTER CHART</a>  </li>
        <li><a href="/17">WAFFLE CHART</a>  </li>
        <li><a href="/24">CORRELATION CHART</a>  </li>
        <li><a href="/3">AREA CHART</a>  </li>


    </ul>


</p>


</header>



<script>
// Used to toggle the menu on small screens when clicking on the menu button
function myFunction() {
    var x = document.getElementById("navDemo");
    if (x.className.indexOf("w3-show") == -1) {
        x.className += " w3-show";
    } else {
        x.className = x.className.replace(" w3-show", "");
    }
}
</script>

</body>
</html>
`

var relation=`<!DOCTYPE html>
<html>
<title>IIITK-VizB-Services</title>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="https://www.w3schools.com/w3css/4/w3.css">
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Lato">
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Montserrat">
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
<style>
body,h1,h2,h3,h4,h5,h6 {font-family: "Lato", sans-serif}
.w3-bar,h1,button {font-family: "Montserrat", sans-serif}
.fa-anchor,.fa-coffee {font-size:200px}
.s1{font-size: 15px}

.navbar {
    overflow: hidden;
    background-color: #e3b048;;
    font-family: Arial, Helvetica, sans-serif;
}

.navbar a {
    float: left;
    font-size: 16px;
    color: white;
    text-align: center;
    padding: 14px 16px;
    text-decoration: none;
}

.dropdown {
    float: left;
    overflow: hidden;
}

.dropdown .dropbtn {
    font-size: 16px;
    border: none;
    outline: none;
    color: white;
    padding: 14px 16px;
    background-color: inherit;
    font-family: inherit;
    margin: 0;
}

.navbar a:hover, .dropdown:hover .dropbtn {
    background-color: red;
}

.dropdown-content {
    display: none;
    position: absolute;
    background-color: #f9f9f9;
    min-width: 160px;
    box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
    z-index: 1;
}

.dropdown-content a {
    float: none;
    color: black;
    padding: 12px 16px;
    text-decoration: none;
    display: block;
    text-align: left;
}

.dropdown-content a:hover {
    background-color: #ddd;
}

.dropdown:hover .dropdown-content {
    display: block;
}
</style>
<body>

<!-- Navbar -->
<div class="navbar">
  <a href="sa.html">Home</a>
  <a href="contact.html">contact</a>
  <a href="about.html">about</a>
  <div class="dropdown">
    <button class="dropbtn">Services
      <i class="fa fa-caret-down"></i>
    </button>
    <div class="dropdown-content">
      <a href="distribution.html">distributioncharts</a>
      <a href="comparison.html">comparisoncharts</a>
      <a href="relationship.html">Relationcharts</a>
      <a href="composition.html">Compositioncharts</a>
    </div>
  </div>
</div>




<!-- Header -->
<header class="s1" style="padding:0px 0px">
   <h1 style="font-size:15px">Relationshio Charts</h1>
   <p style="font-size:15px">A good use of relationship graphs would be to demonstrate how something does or does not affect another variable positively or negatively.
     Somewhat similarly to distribution, you can use data relationships to find correlations, trends, patterns, clusters, averages, or outliers.


   <ul style="font-size:15px">
      Use a distribution or deviation visualizations chart to:
      <li>Visualize the distribution of ages in a group of people</li>
      <li>Visualize the correlations between cause and effect</li>
      <li>Spot bugs or constraints in production or quality control systems</li>
      <li>Visualize sales in various stores, to spot best and worst performing ones</li>
   </ul>
   <ul type="square" style="font-size:15px">
      Some of the most important relationship charts are:
        <li><a href="/9">BUBBLE CHART </a>  </li>
        <li><a href="/10">COUNT CHART</a>  </li>
        <li><a href="/11">JITTER CHART</a>  </li>
        <li><a href="/12">COORELOGRAM CHART</a>  </li>
        <li><a href="/13">MARGINAL HISTOGRAM CHART</a>  </li>
        <li><a href="/14">SCATTER ENCIRCLING CHART</a>  </li>
           <li><a href="/25">LOLLIPOP CHART</a>  </li>

    </ul>


</p>


</header>



<script>
// Used to toggle the menu on small screens when clicking on the menu button
function myFunction() {
    var x = document.getElementById("navDemo");
    if (x.className.indexOf("w3-show") == -1) {
        x.className += " w3-show";
    } else {
        x.className = x.className.replace(" w3-show", "");
    }
}
</script>

</body>
</html>
`

/// Infohandler service - IIITK-VizB-Services vizualization service
func infohandler(w http.ResponseWriter, req *http.Request){
   fmt.Fprintf(w, mainInfo)
}
func infohandler1(w http.ResponseWriter, req *http.Request){
   fmt.Fprintf(w, contact)
}
func infohandler2(w http.ResponseWriter, req *http.Request){
   fmt.Fprintf(w, about)
}
func infohandler3(w http.ResponseWriter, req *http.Request){
   fmt.Fprintf(w, dist)
}
func infohandler4(w http.ResponseWriter, req *http.Request){
   fmt.Fprintf(w, compare)
}
func infohandler5(w http.ResponseWriter, req *http.Request){
   fmt.Fprintf(w, compos)
}
func infohandler6(w http.ResponseWriter, req *http.Request){
   fmt.Fprintf(w, relation)
}
func bigdatacompscatterhandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatacompscatter(w, req)
}
func bigdatadistboxhandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatadistbox(w,req)
}
func bigdatacompareahandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatacomparea(w,req)
}
func bigdatadistdotboxhandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatadistdotbox(w,req)
}
func bigdatadistkernaldensityhandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatadistkernaldensity(w,req)
}
func bigdatadistnotchedboxhandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatadistnotchedbox(w,req)
}
func bigdatadisttufteboxhandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatadisttuftebox(w,req)
}
func bigdatadistviolinhandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatadistviolin(w,req)
}
func bigdatarelbubblehandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatarelbubble(w,req)
}
func bigdatareltablehandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatareltable(w,req)
}
func bigdatarelcounthandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatarelcount(w,req)
}
func Bigdatareljitterhandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatareljitter(w,req)
}
func Bigdatarelcorrelogramhandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatarelcorrelogram(w,req)
}
func Bigdatarelmarginalhistogramhandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatarelmarginalhistogram(w,req)
}
func Bigdatarelscatterencirclinghandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatarelscatterencircling(w,req)
}
func Bigdatacomppiehandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatacomppie(w,req)
}
func Bigdatacompstackedareahandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatacompstackedarea(w,req)
}
func Bigdatacompwafflehandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatacompwaffle(w,req)
}
func Bigdatacmphistogramhandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatacmphistogram(w,req)
}
func Bigdatacmpbdbarhandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatacmpbdbar(w,req)
}
func Bigdatacmpheathandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatacmpheat(w,req)
}
func Bigdatacmphexbinhandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatacmphexbin(w,req)
}
func Bigdatacmpmosaichandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatacmpmosaic(w,req)
}
func Bigdatacmpmaphandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatacmpmap(w,req)
}
func Bigdatacompcorrelationhandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatacompcorrelation(w,req)
}
func Bigdatarellollipophandler(w http.ResponseWriter, req *http.Request){
  bigdata.Bigdatarellollipop(w,req)
}
/**
  IIITK-VizB-Services - main function
*/
func main() {
//Step 1: Identify what kind of dataset
//Step 2: Identify the path of data or collect from database
//Step 3: Collect any hints from the user or automate ourselves.
//Step 4: Initiate the respective handlers :
// BigData vs. smalldata
// Composition; Comparison; DIstribution; Relation
// Modeling; Prediction graphs.
//Step 5: Execute the services.

   //Pursue with the steps

    fmt.Println("Initiating the IIITK-VizB-Services Services !!")
    fmt.Println("IIITK-VizB-Services starts at localhost:8080")
    router := mux.NewRouter()

    router.HandleFunc("/", infohandler)
      router.HandleFunc("/contact.html", infohandler1)
      router.HandleFunc("/about.html", infohandler2)
      router.HandleFunc("/distribution.html", infohandler3)
      router.HandleFunc("/comparison.html", infohandler4)
      router.HandleFunc("/composition.html", infohandler5)
      router.HandleFunc("/relationship.html", infohandler6)
    router.HandleFunc("/1", bigdatacompscatterhandler)
    router.HandleFunc("/2", bigdatadistboxhandler)
    router.HandleFunc("/3", bigdatacompareahandler)
    router.HandleFunc("/4", bigdatadistdotboxhandler)
    router.HandleFunc("/5", bigdatadistkernaldensityhandler)
    router.HandleFunc("/6", bigdatadistnotchedboxhandler)
    router.HandleFunc("/7", bigdatadisttufteboxhandler)
    router.HandleFunc("/8", bigdatadistviolinhandler)
    router.HandleFunc("/9", bigdatarelbubblehandler)
    router.HandleFunc("/10", bigdatarelcounthandler)
    router.HandleFunc("/11", Bigdatareljitterhandler)
    router.HandleFunc("/12", Bigdatarelcorrelogramhandler)
    router.HandleFunc("/13", Bigdatarelmarginalhistogramhandler)
    router.HandleFunc("/14", Bigdatarelscatterencirclinghandler)
    router.HandleFunc("/15", Bigdatacomppiehandler)
    router.HandleFunc("/16", Bigdatacompstackedareahandler)
    router.HandleFunc("/17", Bigdatacompwafflehandler)
    router.HandleFunc("/18", Bigdatacmphistogramhandler)
    router.HandleFunc("/19", Bigdatacmpbdbarhandler)
    router.HandleFunc("/20", Bigdatacmpheathandler)
    router.HandleFunc("/21", Bigdatacmphexbinhandler)
    router.HandleFunc("/22", Bigdatacmpmosaichandler)
    router.HandleFunc("/23", Bigdatacmpmaphandler)
    router.HandleFunc("/24", Bigdatacompcorrelationhandler)
    router.HandleFunc("/25", Bigdatarellollipophandler)

    log.Fatal(http.ListenAndServe(":8080", router))
}
