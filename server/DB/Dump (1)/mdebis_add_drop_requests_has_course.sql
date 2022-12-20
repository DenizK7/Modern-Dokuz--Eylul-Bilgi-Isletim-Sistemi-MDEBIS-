CREATE DATABASE  IF NOT EXISTS `mdebis` /*!40100 DEFAULT CHARACTER SET utf8mb3 */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `mdebis`;
-- MySQL dump 10.13  Distrib 8.0.29, for macos12 (x86_64)
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
-- Table structure for table `add_drop_requests_has_course`
--

DROP TABLE IF EXISTS `add_drop_requests_has_course`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `add_drop_requests_has_course` (
  `Add_Drop_Requests_Request_Id` int NOT NULL,
  `Course_Course_Id` int NOT NULL,
  PRIMARY KEY (`Add_Drop_Requests_Request_Id`,`Course_Course_Id`),
  KEY `fk_Add_Drop_Requests_has_Course_Course1_idx` (`Course_Course_Id`),
  KEY `fk_Add_Drop_Requests_has_Course_Add_Drop_Requests1_idx` (`Add_Drop_Requests_Request_Id`),
  CONSTRAINT `fk_Add_Drop_Requests_has_Course_Add_Drop_Requests1` FOREIGN KEY (`Add_Drop_Requests_Request_Id`) REFERENCES `add_drop_requests` (`Request_Id`),
  CONSTRAINT `fk_Add_Drop_Requests_has_Course_Course1` FOREIGN KEY (`Course_Course_Id`) REFERENCES `course` (`Course_Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `add_drop_requests_has_course`
--

LOCK TABLES `add_drop_requests_has_course` WRITE;
/*!40000 ALTER TABLE `add_drop_requests_has_course` DISABLE KEYS */;
/*!40000 ALTER TABLE `add_drop_requests_has_course` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-12-21  0:43:57
