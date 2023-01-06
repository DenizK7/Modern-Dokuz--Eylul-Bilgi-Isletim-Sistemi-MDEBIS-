# Modern Dokuz Eylül Information Management System
# Modern Dokuz Eylül Bilgi İşletim Sistemi MDEBİS

In case of a need to reach the creators, please use the following e-mails

Emircan Tepe: emircan.tepe@ceng.deu.edu.tr
  
Deniz Küçükkara: deniz.kucukkara@ceng.deu.edu.tr

# Contents
1.	[SYSTEM ARCHITRECTURES](#SystemArc)
2.	[USED TECHNOLOGIES, TOOLS AND PROGRAMMING LANGUAGES](#UsedTechnology)
3.	[DATABASE](#Database)
4.	[HOW TO RUN](#HowToRun)

## SYSTEM ARCHTIRECTURE <a name="SystemArc"></a>
The well-known and widely used Model-View-Controller (MVC) architecture is applied with its all design patterns. It is based on the separating information from its representation. Below figure is the general representation of MVC.

<center>
<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/a/a0/MVC-Process.svg/1000px-MVC-Process.svg.png">

Representation of MVC, taken from [Wikipedia](https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93controller)
</center>


## USED TECHNOLOGIES, TOOLS AND PROGRAMMING LANGUAGES <a name="UsedTechnology"></a>

Below are given the programming languages, technology and tools throughout the project development.
<ol>
<li><b>GOLANG </b>(<i>Model</i>): A functional programming language designed by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but with additionally memory safety, garbage collection, structural typing, and CSP-style concurrency.</li>
<li><b>HTML</b> (<i>View</i>): HTML (Hypertext Markup Language) is a standard markup language for creating web pages and web applications. It is used to structure and organize content on the web, and to define the meaning and structure of that content. HTML consists of a series of elements, which are used to enclose, format, and link different types of content. When a web browser receives an HTML document, it reads the HTML code and uses it to render the content of the page on the user's screen.</li>
<li><b>Cascading Style Sheets</b> (<i>CSS</i>) (View): CSS is a stylesheet language used for describing the look and formatting of a document written in HTML. CSS is used to define the layout and design of web pages, including colors, fonts, and responsive design. It is used to apply styles to web pages, such as specifying that certain text should be displayed in a particular font or that certain elements should be positioned in a certain way on the page. By separating the presentation of content from the content itself, CSS allows web designers and developers to create more flexible and adaptable websites.</li>
<li><b>JavaScript</b> (View & Controller): JavaScript is a programming language that is commonly used to create interactive effects within web browsers. It is an implementation of the ECMAScript specification and is used to enable web pages to be dynamic and interactive. JavaScript code is run on the client side, meaning it is executed by the user's web browser rather than on the server. It can be used to create things like dropdown menus, form validation, and interactive maps. JavaScript is an essential component of modern web development and is used on the majority of websites.</li>
<li><b>REACT </b>(View & Controller): React is a JavaScript library for building user interfaces that was developed by Facebook. It is commonly used for building single-page applications and mobile applications, and allows developers to create reusable UI components that can be rendered on the server or the client. React uses a virtual DOM to optimize the rendering of UI elements, and has a declarative programming style, which means that the code specifies what the UI should look like rather than describing the steps to create it. </li>
<li><b>MYSQL </b>(DB): MYSQL is one of the well-known and most-used relational DBMS.</li>
<li><b>Python</b>: Python is not used any part of the project that can effect any performance. It is only used as a tool when doing web-scraping </li>

## DATABASE <a name="Database"></a>

<p> The database is created using MySql. It is currently in the third normal form. Below is the ER diagram of the DB.
</p>

<center> 
<img src="diagrams/ER_Diagram.png">

ER diagram, created on [VisualParadigm](https://online.visual-paradigm.com/)
</center>



## HOW TO RUN <a name="HowToRun"></a>

<p>
To be able to run this project succesfully with its all components, one needs following tools:
<ul>
	<li>NodeJs (https://nodejs.org/en/download/)</li>
	<li>REACT library (run the command "npm install", then "npm start" under the mdebis folder)</li>
	<li>GOLANG (https://go.dev/dl/)</li>
	<li>MYSQL (https://www.mysql.com/downloads/)</li>
</ul>
</p>
<p> 
Lastly, it is also a need to import the database that will be used by the program to your local mysql server

You can use the dump file provided under this folder to import the mdebis database 

After all these steps, to run the mdebis, please follow the below steps in order:
</p>
<ol>
	<li>Go to server folder</li>
	<li>Run "go run ./"</li>
	<li>If the above command produces an error, please run "go mod init mdebis", then "go install .", then "go run ./"</li>
	<li>GO TO MDEBIS FOLDER</li>
	<li>Run "npm install" command. this command may take some time, please wait.</li>
	<li>Run "npm start" command. after compilation, home page of mdebis will start on your main web browser.</li>
	<li>DONE :)</li>
</ol>
<p>
Our database is created with data taken from the internet using web scraping methods.

So instead of using a randomly chosen user, you can use the below ones to see the all the features of mdebis.
</p>
Student id: 2015537117 | 2019537310 | 2016537990 

Lecturer id: 2020537421

Admin id: 1  

All passwords are 354152.
