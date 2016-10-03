-- MySQL dump 10.13  Distrib 5.5.50, for debian-linux-gnu (x86_64)
--
-- Host: 52.87.63.135    Database: mobint
-- ------------------------------------------------------
-- Server version	5.5.46

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `config_device`
--

DROP TABLE IF EXISTS `config_device`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `config_device` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `id_device` varchar(100) NOT NULL,
  `operation_system` varchar(50) NOT NULL,
  `operation_system_version` varchar(20) NOT NULL,
  `device` varchar(50) NOT NULL,
  `type_connection` varchar(15) NOT NULL,
  `id_search` varchar(50) NOT NULL,
  `id_notification` varchar(300) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_config_device__search_idx` (`id_search`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `config_device`
--

LOCK TABLES `config_device` WRITE;
/*!40000 ALTER TABLE `config_device` DISABLE KEYS */;
/*!40000 ALTER TABLE `config_device` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `modality`
--

DROP TABLE IF EXISTS `modality`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `modality` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `price_km` decimal(15,2) NOT NULL DEFAULT '0.00',
  `time_km` int(11) NOT NULL DEFAULT '0',
  `id_player` int(11) NOT NULL,
  `price_base` decimal(15,2) NOT NULL DEFAULT '0.00',
  `price_time` decimal(15,2) NOT NULL DEFAULT '0.00',
  `minimum_price` decimal(15,2) NOT NULL DEFAULT '0.00',
  `active` int(2) NOT NULL DEFAULT '1',
  `edit_values` int(2) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `modality_player_FK` (`id_player`),
  CONSTRAINT `modality_player_FK` FOREIGN KEY (`id_player`) REFERENCES `player` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `modality`
--

LOCK TABLES `modality` WRITE;
/*!40000 ALTER TABLE `modality` DISABLE KEYS */;
INSERT INTO `modality` VALUES (1,'Táxi',2.30,300,3,0.00,0.00,0.00,1,0),(11,'EasyTaxi 30% OFF',1.30,180,4,2.00,0.20,4.00,1,0),(12,'99TOP',2.50,300,3,0.00,0.00,0.00,1,0),(13,'99POP',2.17,300,3,0.00,0.00,0.00,1,0),(15,'Easy Go',1.63,300,4,2.00,0.30,4.00,1,0),(16,'EasyPlus+',2.00,300,4,4.00,0.40,10.00,1,0),(17,'uberX',0.00,0,1,0.00,0.00,0.00,1,0),(18,'uberPOOL',0.00,0,1,0.00,0.00,0.00,1,0),(19,'uberBAG',0.00,0,1,0.00,0.00,0.00,1,0),(20,'UberBLACK',0.00,0,1,0.00,0.00,0.00,1,0),(21,'uberBIKE',0.00,0,1,0.00,0.00,0.00,1,0),(22,'Cabify Lite',0.00,0,2,0.00,0.00,0.00,1,0),(23,'Cabify Cab',0.00,0,2,0.00,0.00,0.00,1,0),(24,'Táxi 30% OFF',0.00,0,3,0.00,0.00,0.00,1,0),(25,'EasyTaxi',0.00,0,4,0.00,0.00,0.00,1,0);
/*!40000 ALTER TABLE `modality` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `modality_coverage`
--

DROP TABLE IF EXISTS `modality_coverage`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `modality_coverage` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `id_modality` int(11) NOT NULL,
  `zip_code_initial` varchar(8) NOT NULL DEFAULT '1000001',
  `zip_code_final` varchar(8) NOT NULL DEFAULT '8499999',
  PRIMARY KEY (`id`),
  KEY `modality_coverage_modality_FK` (`id_modality`),
  CONSTRAINT `modality_coverage_modality_FK` FOREIGN KEY (`id_modality`) REFERENCES `modality` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `modality_coverage`
--

LOCK TABLES `modality_coverage` WRITE;
/*!40000 ALTER TABLE `modality_coverage` DISABLE KEYS */;
INSERT INTO `modality_coverage` VALUES (16,12,'00000001','99999999'),(17,13,'00000001','99999999'),(18,1,'00000001','99999999'),(32,11,'00000001','99999999'),(33,15,'00000001','99999999'),(34,16,'00000001','99999999');
/*!40000 ALTER TABLE `modality_coverage` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `player`
--

DROP TABLE IF EXISTS `player`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `active` int(2) NOT NULL DEFAULT '1',
  `token` varchar(300) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `player`
--

LOCK TABLES `player` WRITE;
/*!40000 ALTER TABLE `player` DISABLE KEYS */;
INSERT INTO `player` VALUES (1,'UBER',1,'Token N7fFFJoeenUt06hYaIJ73plRNNZuaXawFTZZ0yVr'),(2,'CABIFY',1,'Bearer EYa4WcYJ74sN43vTJZUSpYmeNBm7HYVo8hSZQplYvg8'),(3,'99',1,'febed95f-239a-41c1-a76c-7ca9747c8f1b'),(4,'EASY TAXI',1,'febed95f-239a-41c1-a76c-7ca9747c8f1b');
/*!40000 ALTER TABLE `player` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `promotion`
--

DROP TABLE IF EXISTS `promotion`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `promotion` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `off` decimal(5,2) DEFAULT NULL,
  `limit_off` decimal(5,2) DEFAULT NULL,
  `active` int(2) NOT NULL DEFAULT '1',
  `new_modality` int(2) NOT NULL DEFAULT '1',
  `text` varchar(300) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `promotion`
--

LOCK TABLES `promotion` WRITE;
/*!40000 ALTER TABLE `promotion` DISABLE KEYS */;
INSERT INTO `promotion` VALUES (1,'Claro Clube',50.00,25.00,1,1,'Desconto de 50% sobre o valor da corrida, limitados a R$ 25,00, para clientes Claro Clube. Em meios de pagamento, selecione a opção Claro Clube (50% de desconto)'),(2,'Santander Meia Bandeira',50.00,15.00,1,1,'Desconto de 50% sobre o valor da corrida, limitados a R$ 15,00, para clientes Santander. Em pagamento & promoção, selecione a opção Santander Meia Bandeira – 50%'),(3,'SOMOSMOBILIDADE',25.00,15.00,1,1,'Desconto de 25% sobre o valor da corrida, limitados a R$ 15,00. Promoção válida para apenas 3 (três) corridas. Antes de confirmar a corrida, vá em Promoções e digite o código SOMOSMOBILIDADE para obter o desconto.');
/*!40000 ALTER TABLE `promotion` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `promotion_available`
--

DROP TABLE IF EXISTS `promotion_available`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `promotion_available` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `monday` int(1) NOT NULL DEFAULT '0',
  `tuesday` int(1) NOT NULL DEFAULT '0',
  `wednesday` int(1) NOT NULL DEFAULT '0',
  `thursday` int(1) NOT NULL DEFAULT '0',
  `friday` int(1) NOT NULL DEFAULT '0',
  `saturday` int(1) NOT NULL DEFAULT '0',
  `sunday` int(1) NOT NULL DEFAULT '0',
  `start_hour` varchar(5) NOT NULL DEFAULT '00:01',
  `end_hour` varchar(5) NOT NULL DEFAULT '23:59',
  `id_promotion` bigint(20) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `promotion_available_promotion_FK` (`id_promotion`),
  CONSTRAINT `promotion_available_promotion_FK` FOREIGN KEY (`id_promotion`) REFERENCES `promotion` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `promotion_available`
--

LOCK TABLES `promotion_available` WRITE;
/*!40000 ALTER TABLE `promotion_available` DISABLE KEYS */;
INSERT INTO `promotion_available` VALUES (1,0,0,0,0,1,0,0,'20:00','23:59',1),(2,1,0,0,0,0,0,0,'00:00','06:00',1),(3,0,0,0,0,0,1,1,'00:00','23:59',1),(4,0,0,0,1,1,1,1,'20:00','23:59',2),(5,1,0,0,0,1,1,1,'00:00','06:00',2),(6,1,1,1,1,1,1,1,'00:00','23:59',3);
/*!40000 ALTER TABLE `promotion_available` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `promotion_modality`
--

DROP TABLE IF EXISTS `promotion_modality`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `promotion_modality` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `id_promotion` bigint(20) NOT NULL,
  `id_modality` int(11) NOT NULL,
  `initial_at` date DEFAULT NULL,
  `final_at` date DEFAULT NULL,
  `exibition_name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `promotion_modality_modality_FK` (`id_modality`),
  KEY `promotion_modality_promotion_FK` (`id_promotion`),
  CONSTRAINT `promotion_modality_modality_FK` FOREIGN KEY (`id_modality`) REFERENCES `modality` (`id`) ON DELETE CASCADE,
  CONSTRAINT `promotion_modality_promotion_FK` FOREIGN KEY (`id_promotion`) REFERENCES `promotion` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `promotion_modality`
--

LOCK TABLES `promotion_modality` WRITE;
/*!40000 ALTER TABLE `promotion_modality` DISABLE KEYS */;
INSERT INTO `promotion_modality` VALUES (1,1,24,NULL,NULL,'Claro Clube'),(2,1,13,NULL,NULL,'Claro Clube'),(3,1,12,NULL,NULL,'Claro Clube'),(4,1,1,NULL,NULL,'Claro Clube'),(5,2,25,NULL,NULL,'Santander ½ bandeira'),(6,2,16,NULL,NULL,'Santander ½ bandeira'),(7,2,15,NULL,NULL,'Santander ½ bandeira'),(8,2,11,NULL,NULL,'Santander ½ bandeira'),(27,3,22,NULL,NULL,'SOMOSMOBILIDADE'),(28,3,23,NULL,NULL,'SOMOSMOBILIDADE');
/*!40000 ALTER TABLE `promotion_modality` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `promotion_modality_coverage`
--

DROP TABLE IF EXISTS `promotion_modality_coverage`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `promotion_modality_coverage` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `id_promotion_modality` bigint(20) NOT NULL,
  `zip_code_initial` varchar(8) DEFAULT NULL,
  `zip_code_final` varchar(8) DEFAULT NULL,
  `state` varchar(2) DEFAULT NULL,
  `city` varchar(150) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `promotion_modality_coverage_promotion_modality_FK` (`id_promotion_modality`),
  CONSTRAINT `promotion_modality_coverage_promotion_modality_FK` FOREIGN KEY (`id_promotion_modality`) REFERENCES `promotion_modality` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `promotion_modality_coverage`
--

LOCK TABLES `promotion_modality_coverage` WRITE;
/*!40000 ALTER TABLE `promotion_modality_coverage` DISABLE KEYS */;
INSERT INTO `promotion_modality_coverage` VALUES (17,5,NULL,NULL,'SP','São Paulo'),(18,6,NULL,NULL,'SP','São Paulo'),(19,7,NULL,NULL,'SP','São Paulo'),(20,8,NULL,NULL,'SP','São Paulo');
/*!40000 ALTER TABLE `promotion_modality_coverage` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `search`
--

DROP TABLE IF EXISTS `search`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `search` (
  `id` varchar(50) NOT NULL,
  `date_time` datetime NOT NULL,
  `start_address_id` varchar(50) NOT NULL,
  `end_address_id` varchar(50) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `search_search_address_start_FK` (`start_address_id`),
  KEY `search_search_address_end_FK` (`end_address_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `search`
--

LOCK TABLES `search` WRITE;
/*!40000 ALTER TABLE `search` DISABLE KEYS */;
/*!40000 ALTER TABLE `search` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `search_address`
--

DROP TABLE IF EXISTS `search_address`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `search_address` (
  `id` varchar(50) NOT NULL,
  `lat` decimal(15,12) NOT NULL,
  `lng` decimal(15,12) NOT NULL,
  `address` varchar(100) DEFAULT NULL,
  `district` varchar(100) DEFAULT NULL,
  `city` varchar(100) DEFAULT NULL,
  `state` varchar(2) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `search_address`
--

LOCK TABLES `search_address` WRITE;
/*!40000 ALTER TABLE `search_address` DISABLE KEYS */;
/*!40000 ALTER TABLE `search_address` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `search_results`
--

DROP TABLE IF EXISTS `search_results`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `search_results` (
  `id` varchar(100) NOT NULL,
  `id_player` int(11) NOT NULL,
  `modality` varchar(25) DEFAULT NULL,
  `waiting_time` bigint(20) NOT NULL,
  `tax_value` varchar(50) NOT NULL,
  `id_search` varchar(50) NOT NULL,
  `multiplier` decimal(5,2) DEFAULT NULL,
  `promotion` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `search_results_search_FK` (`id_search`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `search_results`
--

LOCK TABLES `search_results` WRITE;
/*!40000 ALTER TABLE `search_results` DISABLE KEYS */;
/*!40000 ALTER TABLE `search_results` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `search_selected`
--

DROP TABLE IF EXISTS `search_selected`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `search_selected` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `id_search_results` varchar(50) NOT NULL,
  `date_time_click` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_search_selected_results_idx` (`id_search_results`),
  CONSTRAINT `search_selected_search_results_FK` FOREIGN KEY (`id_search_results`) REFERENCES `search_results` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `search_selected`
--

LOCK TABLES `search_selected` WRITE;
/*!40000 ALTER TABLE `search_selected` DISABLE KEYS */;
/*!40000 ALTER TABLE `search_selected` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `mail` varchar(100) NOT NULL,
  `username` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'rafaelvalerini@gmail.com','rafaelvalerini','rafael123*'),(2,'leandroceccato@mobint.com.br','ceccato','123456'),(3,'marcio@thinktwice.com.br','mbern','nelson');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'mobint'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-10-02 23:55:12
