-- MySQL dump 10.13  Distrib 8.0.31, for Win64 (x86_64)
--
-- Host: localhost    Database: mdebis
-- ------------------------------------------------------
-- Server version	8.0.31

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `department`
--

DROP TABLE IF EXISTS `department`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `department` (
  `Department_Id` int NOT NULL,
  `Head_Lecturer_Id` int NOT NULL,
  `Name` varchar(300) NOT NULL,
  PRIMARY KEY (`Department_Id`),
  KEY `fk_Department_Lecturer1_idx` (`Head_Lecturer_Id`),
  CONSTRAINT `fk_Department_Lecturer1` FOREIGN KEY (`Head_Lecturer_Id`) REFERENCES `lecturer` (`Lecturer_Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `department`
--

LOCK TABLES `department` WRITE;
/*!40000 ALTER TABLE `department` DISABLE KEYS */;
INSERT INTO `department` VALUES (1,2019501771,'Elementary Teacher Education'),(2,2018502129,'Pre - School Teacher Education'),(3,2016503694,'Guidance And Psychological Counseling'),(4,2017504186,'Fine Arts Teacher Education'),(5,2014505605,'Music Teacher Education'),(6,2015506534,'English Language Teacher Education'),(7,2001507684,'French Language Teacher Education'),(8,2002508600,'German Language Teacher Education'),(9,2011509845,'Biology Teacher Education'),(10,2015510518,'Chemistry Teacher Education'),(11,2009511188,'Elementary Mathematics Teacher Education'),(12,2000512838,'Mathematics Teacher Education'),(13,2008513716,'Physics Teacher Education'),(14,2006514620,'Science Teacher Education'),(15,2003515110,'Special Teacher Education'),(16,2006516812,'Geography Teacher Education'),(17,2022517216,'History Teacher Education'),(18,2014518945,'Social Studies Teacher Education'),(19,2002519480,'Turkish Language And Literature Teacher Education'),(20,2006520432,'Turkish Language Teacher Education'),(21,2017521305,'Architecture'),(22,2022522785,'City And Regional Planning'),(23,2018523773,'Business Administration'),(24,2021524721,'Business Administration (uolp-suny Albany)'),(25,2007525462,'Economics  (uolp-suny Albany)'),(26,2013526761,'Economics'),(27,2006527801,'International Relations  (uolp-suny Albany)'),(28,2013528917,'International Business And Trade'),(29,2011529563,'International Relations'),(30,2018530490,'Tourism Management'),(31,2008531599,'Econometrics'),(32,2013532318,'Labour Economics And Industrial Relations'),(33,2007533235,'Management Information Systems'),(34,2022534126,'Public Administration'),(35,2007535738,'Public Finance'),(36,2001536542,'Civil Engineering'),(37,2010537195,'Computer Engineering'),(38,2002538777,'Electrical And Electronics Engineering'),(39,2022539706,'Environmental Engineering'),(40,2014540373,'Geological Engineering'),(41,2010541266,'Geophysical Engineering'),(42,2020542257,'Industrial Engineering'),(43,2010543989,'Mechanical Engineering'),(44,2006544990,'Metallurgical And Materials Engineering'),(45,2019545835,'Mining Engineering'),(46,2022546881,'Textile Engineering'),(47,2009547575,'Animation Film Design And Directing'),(48,2005548958,'Sound Editing And Cinematography'),(49,2004549931,'Graphic'),(50,2007550690,'Graphic Design'),(51,2008551792,'Illustration And Printmaking'),(52,2019552194,'Music Sciences'),(53,2003553551,'Music Technology'),(54,2008554481,'Painting'),(55,2013555252,'Performance Arts'),(56,2002556252,'Stage Design'),(57,2018557682,'Photography'),(58,2012558443,'Sculpture'),(59,2015559237,'Accessory Design'),(60,2004560470,'Fashion Clothing Design'),(61,2009561586,'Textile And Fashion Design Department'),(62,2006562347,'Textile Design'),(63,2014563842,'Restorations Of Ancient Tile'),(64,2000564831,'Law'),(65,2011565896,'American Culture And Literature'),(66,2010566588,'Archeology'),(67,2014567890,'Comparative Literature'),(68,2007568911,'History'),(69,2007569454,'Linguistics'),(70,2009570649,'Museology'),(71,2017571355,'Philosophy'),(72,2003572980,'Psychology'),(73,2002573652,'Sociology'),(74,2019574915,'Turkish Language And Literature'),(75,2006575328,'Medicine'),(76,2014576705,'Nursing'),(77,2001577288,'Physiotherapy And Rehabilitation'),(78,2003578904,'Biology'),(79,2020579722,'Chemistry'),(80,2022580116,'Computer Science'),(81,2002581426,'Mathematics'),(82,2017582199,'Physics'),(83,2001583955,'Statistics'),(84,2016584304,'Primary Religious Culture And Ethics Teacher Education'),(85,2008585404,'Theology'),(86,2013586731,'Gastronomy And Culinary Arts'),(87,2019587766,'Logistics Management'),(88,2021588451,'Marine Engineering'),(89,2011589402,'Marine Transportation Engineering'),(90,2005590858,'Maritime Business Administration'),(91,2018591958,'Coaching Education'),(92,2013592632,'Physical Education And Sports Teaching'),(93,2012593353,'International Trade'),(94,2000594310,'Tourism And Hotel Management'),(95,2012595765,'Composition'),(96,2015596191,'Guitar'),(97,2006597886,'Harp'),(98,2001598362,'Piano'),(99,2012599315,'Cello'),(100,2004600701,'Contrabass'),(101,2001601433,'Viola'),(102,2019602110,'Violin'),(103,2009603893,'Ballet'),(104,2006604424,'Opera');
/*!40000 ALTER TABLE `department` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-12-20  0:46:06
