CREATE DATABASE  IF NOT EXISTS `mdebis` /*!40100 DEFAULT CHARACTER SET utf8mb3 */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `mdebis`;
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
-- Table structure for table `general_announcement`
--

DROP TABLE IF EXISTS `general_announcement`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `general_announcement` (
  `General_Announcements_Id` int NOT NULL,
  `Title` tinytext NOT NULL,
  `Content` longtext NOT NULL,
  `link` varchar(1000) DEFAULT NULL,
  PRIMARY KEY (`General_Announcements_Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `general_announcement`
--

LOCK TABLES `general_announcement` WRITE;
/*!40000 ALTER TABLE `general_announcement` DISABLE KEYS */;
INSERT INTO `general_announcement` VALUES (1,'7417 SAYILI AF KANUNU HAKKINDA DUYURU','5 Temmuz 2022 tarih ve 31887 sayılı Resmi Gazete’de yayımlanarak yürürlüğe giren 7417 Sayılı Kanun’un 35 Maddesi ile 2547 Sayılı Kanuna Eklenen Geçici 83 üncü maddeden yararlanmak isteyen (öğrenci affı) 5 Temmuz 2022 tarihinden önce ilişiği kesilen üniversitemiz öğrencileri, en geç 05.07.2022 tarihinden itibaren 4 ay içinde, kayıtlarının silindiği veya kayıt hakkı kazandıkları Fakülte/Enstitü veya Yüksekokula aşağıdaki başvuru formu ve istenen belgeler ile başvurmaları ve kanunda belirtilen şartları taşımaları halinde 2022-2023 eğitim öğretim yılında öğrenime başlayabileceklerdir.','https://www.deu.edu.tr/duyurular/7417-sayili-af-kanunu-hakkinda-duyuru/'),(2,'Lisansüstü Programlara Öğrenci Kabulü (Student Admission to Graduate Programs)','2022-2023 eğitim-öğretim yılı güz yarıyılında Üniversitemize bağlı enstitülerin lisansüstü programlarına kabul edilecek öğrenci kontenjanları, başvuru ve kabul şartları aşağıda yer alan web sayfalarında ilan edilmiştir. (In the fall semester of the 2022-2023 academic year, the quota of students to be admitted to the graduate programs of the institutes of our university, application and admission requirements are announced on the following web pages.)','https://www.deu.edu.tr/duyurular/lisansustu-programlara-ogrenci-kabulu-student-admission-to-graduate-programs-2/'),(3,'2022-2023 Eğitim-Öğretim Yılı Güz Yarıyılı Yatay Geçiş Başvuru İşlemleri','Yatay geçiş başvuruları https://ogrbasvuru.deu.edu.tr/login.php internet üzerinden ONLINE yapılacaktır. (Belgeleri eksik olan ve/veya tercih yapmayan adayların başvuruları değerlendirmeye alınmaz).','https://www.deu.edu.tr/duyurular/2022-2023-egitim-ogretim-yili-guz-yariyili-yatay-gecis-basvuru-islemleri/'),(4,'Sosyal Hizmetler Kanunu Kapsamında Üniversitemize Yerleşen Adaylardan İstenilen Belgele','-Yerleşen adayların istenilen belgeler ve başvuru dilekçesi (web sayfasından alınacaktır) ile birlikte 10/02/2021 tarihi mesai bitimine kadar Dokuz Eylül Üniversitesi Personel Daire Başkanlığı Atama Şube Müdürlüğüne başvurmaları gerekmektedir.','https://www.deu.edu.tr/duyurular/sosyal-hizmetler-kanunu-kapsaminda-universitemize-yerlesen-adaylardan-istenilen-belgeler-6/');
/*!40000 ALTER TABLE `general_announcement` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-12-19 20:44:14
