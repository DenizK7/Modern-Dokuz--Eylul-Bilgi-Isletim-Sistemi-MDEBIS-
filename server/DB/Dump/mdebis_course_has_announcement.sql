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
-- Table structure for table `course_has_announcement`
--

DROP TABLE IF EXISTS `course_has_announcement`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `course_has_announcement` (
  `Course_Announcement_Id` int NOT NULL AUTO_INCREMENT,
  `Course_Id` int NOT NULL,
  `Title` tinytext,
  `Content` longtext,
  `Lecturer_Id` int DEFAULT NULL,
  PRIMARY KEY (`Course_Announcement_Id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `course_has_announcement`
--

LOCK TABLES `course_has_announcement` WRITE;
/*!40000 ALTER TABLE `course_has_announcement` DISABLE KEYS */;
INSERT INTO `course_has_announcement` VALUES (1,288,'FIRST ANNOUNCEMENT OF THE COURSE','WELCOME TO CLASS!! WISH YOU ALL SUCCESS!',2000506140),(8,288,'SECOND ANNOUNCEMENT OF THE COURSE','PLEASE before the class, read the chapter shared with you in the resources page of the class.',2000506140),(9,288,'SECOND ANNOUNCEMENT OF THE COURSE','PLEASE before the class, read the chapter shared with you in the resources page of the class.',2000506140),(10,288,'SECOND ANNOUNCEMENT OF THE COURSE','PLEASE before the class, read the chapter shared with you in the resources page of the class.',2000506140),(11,288,'SECOND ANNOUNCEMENT OF THE COURSE','PLEASE before the class, read the chapter shared with you in the resources page of the class.',2000506140),(12,288,'SECOND ANNOUNCEMENT OF THE COURSE','PLEASE before the class, read the chapter shared with you in the resources page of the class.',2000506140);
/*!40000 ALTER TABLE `course_has_announcement` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-12-18 20:13:00
