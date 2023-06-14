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


DROP TABLE IF EXISTS `documentreport`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `documentreport` (
  `id` int NOT NULL AUTO_INCREMENT,
  `documentuniqueid` varchar(256) NOT NULL,
  `timereceived` datetime NOT NULL,
  `typecode` varchar(256) NOT NULL,
  `cpr` varchar(256),
  `projectreference` varchar(512),
  `timesent` datetime,
  PRIMARY KEY (`id`),
  KEY `timereceived_documentreport_idx` (`timereceived`, `id`),
  KEY `documentuniqueid_documentreport_idx` (`documentuniqueid`, `id`),
  KEY `typecode_documentreport_idx` (`typecode`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `project`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `project` (
  `id` int NOT NULL AUTO_INCREMENT,
  `code` varchar(256) NOT NULL UNIQUE,
  `name` varchar(256) NOT NULL UNIQUE,
  `description` varchar(512),
  `path` varchar(1024) NOT NULL,
  `active` varchar(1) NOT NULL DEFAULT 'Y',
  `deleted` varchar(1) NOT NULL DEFAULT 'N',
  `projectreference` varchar(512),
  `timeupdated` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `projectreference_project_idx` (`projectreference`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1; 
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `projectcpr`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `projectcpr` (
  `id` int NOT NULL AUTO_INCREMENT,
  `projectid` int NOT NULL,
  `cpr` varchar(256) NOT NULL,
  `timeupdated` datetime NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT projectid_cpr_projectcpr_uc UNIQUE (projectid, cpr),
  CONSTRAINT projectid_projectcpr_fk foreign key (projectid) references project(id),
  KEY `projectcpr_projectid_idx` (`projectid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1; 
/*!40101 SET character_set_client = @saved_cs_client */;




/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
