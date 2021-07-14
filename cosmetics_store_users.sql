-- MySQL dump 10.13  Distrib 8.0.25, for macos11 (x86_64)
--
-- Host: localhost    Database: cosmetics_store
-- ------------------------------------------------------
-- Server version	8.0.25

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
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `full_name` varchar(100) NOT NULL,
  `email` varchar(191) NOT NULL,
  `password` longtext NOT NULL,
  `phone` varchar(12) DEFAULT NULL,
  `address` longtext,
  `date_of_birth` text,
  `gender` enum('Male','Female') DEFAULT 'Male',
  `avatar` varchar(300) DEFAULT NULL,
  `role_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email_UNIQUE` (`email`),
  KEY `role_id` (`role_id`),
  CONSTRAINT `users_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=76 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (3,'Nhân béo','asd@gmail.com','$2a$14$1Xw.SvkqhmMd4f1VTvNkw.Yz9O2YPOJsOCKVZWjL5AULNa.0KR1ui','0981838161','abc xyz gmh','26/04/1997','Male','https://scontent.fhan3-2.fna.fbcdn.net/v/t1.6435-9/196437483_838057373517144_5524917560022717507_n.jpg?_nc_cat=107&_nc_rgb565=1&ccb=1-3&_nc_sid=09cbfe&_nc_ohc=7fW3zg5T0sMAX9nGc4K&_nc_ht=scontent.fhan3-2.fna&oh=8995567f610b94bd7c5c72a37b79db27&oe=60F0F39D',1),(40,'abc','asdaaaaaa@gmail.com','$2a$14$0ht/IELs98LLBOxkBtLzlu2rhhvWEllTy6EMCrdbXjeGPlpSP9bY2',NULL,NULL,NULL,'Male',NULL,2),(47,'abc','asdaaaacaa@gmail.com','$2a$14$gtmGLtZ/a5.49hnlLOiDSeLN4eNubdLvW3UyxOHEcB3uBbw883aYK',NULL,NULL,NULL,'Male',NULL,2),(51,'nhan','ascccd@gmail.com','$2a$14$shyVzfg4UkxnF6qrslwGT.TC.KqPO.rxX8qlDEGEbDLodYkFM8mvS',NULL,NULL,NULL,'Male',NULL,2),(53,'nhanbeo','abc@gmail.com','$2a$14$g2cDDan0yYob/qxU1vpAauOF.5Oef5jfbBvvxhQg7vn4bo0TPxUaS',NULL,NULL,NULL,'Male',NULL,2),(55,'nhanbeo','acccbc@gmail.com','$2a$14$zcaBdFCZppWqH645xG8FCeXOs/wHDdwsDcico3c0y9f8t8AqphuvC',NULL,NULL,NULL,'Male',NULL,2),(56,'nhan','abccccxc@gmail.com','$2a$14$NthTVm9TyEAjQWcl/NeyUex9keUY/s/3tE2Oq54vt/OUVGubUMTXG',NULL,NULL,NULL,'Male',NULL,2),(59,'nhan','abccccxxc@gmail.com','$2a$14$RDkFIq0b9lBudC1Udv.W6.tc6/kf0TNvxuQa4lcjWEgjDFohkp0vW',NULL,NULL,NULL,'Male',NULL,2),(61,'nhan','aaaasd@gmail.com','$2a$14$P6/69VNoLszgrGQ5QTWl0eIYZ5xlF5vEdqAJb54R/J7OtdstIMpZy',NULL,NULL,NULL,'Male',NULL,2),(64,'nhan','aaacasd@gmail.com','$2a$14$ARZCyvDAwjdVqMRJD7k76u2Tb3cC53wGoGVkUtT1wTC25RPaUZAlO',NULL,NULL,NULL,'Male',NULL,2),(66,'nhan','ascxxxd@gmail.com','$2a$14$2GNtGuJ20mRphgoFxGPfQeaWhN6sOWpV8KPeVaSiNtj8VGXjYQCre',NULL,NULL,NULL,'Male',NULL,2),(68,'nhan','ascxxxdx@gmail.com','$2a$14$5gAz0tcpOhZC3pNPVWRJpO8bhaRs1Rj25cnHyy3Mep9HxJgKl/Iu6',NULL,NULL,NULL,'Male',NULL,2),(74,'nhan','abczzcc@gmail.com','$2a$14$s6HOk5VlLOqvmS3AlKMEcuLSWXxfDGCJ7VvqXCjOrwuuvJx0vrBP2',NULL,NULL,NULL,'Male',NULL,2);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-07-14 11:25:45
