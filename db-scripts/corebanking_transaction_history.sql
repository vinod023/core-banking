-- MySQL dump 10.13  Distrib 8.0.29, for Win64 (x86_64)
--
-- Host: localhost    Database: corebanking
-- ------------------------------------------------------
-- Server version	8.0.29

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
-- Table structure for table `transaction_history`
--

DROP TABLE IF EXISTS transaction_history;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE transaction_history (
  id int NOT NULL AUTO_INCREMENT,
  userid int NOT NULL,
  `type` varchar(45) DEFAULT NULL,
  amount decimal(10,0) DEFAULT NULL,
  balance decimal(10,0) DEFAULT NULL,
  trans_userid int DEFAULT NULL,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transaction_history`
--

/*!40000 ALTER TABLE transaction_history DISABLE KEYS */;
INSERT INTO transaction_history (id, userid, type, amount, balance, trans_userid, created_at, updated_at) VALUES (1,2,'DEPOSIT',200,200,0,'2022-07-18 02:10:57','2022-07-18 02:10:57');
INSERT INTO transaction_history (id, userid, type, amount, balance, trans_userid, created_at, updated_at) VALUES (2,2,'WITHDRAW',100,100,0,'2022-07-18 11:09:40','2022-07-18 11:09:40');
INSERT INTO transaction_history (id, userid, type, amount, balance, trans_userid, created_at, updated_at) VALUES (3,3,'DEPOSIT',301,301,0,'2022-07-19 11:10:02','2022-07-19 11:10:02');
INSERT INTO transaction_history (id, userid, type, amount, balance, trans_userid, created_at, updated_at) VALUES (4,3,'WITHDRAW',100,201,0,'2022-07-19 11:10:19','2022-07-19 11:10:19');
INSERT INTO transaction_history (id, userid, type, amount, balance, trans_userid, created_at, updated_at) VALUES (5,5,'DEPOSIT',800,800,0,'2022-07-18 11:53:02','2022-07-18 11:53:02');
INSERT INTO transaction_history (id, userid, type, amount, balance, trans_userid, created_at, updated_at) VALUES (6,5,'WITHDRAW',151,649,0,'2022-07-18 11:54:14','2022-07-18 11:54:14');
INSERT INTO transaction_history (id, userid, type, amount, balance, trans_userid, created_at, updated_at) VALUES (7,5,'WITHDRAW',188,461,0,'2022-07-18 11:58:32','2022-07-18 11:58:32');
INSERT INTO transaction_history (id, userid, type, amount, balance, trans_userid, created_at, updated_at) VALUES (8,5,'DEPOSIT',501,962,0,'2022-07-19 11:59:37','2022-07-19 11:59:37');
INSERT INTO transaction_history (id, userid, type, amount, balance, trans_userid, created_at, updated_at) VALUES (9,5,'DEPOSIT',11,973,0,'2022-07-19 12:00:17','2022-07-19 12:00:17');
INSERT INTO transaction_history (id, userid, type, amount, balance, trans_userid, created_at, updated_at) VALUES (10,5,'TRANSFER',50,973,1,'2022-07-19 12:02:46','2022-07-19 12:02:46');
INSERT INTO transaction_history (id, userid, type, amount, balance, trans_userid, created_at, updated_at) VALUES (11,5,'TRANSFER',151,973,1,'2022-07-20 12:13:02','2022-07-20 12:13:02');
INSERT INTO transaction_history (id, userid, type, amount, balance, trans_userid, created_at, updated_at) VALUES (12,5,'TRANSFER',698,426,1,'2022-07-20 12:17:09','2022-07-20 12:17:09');
INSERT INTO transaction_history (id, userid, type, amount, balance, trans_userid, created_at, updated_at) VALUES (13,5,'TRANSFER',589,535,1,'2022-07-20 12:22:34','2022-07-20 12:22:34');
/*!40000 ALTER TABLE transaction_history ENABLE KEYS */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-07-21 22:52:12
