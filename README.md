# Modern Dokuz Eylül Information Management System
# Modern Dokuz Eylül Bilgi İşletim Sistemi MDEBİS

In case of a need to reach the creators, please use the following e-mails
<ul>
  <li>Emircan Tepe: emircan.tepe@ceng.deu.edu.tr</li>
  <li>Deniz Küçükkara: deniz.kucukkara@ceng.deu.edu.tr</li>
</ul>


## System Architecture

## Database

<p> The database is created using MySql. It is currently in the third normal form. Below is the ER diagram of the DB.
	
</p>

## GO Architecture


## HOW TO RUN

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
Our database is created with data taken from the internet using web scraping methods
So instead of using a randomly chosen user, you can use the below ones to see the all the features of .
</p>
Student id: 2015537117 | 2019537310 | 2016537990
Lecturer id: 2020537421
Admin id: 1
All passwords are 354152.
