-- MySQL dump 10.13  Distrib 5.7.18, for Linux (x86_64)
--
-- Host: localhost    Database: openxds
-- ------------------------------------------------------
-- Server version	5.7.18

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
-- Table structure for table `adhocquery`
--

DROP TABLE IF EXISTS `adhocquery`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `adhocquery` (
  `id` varchar(256) NOT NULL,
  `home` varchar(256) DEFAULT NULL,
  `lid` varchar(256) NOT NULL,
  `objectType` varchar(256) DEFAULT NULL,
  `status` varchar(256) NOT NULL,
  `versionName` varchar(16) DEFAULT NULL,
  `comment_` varchar(256) DEFAULT NULL,
  `queryLanguage` varchar(256) NOT NULL,
  `query` varchar(4096) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `lid_AdhQuery_idx` (`lid`),
  KEY `id_AdhQuery_idx` (`id`),
  KEY `home_AdhQuery_idx` (`home`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `affectedobject`
--

DROP TABLE IF EXISTS `affectedobject`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `affectedobject` (
  `id` varchar(256) NOT NULL,
  `home` varchar(256) DEFAULT NULL,
  `eventId` varchar(256) NOT NULL,
  PRIMARY KEY (`id`,`eventId`),
  KEY `id_evId_AFOBJ_idx` (`id`,`eventId`),
  KEY `id_AFOBJ_idx` (`id`),
  KEY `evid_AFOBJ_idx` (`eventId`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `association`
--

DROP TABLE IF EXISTS `association`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `association` (
  `id` varchar(256) NOT NULL,
  `home` varchar(256) DEFAULT NULL,
  `lid` varchar(256) NOT NULL,
  `objectType` varchar(256) DEFAULT NULL,
  `status` varchar(256) NOT NULL,
  `versionName` varchar(16) DEFAULT NULL,
  `comment_` varchar(256) DEFAULT NULL,
  `associationType` varchar(256) NOT NULL,
  `sourceObject` varchar(256) NOT NULL,
  `targetObject` varchar(256) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `lid_Assoc_idx` (`lid`),
  KEY `id_Assoc_idx` (`id`),
  KEY `home_Assoc_idx` (`home`),
  KEY `src_Ass_idx` (`sourceObject`),
  KEY `tgt_Ass_idx` (`targetObject`),
  KEY `type_Ass_idx` (`associationType`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `auditableevent`
--

DROP TABLE IF EXISTS `auditableevent`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `auditableevent` (
  `id` varchar(256) NOT NULL,
  `home` varchar(256) DEFAULT NULL,
  `lid` varchar(256) NOT NULL,
  `objectType` varchar(256) DEFAULT NULL,
  `status` varchar(256) NOT NULL,
  `versionName` varchar(16) DEFAULT NULL,
  `comment_` varchar(256) DEFAULT NULL,
  `requestId` varchar(256) NOT NULL,
  `eventType` varchar(256) NOT NULL,
  `timeStamp_` varchar(30) NOT NULL,
  `user_` varchar(256) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `lid_AUEVENT_idx` (`lid`),
  KEY `id_AUEVENT_idx` (`id`),
  KEY `home_AUEVENT_idx` (`home`),
  KEY `lid_AUEVENT_evtTyp` (`eventType`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `classification`
--

DROP TABLE IF EXISTS `classification`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `classification` (
  `id` varchar(256) NOT NULL,
  `home` varchar(256) DEFAULT NULL,
  `lid` varchar(256) NOT NULL,
  `objectType` varchar(256) DEFAULT NULL,
  `status` varchar(256) NOT NULL,
  `versionName` varchar(16) DEFAULT NULL,
  `comment_` varchar(256) DEFAULT NULL,
  `classificationNode` varchar(256) DEFAULT NULL,
  `classificationScheme` varchar(256) DEFAULT NULL,
  `classifiedObject` varchar(256) NOT NULL,
  `nodeRepresentation` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `lid_Class_idx` (`lid`),
  KEY `id_Class_idx` (`id`),
  KEY `home_Class_idx` (`home`),
  KEY `clsObj_Class_idx` (`classifiedObject`),
  KEY `node_Class_idx` (`classificationNode`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `classificationnode`
--

DROP TABLE IF EXISTS `classificationnode`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `classificationnode` (
  `id` varchar(256) NOT NULL,
  `home` varchar(256) DEFAULT NULL,
  `lid` varchar(256) NOT NULL,
  `objectType` varchar(256) DEFAULT NULL,
  `status` varchar(256) NOT NULL,
  `versionName` varchar(16) DEFAULT NULL,
  `comment_` varchar(256) DEFAULT NULL,
  `code` varchar(256) DEFAULT NULL,
  `parent` varchar(256) DEFAULT NULL,
  `path` varchar(1024) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `lid_Node_idx` (`lid`),
  KEY `id_Node_idx` (`id`),
  KEY `home_Node_idx` (`home`),
  KEY `parent_Node_idx` (`parent`),
  KEY `code_Node_idx` (`code`),
  KEY `path_Node_idx` (`path`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `classscheme`
--

DROP TABLE IF EXISTS `classscheme`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `classscheme` (
  `id` varchar(256) NOT NULL,
  `home` varchar(256) DEFAULT NULL,
  `lid` varchar(256) NOT NULL,
  `objectType` varchar(256) DEFAULT NULL,
  `status` varchar(256) NOT NULL,
  `versionName` varchar(16) DEFAULT NULL,
  `comment_` varchar(256) DEFAULT NULL,
  `isInternal` varchar(1) NOT NULL,
  `nodeType` varchar(256) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `lid_SCHEME_idx` (`lid`),
  KEY `id_SCHEME_idx` (`id`),
  KEY `home_SCHEME_idx` (`home`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `description`
--

DROP TABLE IF EXISTS `description`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `description` (
  `charset` varchar(32) DEFAULT NULL,
  `lang` varchar(32) NOT NULL,
  `value` varchar(1024) NOT NULL,
  `parent` varchar(256) NOT NULL,
  PRIMARY KEY (`parent`,`lang`),
  KEY `value_Desc_idx` (`value`),
  KEY `lngval_Desc_idx` (`lang`,`value`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `emailaddress`
--

DROP TABLE IF EXISTS `emailaddress`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `emailaddress` (
  `address` varchar(64) NOT NULL,
  `type` varchar(256) DEFAULT NULL,
  `parent` varchar(256) NOT NULL,
  KEY `parent_EmlAdr_idx` (`parent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `externalidentifier`
--

DROP TABLE IF EXISTS `externalidentifier`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `externalidentifier` (
  `id` varchar(256) NOT NULL,
  `home` varchar(256) DEFAULT NULL,
  `lid` varchar(256) NOT NULL,
  `objectType` varchar(256) DEFAULT NULL,
  `status` varchar(256) NOT NULL,
  `versionName` varchar(16) DEFAULT NULL,
  `comment_` varchar(256) DEFAULT NULL,
  `registryObject` varchar(256) NOT NULL,
  `identificationScheme` varchar(256) NOT NULL,
  `value` varchar(256) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `lid_EID_idx` (`lid`),
  KEY `id_EID_idx` (`id`),
  KEY `home_EID_idx` (`home`),
  KEY `ro_EID_idx` (`registryObject`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `externallink`
--

DROP TABLE IF EXISTS `externallink`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `externallink` (
  `id` varchar(256) NOT NULL,
  `home` varchar(256) DEFAULT NULL,
  `lid` varchar(256) NOT NULL,
  `objectType` varchar(256) DEFAULT NULL,
  `status` varchar(256) NOT NULL,
  `versionName` varchar(16) DEFAULT NULL,
  `comment_` varchar(256) DEFAULT NULL,
  `externalURI` varchar(256) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `lid_ExLink_idx` (`lid`),
  KEY `id_ExLink_idx` (`id`),
  KEY `home_ExLink_idx` (`home`),
  KEY `uri_ExLink_idx` (`externalURI`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `extrinsicobject`
--

DROP TABLE IF EXISTS `extrinsicobject`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `extrinsicobject` (
  `id` varchar(256) NOT NULL,
  `home` varchar(256) DEFAULT NULL,
  `lid` varchar(256) NOT NULL,
  `objectType` varchar(256) DEFAULT NULL,
  `status` varchar(256) NOT NULL,
  `versionName` varchar(16) DEFAULT NULL,
  `comment_` varchar(256) DEFAULT NULL,
  `isOpaque` varchar(1) NOT NULL,
  `mimeType` varchar(256) DEFAULT NULL,
  `contentVersionName` varchar(16) DEFAULT NULL,
  `contentVersionComment` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `lid_EXTOBJ_idx` (`lid`),
  KEY `id_EXTOBJ_idx` (`id`),
  KEY `home_EXTOBJ_idx` (`home`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `federation`
--

DROP TABLE IF EXISTS `federation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `federation` (
  `id` varchar(256) NOT NULL,
  `home` varchar(256) DEFAULT NULL,
  `lid` varchar(256) NOT NULL,
  `objectType` varchar(256) DEFAULT NULL,
  `status` varchar(256) NOT NULL,
  `versionName` varchar(16) DEFAULT NULL,
  `comment_` varchar(256) DEFAULT NULL,
  `replicationSyncLatency` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `lid_FED_idx` (`lid`),
  KEY `id_FED_idx` (`id`),
  KEY `home_FED_idx` (`home`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Temporary table structure for view `identifiable`
--

DROP TABLE IF EXISTS `identifiable`;
/*!50001 DROP VIEW IF EXISTS `identifiable`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE VIEW `identifiable` AS SELECT 
 1 AS `id`,
 1 AS `home`*/;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `name_`
--

DROP TABLE IF EXISTS `name_`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `name_` (
  `charset` varchar(32) DEFAULT NULL,
  `lang` varchar(32) NOT NULL,
  `value` varchar(1024) NOT NULL,
  `parent` varchar(256) NOT NULL,
  PRIMARY KEY (`parent`,`lang`),
  KEY `value_Name_idx` (`value`),
  KEY `lngval_Name_idx` (`lang`,`value`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `notification`
--

DROP TABLE IF EXISTS `notification`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `notification` (
  `id` varchar(256) NOT NULL,
  `home` varchar(256) DEFAULT NULL,
  `lid` varchar(256) NOT NULL,
  `objectType` varchar(256) DEFAULT NULL,
  `status` varchar(256) NOT NULL,
  `versionName` varchar(16) DEFAULT NULL,
  `comment_` varchar(256) DEFAULT NULL,
  `subscription` varchar(256) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `notificationobject`
--

DROP TABLE IF EXISTS `notificationobject`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `notificationobject` (
  `notificationId` varchar(256) NOT NULL,
  `registryObjectId` varchar(256) NOT NULL,
  PRIMARY KEY (`notificationId`,`registryObjectId`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `notifyaction`
--

DROP TABLE IF EXISTS `notifyaction`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `notifyaction` (
  `notificationOption` varchar(256) NOT NULL,
  `endPoint` varchar(256) NOT NULL,
  `parent` varchar(256) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `objectref`
--

DROP TABLE IF EXISTS `objectref`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `objectref` (
  `id` varchar(256) NOT NULL,
  `home` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_ObjectRef_idx` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `organization`
--

DROP TABLE IF EXISTS `organization`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `organization` (
  `id` varchar(256) NOT NULL,
  `home` varchar(256) DEFAULT NULL,
  `lid` varchar(256) NOT NULL,
  `objectType` varchar(256) DEFAULT NULL,
  `status` varchar(256) NOT NULL,
  `versionName` varchar(16) DEFAULT NULL,
  `comment_` varchar(256) DEFAULT NULL,
  `parent` varchar(256) DEFAULT NULL,
  `primaryContact` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `lid_ORG_idx` (`lid`),
  KEY `id_ORG_idx` (`id`),
  KEY `home_ORG_idx` (`home`),
  KEY `parent_ORG_idx` (`parent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `person`
--

DROP TABLE IF EXISTS `person`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `person` (
  `id` varchar(256) NOT NULL,
  `home` varchar(256) DEFAULT NULL,
  `lid` varchar(256) NOT NULL,
  `objectType` varchar(256) DEFAULT NULL,
  `status` varchar(256) NOT NULL,
  `versionName` varchar(16) DEFAULT NULL,
  `comment_` varchar(256) DEFAULT NULL,
  `personName_firstName` varchar(64) DEFAULT NULL,
  `personName_middleName` varchar(64) DEFAULT NULL,
  `personName_lastName` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `lid_Person_idx` (`lid`),
  KEY `id_Person_idx` (`id`),
  KEY `home_Person_idx` (`home`),
  KEY `lastNm_Person_idx` (`personName_lastName`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `personidentifier`
--

DROP TABLE IF EXISTS `personidentifier`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `personidentifier` (
  `registrypatientid` varchar(255) NOT NULL,
  `assigningauthority` varchar(255) NOT NULL,
  `patientid` varchar(255) NOT NULL,
  `deleted` varchar(1) DEFAULT NULL,
  `merged` varchar(1) DEFAULT NULL,
  `survivingpatientid` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`registrypatientid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `postaladdress`
--

DROP TABLE IF EXISTS `postaladdress`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `postaladdress` (
  `city` varchar(64) DEFAULT NULL,
  `country` varchar(64) DEFAULT NULL,
  `postalCode` varchar(64) DEFAULT NULL,
  `state` varchar(64) DEFAULT NULL,
  `street` varchar(64) DEFAULT NULL,
  `streetNumber` varchar(32) DEFAULT NULL,
  `parent` varchar(256) NOT NULL,
  KEY `parent_PstlAdr_idx` (`parent`),
  KEY `city_PstlAdr_idx` (`city`),
  KEY `cntry_PstlAdr_idx` (`country`),
  KEY `pCode_PstlAdr_idx` (`postalCode`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `registry`
--

DROP TABLE IF EXISTS `registry`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `registry` (
  `id` varchar(256) NOT NULL,
  `home` varchar(256) DEFAULT NULL,
  `lid` varchar(256) NOT NULL,
  `objectType` varchar(256) DEFAULT NULL,
  `status` varchar(256) NOT NULL,
  `versionName` varchar(16) DEFAULT NULL,
  `comment_` varchar(256) DEFAULT NULL,
  `catalogingSyncLatency` varchar(32) DEFAULT 'P1D',
  `conformanceProfile` varchar(16) DEFAULT NULL,
  `operator` varchar(256) NOT NULL,
  `replicationSyncLatency` varchar(32) DEFAULT 'P1D',
  `specificationVersion` varchar(8) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `lid_Registry_idx` (`lid`),
  KEY `id_Registry_idx` (`id`),
  KEY `home_Registry_idx` (`home`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Temporary table structure for view `registryobject`
--

DROP TABLE IF EXISTS `registryobject`;
/*!50001 DROP VIEW IF EXISTS `registryobject`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE VIEW `registryobject` AS SELECT 
 1 AS `id`,
 1 AS `home`,
 1 AS `lid`,
 1 AS `objectType`,
 1 AS `status`,
 1 AS `versionName`,
 1 AS `comment_`*/;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `registrypackage`
--

DROP TABLE IF EXISTS `registrypackage`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `registrypackage` (
  `id` varchar(256) NOT NULL,
  `home` varchar(256) DEFAULT NULL,
  `lid` varchar(256) NOT NULL,
  `objectType` varchar(256) DEFAULT NULL,
  `status` varchar(256) NOT NULL,
  `versionName` varchar(16) DEFAULT NULL,
  `comment_` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `lid_PKG_idx` (`lid`),
  KEY `id_PKG_idx` (`id`),
  KEY `home_PKG_idx` (`home`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `repository`
--

DROP TABLE IF EXISTS `repository`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `repository` (
  `documentuniqueid` varchar(255) NOT NULL,
  `mimetype` varchar(255) NOT NULL,
  `documentsize` int(11) NOT NULL,
  `hash` varchar(255) NOT NULL,
  `content` mediumblob,
  PRIMARY KEY (`documentuniqueid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `repositoryitem`
--

DROP TABLE IF EXISTS `repositoryitem`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `repositoryitem` (
  `lid` varchar(256) NOT NULL,
  `versionname` varchar(16) NOT NULL,
  `content` mediumblob,
  PRIMARY KEY (`lid`,`versionname`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `service`
--

DROP TABLE IF EXISTS `service`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `service` (
  `id` varchar(256) NOT NULL,
  `home` varchar(256) DEFAULT NULL,
  `lid` varchar(256) NOT NULL,
  `objectType` varchar(256) DEFAULT NULL,
  `status` varchar(256) NOT NULL,
  `versionName` varchar(16) DEFAULT NULL,
  `comment_` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `lid_Service_idx` (`lid`),
  KEY `id_Service_idx` (`id`),
  KEY `home_Service_idx` (`home`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `servicebinding`
--

DROP TABLE IF EXISTS `servicebinding`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `servicebinding` (
  `id` varchar(256) NOT NULL,
  `home` varchar(256) DEFAULT NULL,
  `lid` varchar(256) NOT NULL,
  `objectType` varchar(256) DEFAULT NULL,
  `status` varchar(256) NOT NULL,
  `versionName` varchar(16) DEFAULT NULL,
  `comment_` varchar(256) DEFAULT NULL,
  `service` varchar(256) NOT NULL,
  `accessURI` varchar(256) DEFAULT NULL,
  `targetBinding` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `lid_BIND_idx` (`lid`),
  KEY `id_BIND_idx` (`id`),
  KEY `home_BIND_idx` (`home`),
  KEY `service_BIND_idx` (`service`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `slot`
--

DROP TABLE IF EXISTS `slot`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `slot` (
  `sequenceId` int(11) NOT NULL,
  `name_` varchar(256) NOT NULL,
  `slotType` varchar(256) DEFAULT NULL,
  `value` varchar(256) DEFAULT NULL,
  `parent` varchar(256) NOT NULL,
  PRIMARY KEY (`parent`,`name_`,`sequenceId`),
  KEY `parent_Slot_idx` (`parent`),
  KEY `name_Slot_idx` (`name_`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `specificationlink`
--

DROP TABLE IF EXISTS `specificationlink`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `specificationlink` (
  `id` varchar(256) NOT NULL,
  `home` varchar(256) DEFAULT NULL,
  `lid` varchar(256) NOT NULL,
  `objectType` varchar(256) DEFAULT NULL,
  `status` varchar(256) NOT NULL,
  `versionName` varchar(16) DEFAULT NULL,
  `comment_` varchar(256) DEFAULT NULL,
  `serviceBinding` varchar(256) NOT NULL,
  `specificationObject` varchar(256) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `lid_SLnk_idx` (`lid`),
  KEY `id_SLnk_idx` (`id`),
  KEY `home_SLnk_idx` (`home`),
  KEY `binding_SLnk_idx` (`serviceBinding`),
  KEY `spec_SLnk_idx` (`specificationObject`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `subscription`
--

DROP TABLE IF EXISTS `subscription`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `subscription` (
  `id` varchar(256) NOT NULL,
  `home` varchar(256) DEFAULT NULL,
  `lid` varchar(256) NOT NULL,
  `objectType` varchar(256) DEFAULT NULL,
  `status` varchar(256) NOT NULL,
  `versionName` varchar(16) DEFAULT NULL,
  `comment_` varchar(256) DEFAULT NULL,
  `selector` varchar(256) NOT NULL,
  `endTime` varchar(30) DEFAULT NULL,
  `notificationInterval` varchar(32) DEFAULT 'P1D',
  `startTime` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `lid_SUBS_idx` (`lid`),
  KEY `id_SUBS_idx` (`id`),
  KEY `home_SUBS_idx` (`home`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `telephonenumber`
--

DROP TABLE IF EXISTS `telephonenumber`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `telephonenumber` (
  `areaCode` varchar(8) DEFAULT NULL,
  `countryCode` varchar(8) DEFAULT NULL,
  `extension` varchar(8) DEFAULT NULL,
  `number_` varchar(16) DEFAULT NULL,
  `phoneType` varchar(256) DEFAULT NULL,
  `parent` varchar(256) NOT NULL,
  KEY `parent_Phone_idx` (`parent`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `usagedescription`
--

DROP TABLE IF EXISTS `usagedescription`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `usagedescription` (
  `charset` varchar(32) DEFAULT NULL,
  `lang` varchar(32) NOT NULL,
  `value` varchar(1024) NOT NULL,
  `parent` varchar(256) NOT NULL,
  PRIMARY KEY (`parent`,`lang`),
  KEY `value_UsgDes_idx` (`value`),
  KEY `lngval_UsgDes_idx` (`lang`,`value`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `usageparameter`
--

DROP TABLE IF EXISTS `usageparameter`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `usageparameter` (
  `value` varchar(1024) NOT NULL,
  `parent` varchar(256) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user_`
--

DROP TABLE IF EXISTS `user_`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_` (
  `id` varchar(256) NOT NULL,
  `home` varchar(256) DEFAULT NULL,
  `lid` varchar(256) NOT NULL,
  `objectType` varchar(256) DEFAULT NULL,
  `status` varchar(256) NOT NULL,
  `versionName` varchar(16) DEFAULT NULL,
  `comment_` varchar(256) DEFAULT NULL,
  `personName_firstName` varchar(64) DEFAULT NULL,
  `personName_middleName` varchar(64) DEFAULT NULL,
  `personName_lastName` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `lid_User_idx` (`lid`),
  KEY `id_User_idx` (`id`),
  KEY `home_User_idx` (`home`),
  KEY `lastNm_User_idx` (`personName_lastName`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Final view structure for view `identifiable`
--

/*!50001 DROP VIEW IF EXISTS `identifiable`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = latin1 */;
/*!50001 SET character_set_results     = latin1 */;
/*!50001 SET collation_connection      = latin1_swedish_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `identifiable` AS select `adhocquery`.`id` AS `id`,`adhocquery`.`home` AS `home` from `adhocquery` union all select `association`.`id` AS `id`,`association`.`home` AS `home` from `association` union all select `auditableevent`.`id` AS `id`,`auditableevent`.`home` AS `home` from `auditableevent` union all select `classification`.`id` AS `id`,`classification`.`home` AS `home` from `classification` union all select `classificationnode`.`id` AS `id`,`classificationnode`.`home` AS `home` from `classificationnode` union all select `classscheme`.`id` AS `id`,`classscheme`.`home` AS `home` from `classscheme` union all select `externalidentifier`.`id` AS `id`,`externalidentifier`.`home` AS `home` from `externalidentifier` union all select `externallink`.`id` AS `id`,`externallink`.`home` AS `home` from `externallink` union all select `extrinsicobject`.`id` AS `id`,`extrinsicobject`.`home` AS `home` from `extrinsicobject` union all select `federation`.`id` AS `id`,`federation`.`home` AS `home` from `federation` union all select `organization`.`id` AS `id`,`organization`.`home` AS `home` from `organization` union all select `registry`.`id` AS `id`,`registry`.`home` AS `home` from `registry` union all select `registrypackage`.`id` AS `id`,`registrypackage`.`home` AS `home` from `registrypackage` union all select `service`.`id` AS `id`,`service`.`home` AS `home` from `service` union all select `servicebinding`.`id` AS `id`,`servicebinding`.`home` AS `home` from `servicebinding` union all select `specificationlink`.`id` AS `id`,`specificationlink`.`home` AS `home` from `specificationlink` union all select `subscription`.`id` AS `id`,`subscription`.`home` AS `home` from `subscription` union all select `user_`.`id` AS `id`,`user_`.`home` AS `home` from `user_` union all select `person`.`id` AS `id`,`person`.`home` AS `home` from `person` union all select `objectref`.`id` AS `id`,`objectref`.`home` AS `home` from `objectref` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `registryobject`
--

/*!50001 DROP VIEW IF EXISTS `registryobject`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = latin1 */;
/*!50001 SET character_set_results     = latin1 */;
/*!50001 SET collation_connection      = latin1_swedish_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `registryobject` AS select `adhocquery`.`id` AS `id`,`adhocquery`.`home` AS `home`,`adhocquery`.`lid` AS `lid`,`adhocquery`.`objectType` AS `objectType`,`adhocquery`.`status` AS `status`,`adhocquery`.`versionName` AS `versionName`,`adhocquery`.`comment_` AS `comment_` from `adhocquery` union all select `association`.`id` AS `id`,`association`.`home` AS `home`,`association`.`lid` AS `lid`,`association`.`objectType` AS `objectType`,`association`.`status` AS `status`,`association`.`versionName` AS `versionName`,`association`.`comment_` AS `comment_` from `association` union all select `auditableevent`.`id` AS `id`,`auditableevent`.`home` AS `home`,`auditableevent`.`lid` AS `lid`,`auditableevent`.`objectType` AS `objectType`,`auditableevent`.`status` AS `status`,`auditableevent`.`versionName` AS `versionName`,`auditableevent`.`comment_` AS `comment_` from `auditableevent` union all select `classification`.`id` AS `id`,`classification`.`home` AS `home`,`classification`.`lid` AS `lid`,`classification`.`objectType` AS `objectType`,`classification`.`status` AS `status`,`classification`.`versionName` AS `versionName`,`classification`.`comment_` AS `comment_` from `classification` union all select `classificationnode`.`id` AS `id`,`classificationnode`.`home` AS `home`,`classificationnode`.`lid` AS `lid`,`classificationnode`.`objectType` AS `objectType`,`classificationnode`.`status` AS `status`,`classificationnode`.`versionName` AS `versionName`,`classificationnode`.`comment_` AS `comment_` from `classificationnode` union all select `classscheme`.`id` AS `id`,`classscheme`.`home` AS `home`,`classscheme`.`lid` AS `lid`,`classscheme`.`objectType` AS `objectType`,`classscheme`.`status` AS `status`,`classscheme`.`versionName` AS `versionName`,`classscheme`.`comment_` AS `comment_` from `classscheme` union all select `externalidentifier`.`id` AS `id`,`externalidentifier`.`home` AS `home`,`externalidentifier`.`lid` AS `lid`,`externalidentifier`.`objectType` AS `objectType`,`externalidentifier`.`status` AS `status`,`externalidentifier`.`versionName` AS `versionName`,`externalidentifier`.`comment_` AS `comment_` from `externalidentifier` union all select `externallink`.`id` AS `id`,`externallink`.`home` AS `home`,`externallink`.`lid` AS `lid`,`externallink`.`objectType` AS `objectType`,`externallink`.`status` AS `status`,`externallink`.`versionName` AS `versionName`,`externallink`.`comment_` AS `comment_` from `externallink` union all select `extrinsicobject`.`id` AS `id`,`extrinsicobject`.`home` AS `home`,`extrinsicobject`.`lid` AS `lid`,`extrinsicobject`.`objectType` AS `objectType`,`extrinsicobject`.`status` AS `status`,`extrinsicobject`.`versionName` AS `versionName`,`extrinsicobject`.`comment_` AS `comment_` from `extrinsicobject` union all select `federation`.`id` AS `id`,`federation`.`home` AS `home`,`federation`.`lid` AS `lid`,`federation`.`objectType` AS `objectType`,`federation`.`status` AS `status`,`federation`.`versionName` AS `versionName`,`federation`.`comment_` AS `comment_` from `federation` union all select `organization`.`id` AS `id`,`organization`.`home` AS `home`,`organization`.`lid` AS `lid`,`organization`.`objectType` AS `objectType`,`organization`.`status` AS `status`,`organization`.`versionName` AS `versionName`,`organization`.`comment_` AS `comment_` from `organization` union all select `registry`.`id` AS `id`,`registry`.`home` AS `home`,`registry`.`lid` AS `lid`,`registry`.`objectType` AS `objectType`,`registry`.`status` AS `status`,`registry`.`versionName` AS `versionName`,`registry`.`comment_` AS `comment_` from `registry` union all select `registrypackage`.`id` AS `id`,`registrypackage`.`home` AS `home`,`registrypackage`.`lid` AS `lid`,`registrypackage`.`objectType` AS `objectType`,`registrypackage`.`status` AS `status`,`registrypackage`.`versionName` AS `versionName`,`registrypackage`.`comment_` AS `comment_` from `registrypackage` union all select `service`.`id` AS `id`,`service`.`home` AS `home`,`service`.`lid` AS `lid`,`service`.`objectType` AS `objectType`,`service`.`status` AS `status`,`service`.`versionName` AS `versionName`,`service`.`comment_` AS `comment_` from `service` union all select `servicebinding`.`id` AS `id`,`servicebinding`.`home` AS `home`,`servicebinding`.`lid` AS `lid`,`servicebinding`.`objectType` AS `objectType`,`servicebinding`.`status` AS `status`,`servicebinding`.`versionName` AS `versionName`,`servicebinding`.`comment_` AS `comment_` from `servicebinding` union all select `specificationlink`.`id` AS `id`,`specificationlink`.`home` AS `home`,`specificationlink`.`lid` AS `lid`,`specificationlink`.`objectType` AS `objectType`,`specificationlink`.`status` AS `status`,`specificationlink`.`versionName` AS `versionName`,`specificationlink`.`comment_` AS `comment_` from `specificationlink` union all select `subscription`.`id` AS `id`,`subscription`.`home` AS `home`,`subscription`.`lid` AS `lid`,`subscription`.`objectType` AS `objectType`,`subscription`.`status` AS `status`,`subscription`.`versionName` AS `versionName`,`subscription`.`comment_` AS `comment_` from `subscription` union all select `user_`.`id` AS `id`,`user_`.`home` AS `home`,`user_`.`lid` AS `lid`,`user_`.`objectType` AS `objectType`,`user_`.`status` AS `status`,`user_`.`versionName` AS `versionName`,`user_`.`comment_` AS `comment_` from `user_` union all select `person`.`id` AS `id`,`person`.`home` AS `home`,`person`.`lid` AS `lid`,`person`.`objectType` AS `objectType`,`person`.`status` AS `status`,`person`.`versionName` AS `versionName`,`person`.`comment_` AS `comment_` from `person` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-08-20 19:59:29

INSERT INTO `user_` VALUES('urn:freebxml:registry:predefinedusers:registryoperator', null, 'urn:freebxml:registry:predefinedUser:RegistryOperator', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Person:User', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.0', null,  'Registry', null, 'Operator');
INSERT INTO PostalAddress VALUES('Burlington', 'USA', '01803', 'MA', 'Network Dr.', '1', 'urn:freebxml:registry:predefinedusers:registryoperator');
INSERT INTO TelephoneNumber VALUES('781', '01', '', '442-0703', 'OfficePhone', 'urn:freebxml:registry:predefinedusers:registryoperator');
INSERT INTO EmailAddress VALUES('registryOperator@ebxmlrr.com', 'OfficeEmail', 'urn:freebxml:registry:predefinedusers:registryoperator');
-- We make the owner of the User object to be the user itself.
INSERT INTO AuditableEvent VALUES('urn:freebxml:registry:predefinedEvent:createRegistryOperator', null, 'urn:freebxml:registry:predefinedEvent:createRegistryOperator', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:AuditableEvent', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.0', null,  '//TODO', 'urn:oasis:names:tc:ebxml-regrep:EventType:Created', '2003-12-10 10:53:24', 'urn:freebxml:registry:predefinedusers:registryoperator');
INSERT INTO AffectedObject VALUES('urn:freebxml:registry:predefinedusers:registryoperator', null, 'urn:freebxml:registry:predefinedEvent:createRegistryOperator');

INSERT INTO `user_` VALUES('urn:freebxml:registry:predefinedusers:registryguest', null, 'urn:freebxml:registry:predefinedusers:registryguest', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Person:User', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted',   '1.0', null,  'Registry', null, 'Guest');
INSERT INTO PostalAddress VALUES('Burlington', 'USA', '01803', 'MA', 'Network Dr.', '1', 'urn:freebxml:registry:predefinedusers:registryguest');
INSERT INTO TelephoneNumber VALUES('781', '01', '', '442-0703', 'OfficePhone', 'urn:freebxml:registry:predefinedusers:registryguest');
INSERT INTO EmailAddress VALUES('registryGuest@ebxmlrr.com', 'OfficeEmail', 'urn:freebxml:registry:predefinedusers:registryguest');
-- We make the owner of the User object to be the user itself.
INSERT INTO AuditableEvent VALUES('urn:freebxml:registry:predefinedEvent:createRegistryGuest', null, 'urn:freebxml:registry:predefinedEvent:createRegistryGuest', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:AuditableEvent', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.0', null,  '//TODO', 'urn:oasis:names:tc:ebxml-regrep:EventType:Created', '2003-12-10 10:53:24', 'urn:freebxml:registry:predefinedusers:registryguest');
INSERT INTO AffectedObject VALUES('urn:freebxml:registry:predefinedusers:registryguest', null, 'urn:freebxml:registry:predefinedEvent:createRegistryGuest');

INSERT INTO `user_` VALUES('urn:freebxml:registry:predefinedusers:farrukh', null, 'urn:freebxml:registry:predefinedusers:farrukh', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Person:User', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted',  '1.0', null,  'Farrukh', 'Salahudin', 'Najmi');
INSERT INTO PostalAddress VALUES('Burlington', 'USA', '01803', 'MA', 'Network Dr.', '1', 'urn:freebxml:registry:predefinedusers:farrukh');
INSERT INTO TelephoneNumber VALUES('781', '01', '', '442-0703', 'OfficePhone', 'urn:freebxml:registry:predefinedusers:farrukh');
INSERT INTO EmailAddress VALUES('farrukh.najmi@sun.com', 'OfficeEmail', 'urn:freebxml:registry:predefinedusers:farrukh');
-- We make the owner of the User object to be the user itself.
INSERT INTO AuditableEvent VALUES('urn:freebxml:registry:predefinedEvent:createFarrukh', null, 'urn:freebxml:registry:predefinedEvent:createFarrukh', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:AuditableEvent', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.0', null,  '//TODO', 'urn:oasis:names:tc:ebxml-regrep:EventType:Created', '2003-12-10 10:53:24', 'urn:freebxml:registry:predefinedusers:farrukh');
INSERT INTO AffectedObject VALUES('urn:freebxml:registry:predefinedusers:farrukh', null, 'urn:freebxml:registry:predefinedEvent:createFarrukh');

INSERT INTO `user_` VALUES('urn:freebxml:registry:predefinedusers:nikola', null, 'urn:freebxml:registry:predefinedusers:nikola', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Person:User', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted',  '1.0', null,  'Nikola', null, 'Stojanovic');
INSERT INTO PostalAddress VALUES('Ithaca', 'France', null, 'NY', 'Terazije', '19', 'urn:freebxml:registry:predefinedusers:nikola');
INSERT INTO TelephoneNumber VALUES('11', '381', '', '222-2222', 'OfficePhone', 'urn:freebxml:registry:predefinedusers:nikola');
INSERT INTO TelephoneNumber VALUES('11', '381', '', '444-4444', 'HomePhone', 'urn:freebxml:registry:predefinedusers:nikola');
INSERT INTO EmailAddress VALUES('nhomest1@twcny.rr.com', 'OfficeEmail', 'urn:freebxml:registry:predefinedusers:nikola');
-- We make the owner of the User object to be the user itself.
INSERT INTO AuditableEvent VALUES('urn:freebxml:registry:predefinedEvent:createNikola', null, 'urn:freebxml:registry:predefinedEvent:createNikola', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:AuditableEvent', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.0', null,  '//TODO', 'urn:oasis:names:tc:ebxml-regrep:EventType:Created', '2003-12-10 10:53:24', 'urn:freebxml:registry:predefinedusers:nikola');
INSERT INTO AffectedObject VALUES('urn:freebxml:registry:predefinedusers:nikola', null, 'urn:freebxml:registry:predefinedEvent:createNikola');

-- Insert IHE defined clssifications

-- XDSSubmissionSet
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:uuid:a54d6aa5-d40d-43f9-88c5-b4633d873bdd', NULL, 'urn:uuid:a54d6aa5-d40d-43f9-88c5-b4633d873bdd', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', NULL, 'XDSSubmissionSet', NULL, NULL);
INSERT INTO CLASSSCHEME VALUES ('urn:uuid:a7058bb9-b4e4-4307-ba5b-e3f0ab85e12d', NULL, 'urn:uuid:a7058bb9-b4e4-4307-ba5b-e3f0ab85e12d', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', 'XDSSubmissionSet.author', 'F', 'urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSSCHEME VALUES ('urn:uuid:aa543740-bdda-424e-8c96-df4873be8500', NULL, 'urn:uuid:aa543740-bdda-424e-8c96-df4873be8500', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', 'XDSSubmissionSet.contentTypeCode', 'F', 'urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSSCHEME VALUES ('urn:uuid:6b5aea1a-874d-4603-a4bc-96a0a7b38446', NULL, 'urn:uuid:6b5aea1a-874d-4603-a4bc-96a0a7b38446', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', 'XDSSubmissionSet.patientId', 'F', 'urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSSCHEME VALUES ('urn:uuid:554ac39e-e3fe-47fe-b233-965d2a147832', NULL, 'urn:uuid:554ac39e-e3fe-47fe-b233-965d2a147832', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', 'XDSSubmissionSet.sourceId', 'F', 'urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSSCHEME VALUES ('urn:uuid:96fdda7c-d067-4183-912e-bf5ee74998a8', NULL, 'urn:uuid:96fdda7c-d067-4183-912e-bf5ee74998a8', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', 'XDSSubmissionSet.uniqueId', 'F', 'urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');

-- XDSDocumentEntry
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:uuid:7edca82f-054d-47f2-a032-9b2a5b5186c1', NULL, 'urn:uuid:7edca82f-054d-47f2-a032-9b2a5b5186c1', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', NULL, 'XDSDocumentEntry', NULL, NULL);
INSERT INTO CLASSSCHEME VALUES ('urn:uuid:93606bcf-9494-43ec-9b4e-a7748d1a838d', NULL, 'urn:uuid:93606bcf-9494-43ec-9b4e-a7748d1a838d', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', 'XDSDocumentEntry.author', 'F', 'urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSSCHEME VALUES ('urn:uuid:41a5887f-8865-4c09-adf7-e362475b143a', NULL, 'urn:uuid:41a5887f-8865-4c09-adf7-e362475b143a', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', 'XDSDocumentEntry.classCode', 'F', 'urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSSCHEME VALUES ('urn:uuid:f4f85eac-e6cb-4883-b524-f2705394840f', NULL, 'urn:uuid:f4f85eac-e6cb-4883-b524-f2705394840f', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', 'XDSDocumentEntry.confidentialityCode', 'F', 'urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSSCHEME VALUES ('urn:uuid:2c6b8cb7-8b2a-4051-b291-b1ae6a575ef4', NULL, 'urn:uuid:2c6b8cb7-8b2a-4051-b291-b1ae6a575ef4', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', 'XDSDocumentEntry.eventCodeList', 'F', 'urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSSCHEME VALUES ('urn:uuid:a09d5840-386c-46f2-b5ad-9c3699a4309d', NULL, 'urn:uuid:a09d5840-386c-46f2-b5ad-9c3699a4309d', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', 'XDSDocumentEntry.formatCode', 'F', 'urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSSCHEME VALUES ('urn:uuid:f33fb8ac-18af-42cc-ae0e-ed0b0bdb91e1', NULL, 'urn:uuid:f33fb8ac-18af-42cc-ae0e-ed0b0bdb91e1', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', 'XDSDocumentEntry.healthCareFacilityTypeCode', 'F', 'urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSSCHEME VALUES ('urn:uuid:58a6f841-87b3-4a3e-92fd-a8ffeff98427', NULL, 'urn:uuid:58a6f841-87b3-4a3e-92fd-a8ffeff98427', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', 'XDSDocumentEntry.patientId', 'F', 'urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSSCHEME VALUES ('urn:uuid:cccf5598-8b07-4b77-a05e-ae952c785ead', NULL, 'urn:uuid:cccf5598-8b07-4b77-a05e-ae952c785ead', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', 'XDSDocumentEntry.practiceSettingCode', 'F', 'urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSSCHEME VALUES ('urn:uuid:f0306f51-975f-434e-a61c-c59651d33983', NULL, 'urn:uuid:f0306f51-975f-434e-a61c-c59651d33983', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', 'XDSDocumentEntry.typeCode', 'F', 'urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSSCHEME VALUES ('urn:uuid:2e82c1f6-a085-4c72-9da3-8640a32e42ab', NULL, 'urn:uuid:2e82c1f6-a085-4c72-9da3-8640a32e42ab', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', 'XDSDocumentEntry.uniqueId', 'F', 'urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');

-- XDSFolder
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:uuid:d9d542f3-6cc4-48b6-8870-ea235fbc94c2', NULL, 'urn:uuid:d9d542f3-6cc4-48b6-8870-ea235fbc94c2', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', NULL, 'XDSFolder', NULL, NULL);
INSERT INTO CLASSSCHEME VALUES ('urn:uuid:1ba97051-7806-41a8-a48b-8fce7af683c5', NULL, 'urn:uuid:1ba97051-7806-41a8-a48b-8fce7af683c5', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', 'XDSFolder.codeList', 'F', 'urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSSCHEME VALUES ('urn:uuid:f64ffdf0-4b97-4e06-b79f-a52b38ec2f8a', NULL, 'urn:uuid:f64ffdf0-4b97-4e06-b79f-a52b38ec2f8a', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', 'XDSFolder.patientId', 'F', 'urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSSCHEME VALUES ('urn:uuid:75df8f67-9973-4fbe-a900-df66cefecc5a', NULL, 'urn:uuid:75df8f67-9973-4fbe-a900-df66cefecc5a', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', 'XDSFolder.uniqueId', 'F', 'urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');

-- Document Relationships
INSERT INTO CLASSSCHEME VALUES ('urn:oasis:names:tc:ebxml-regrep:classificationScheme:AssociationType',NULL,'urn:oasis:names:tc:ebxml-regrep:classificationScheme:AssociationType','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1','AssociationType','F','urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:AssociationType:HasMember', NULL, 'urn:oasis:names:tc:ebxml-regrep:AssociationType:HasMember', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', NULL, 'HasMember', 'urn:oasis:names:tc:ebxml-regrep:classificationScheme:AssociationType', '/urn:oasis:names:tc:ebxml-regrep:classificationScheme:AssociationType/HasMember');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:ihe:iti:2007:AssociationType:APND', NULL, 'urn:ihe:iti:2007:AssociationType:APND', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', NULL, 'APND', 'urn:oasis:names:tc:ebxml-regrep:classificationScheme:AssociationType', '/urn:oasis:names:tc:ebxml-regrep:classificationScheme:AssociationType/APND');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:ihe:iti:2007:AssociationType:RPLC', NULL, 'urn:ihe:iti:2007:AssociationType:RPLC', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', NULL, 'RPLC', 'urn:oasis:names:tc:ebxml-regrep:classificationScheme:AssociationType', '/urn:oasis:names:tc:ebxml-regrep:classificationScheme:AssociationType/RPLC');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:ihe:iti:2007:AssociationType:XFRM', NULL, 'urn:ihe:iti:2007:AssociationType:XFRM', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', NULL, 'XFRM', 'urn:oasis:names:tc:ebxml-regrep:classificationScheme:AssociationType', '/urn:oasis:names:tc:ebxml-regrep:classificationScheme:AssociationType/XFRM');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:ihe:iti:2007:AssociationType:XFRM_RPLC', NULL, 'urn:ihe:iti:2007:AssociationType:XFRM_RPLC', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', NULL, 'XFRM_RPLC', 'urn:oasis:names:tc:ebxml-regrep:classificationScheme:AssociationType', '/urn:oasis:names:tc:ebxml-regrep:classificationScheme:AssociationType/XFRM_RPLC');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:ihe:iti:2007:AssociationType:signs', NULL, 'urn:ihe:iti:2007:AssociationType:signs', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', NULL, 'signs', 'urn:oasis:names:tc:ebxml-regrep:classificationScheme:AssociationType', '/urn:oasis:names:tc:ebxml-regrep:classificationScheme:AssociationType/signs');

-- Others
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:uuid:10aa1a4b-715a-4120-bfd0-9760414112c8', NULL, 'urn:uuid:10aa1a4b-715a-4120-bfd0-9760414112c8', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', NULL, 'XDSDocumentEntryStub', NULL, NULL);
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:uuid:abd807a3-4432-4053-87b4-fd82c643d1f3', NULL, 'urn:uuid:abd807a3-4432-4053-87b4-fd82c643d1f3', 'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode', 'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted', '1.1', NULL, 'Association Documentation', NULL, NULL);

-- Insert Factory defined classifications
-- StatusTypes
INSERT INTO CLASSSCHEME VALUES ('urn:oasis:names:tc:ebxml-regrep:classificationScheme:StatusType',NULL,'urn:oasis:names:tc:ebxml-regrep:classificationScheme:StatusType','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1','StatusType','T','urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:StatusType:Approved',NULL,'urn:oasis:names:tc:ebxml-regrep:StatusType:Approved','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Approved','urn:oasis:names:tc:ebxml-regrep:classificationScheme:StatusType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:StatusType/Approved');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted',NULL,'urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Approved','urn:oasis:names:tc:ebxml-regrep:classificationScheme:StatusType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:StatusType/Submitted');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:StatusType:Deprecated',NULL,'urn:oasis:names:tc:ebxml-regrep:StatusType:Deprecated','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Approved','urn:oasis:names:tc:ebxml-regrep:classificationScheme:StatusType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:StatusType/Deprecated');

-- DataTypes
INSERT INTO CLASSSCHEME VALUES ('urn:oasis:names:tc:ebxml-regrep:classificationScheme:DataType',NULL,'urn:oasis:names:tc:ebxml-regrep:classificationScheme:DataType','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1','DataType','T','urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:DataType:Boolean',NULL,'urn:oasis:names:tc:ebxml-regrep:DataType:Boolean','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Boolean','urn:oasis:names:tc:ebxml-regrep:classificationScheme:DataType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:DataType/Boolean');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:DataType:Date',NULL,'urn:oasis:names:tc:ebxml-regrep:DataType:Date','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Date','urn:oasis:names:tc:ebxml-regrep:classificationScheme:DataType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:DataType/Date');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:DataType:DateTime',NULL,'urn:oasis:names:tc:ebxml-regrep:DataType:DateTime','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'DateTime','urn:oasis:names:tc:ebxml-regrep:classificationScheme:DataType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:DataType/DateTime');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:DataType:Double',NULL,'urn:oasis:names:tc:ebxml-regrep:DataType:Double','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Double','urn:oasis:names:tc:ebxml-regrep:classificationScheme:DataType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:DataType/Double');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:DataType:Duration',NULL,'urn:oasis:names:tc:ebxml-regrep:DataType:Duration','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Duration','urn:oasis:names:tc:ebxml-regrep:classificationScheme:DataType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:DataType/Duration');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:DataType:Float',NULL,'urn:oasis:names:tc:ebxml-regrep:DataType:Float','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Float','urn:oasis:names:tc:ebxml-regrep:classificationScheme:DataType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:DataType/Float');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:DataType:Integer',NULL,'urn:oasis:names:tc:ebxml-regrep:DataType:Integer','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Integer','urn:oasis:names:tc:ebxml-regrep:classificationScheme:DataType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:DataType/Integer');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:DataType:ObjectRef',NULL,'urn:oasis:names:tc:ebxml-regrep:DataType:ObjectRef','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'ObjectRef','urn:oasis:names:tc:ebxml-regrep:classificationScheme:DataType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:DataType/ObjectRef');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:DataType:String',NULL,'urn:oasis:names:tc:ebxml-regrep:DataType:String','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'String','urn:oasis:names:tc:ebxml-regrep:classificationScheme:DataType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:DataType/String');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:DataType:Time',NULL,'urn:oasis:names:tc:ebxml-regrep:DataType:Time','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Time','urn:oasis:names:tc:ebxml-regrep:classificationScheme:DataType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:DataType/Time');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:DataType:URI',NULL,'urn:oasis:names:tc:ebxml-regrep:DataType:URI','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'URI','urn:oasis:names:tc:ebxml-regrep:classificationScheme:DataType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:DataType/URI');

-- EventType
INSERT INTO CLASSSCHEME VALUES ('urn:oasis:names:tc:ebxml-regrep:classificationScheme:EventType',NULL,'urn:oasis:names:tc:ebxml-regrep:classificationScheme:EventType','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1','EventType','T','urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:EventType:Approved',NULL,'urn:oasis:names:tc:ebxml-regrep:EventType:Approved','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Approved','urn:oasis:names:tc:ebxml-regrep:classificationScheme:EventType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:EventType/Approved');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:EventType:Created',NULL,'urn:oasis:names:tc:ebxml-regrep:EventType:Created','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Created','urn:oasis:names:tc:ebxml-regrep:classificationScheme:EventType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:EventType/Created');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:EventType:Deleted',NULL,'urn:oasis:names:tc:ebxml-regrep:EventType:Deleted','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Deleted','urn:oasis:names:tc:ebxml-regrep:classificationScheme:EventType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:EventType/Deleted');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:EventType:Deprecated',NULL,'urn:oasis:names:tc:ebxml-regrep:EventType:Deprecated','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Deprecated','urn:oasis:names:tc:ebxml-regrep:classificationScheme:EventType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:EventType/Deprecated');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:EventType:Downloaded',NULL,'urn:oasis:names:tc:ebxml-regrep:EventType:Downloaded','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Downloaded','urn:oasis:names:tc:ebxml-regrep:classificationScheme:EventType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:EventType/Downloaded');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:EventType:Relocated',NULL,'urn:oasis:names:tc:ebxml-regrep:EventType:Relocated','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Relocated','urn:oasis:names:tc:ebxml-regrep:classificationScheme:EventType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:EventType/Relocated');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:EventType:Undeprecated',NULL,'urn:oasis:names:tc:ebxml-regrep:EventType:Undeprecated','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Undeprecated','urn:oasis:names:tc:ebxml-regrep:classificationScheme:EventType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:EventType/Undeprecated');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:EventType:Updated',NULL,'urn:oasis:names:tc:ebxml-regrep:EventType:Updated','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Updated','urn:oasis:names:tc:ebxml-regrep:classificationScheme:EventType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:EventType/Updated');

-- EmailTypes
INSERT INTO CLASSSCHEME VALUES ('urn:oasis:names:tc:ebxml-regrep:classificationScheme:EmailType',NULL,'urn:oasis:names:tc:ebxml-regrep:classificationScheme:EmailType','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1','EmailType','T','urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:EmailType:HomeEmail',NULL,'urn:oasis:names:tc:ebxml-regrep:EmailType:HomeEmail','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'HomeEmail','urn:oasis:names:tc:ebxml-regrep:classificationScheme:EmailType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:EmailType/HomeEmail');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:EmailType:OfficeEmail',NULL,'urn:oasis:names:tc:ebxml-regrep:EmailType:OfficeEmail','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'OfficeEmail','urn:oasis:names:tc:ebxml-regrep:classificationScheme:EmailType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:EmailType/OfficeEmail');

-- NodeTypes
INSERT INTO CLASSSCHEME VALUES ('urn:oasis:names:tc:ebxml-regrep:classificationScheme:NodeType',NULL,'urn:oasis:names:tc:ebxml-regrep:classificationScheme:NodeType','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1','NodeType','T','urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:NodeType:EmbeddedPath',NULL,'urn:oasis:names:tc:ebxml-regrep:NodeType:EmbeddedPath','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'EmbeddedPath','urn:oasis:names:tc:ebxml-regrep:classificationScheme:NodeType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:NodeType/EmbeddedPath');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:NodeType:NonUniqueCode',NULL,'urn:oasis:names:tc:ebxml-regrep:NodeType:NonUniqueCode','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'NonUniqueCode','urn:oasis:names:tc:ebxml-regrep:classificationScheme:NodeType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:NodeType/NonUniqueCode');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode',NULL,'urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'UniqueCode','urn:oasis:names:tc:ebxml-regrep:classificationScheme:NodeType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:NodeType/UniqueCode');

-- NotificationOptionType
INSERT INTO CLASSSCHEME VALUES ('urn:oasis:names:tc:ebxml-regrep:classificationScheme:NotificationOptionType',NULL,'urn:oasis:names:tc:ebxml-regrep:classificationScheme:NotificationOptionType','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1','NotificationOptionType','T','urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:NotificationOptionType:ObjectRefs',NULL,'urn:oasis:names:tc:ebxml-regrep:NotificationOptionType:ObjectRefs','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'ObjectRefs','urn:oasis:names:tc:ebxml-regrep:classificationScheme:NotificationOptionType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:NotificationOptionType/ObjectRefs');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:NotificationOptionType:Objects',NULL,'urn:oasis:names:tc:ebxml-regrep:NotificationOptionType:Objects','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Objects','urn:oasis:names:tc:ebxml-regrep:classificationScheme:NotificationOptionType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:NotificationOptionType/Objects');

-- PhoneType
INSERT INTO CLASSSCHEME VALUES ('urn:oasis:names:tc:ebxml-regrep:classificationScheme:PhoneType',NULL,'urn:oasis:names:tc:ebxml-regrep:classificationScheme:PhoneType','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1','PhoneType','T','urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:PhoneType:Beeper',NULL,'urn:oasis:names:tc:ebxml-regrep:PhoneType:Beeper','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Beeper','urn:oasis:names:tc:ebxml-regrep:classificationScheme:PhoneType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:PhoneType/Beeper');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:PhoneType:FAX',NULL,'urn:oasis:names:tc:ebxml-regrep:PhoneType:FAX','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'FAX','urn:oasis:names:tc:ebxml-regrep:classificationScheme:PhoneType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:PhoneType/FAX');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:PhoneType:HomePhone',NULL,'urn:oasis:names:tc:ebxml-regrep:PhoneType:HomePhone','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'HomePhone','urn:oasis:names:tc:ebxml-regrep:classificationScheme:PhoneType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:PhoneType/HomePhone');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:PhoneType:MobilePhone',NULL,'urn:oasis:names:tc:ebxml-regrep:PhoneType:MobilePhone','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'MobilePhone','urn:oasis:names:tc:ebxml-regrep:classificationScheme:PhoneType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:PhoneType/MobilePhone');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:PhoneType:OfficePhone',NULL,'urn:oasis:names:tc:ebxml-regrep:PhoneType:OfficePhone','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'OfficePhone','urn:oasis:names:tc:ebxml-regrep:classificationScheme:PhoneType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:PhoneType/OfficePhone');

-- QueryLanguage
INSERT INTO CLASSSCHEME VALUES ('urn:oasis:names:tc:ebxml-regrep:classificationScheme:QueryLanguage',NULL,'urn:oasis:names:tc:ebxml-regrep:classificationScheme:QueryLanguage','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1','QueryLanguage','T','urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:QueryLanguage:ebRSFilterQuery',NULL,'urn:oasis:names:tc:ebxml-regrep:QueryLanguage:ebRSFilterQuery','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'ebRSFilterQuery','urn:oasis:names:tc:ebxml-regrep:classificationScheme:QueryLanguage','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:QueryLanguage/ebRSFilterQuery');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:QueryLanguage:SQL-92',NULL,'urn:oasis:names:tc:ebxml-regrep:QueryLanguage:SQL-92','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'SQL-92','urn:oasis:names:tc:ebxml-regrep:classificationScheme:QueryLanguage','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:QueryLanguage/SQL-92');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:QueryLanguage:XPath',NULL,'urn:oasis:names:tc:ebxml-regrep:QueryLanguage:XPath','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'XPath','urn:oasis:names:tc:ebxml-regrep:classificationScheme:QueryLanguage','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:QueryLanguage/XPath');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:QueryLanguage:XQuery',NULL,'urn:oasis:names:tc:ebxml-regrep:QueryLanguage:XQuery','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'XQuery','urn:oasis:names:tc:ebxml-regrep:classificationScheme:QueryLanguage','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:QueryLanguage/XQuery');

-- SubjectRole
INSERT INTO CLASSSCHEME VALUES ('urn:oasis:names:tc:ebxml-regrep:classificationScheme:SubjectRole',NULL,'urn:oasis:names:tc:ebxml-regrep:classificationScheme:SubjectRole','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1','SubjectRole','T','urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:SubjectRole:ContentOwner',NULL,'urn:oasis:names:tc:ebxml-regrep:SubjectRole:ContentOwner','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'ContentOwner','urn:oasis:names:tc:ebxml-regrep:classificationScheme:SubjectRole','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:SubjectRole/ContentOwner');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:SubjectRole:Intermediary',NULL,'urn:oasis:names:tc:ebxml-regrep:SubjectRole:Intermediary','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Intermediary','urn:oasis:names:tc:ebxml-regrep:classificationScheme:SubjectRole','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:SubjectRole/Intermediary');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:SubjectRole:RegistryAdministrator',NULL,'urn:oasis:names:tc:ebxml-regrep:SubjectRole:RegistryAdministrator','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'RegistryAdministrator','urn:oasis:names:tc:ebxml-regrep:classificationScheme:SubjectRole','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:SubjectRole/RegistryAdministrator');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:SubjectRole:RegistryGuest',NULL,'urn:oasis:names:tc:ebxml-regrep:SubjectRole:RegistryGuest','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'RegistryGuest','urn:oasis:names:tc:ebxml-regrep:classificationScheme:SubjectRole','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:SubjectRole/RegistryGuest');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:freebxml:registry:demoDB:SubjectRole:ProjectMember',NULL,'urn:freebxml:registry:demoDB:SubjectRole:ProjectMember','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'ProjectMember','urn:oasis:names:tc:ebxml-regrep:classificationScheme:SubjectRole','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:SubjectRole/ProjectMember');
-- ProtocolType
INSERT INTO CLASSSCHEME VALUES ('urn:oasis:names:tc:ebxml-regrep:profile:ws:classificationScheme:ProtocolType',NULL,'urn:oasis:names:tc:ebxml-regrep:profile:ws:classificationScheme:ProtocolType','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1','ProtocolType','T','urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:profile:ws:ProtocolType:AS2',NULL,'urn:oasis:names:tc:ebxml-regrep:profile:ws:ProtocolType:AS2','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'AS2','urn:oasis:names:tc:ebxml-regrep:profile:ws:classificationScheme:ProtocolType','/urn:oasis:names:tc:ebxml-regrep:profile:ws:classificationScheme:ProtocolType/AS2');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:profile:ws:ProtocolType:Atom',NULL,'urn:oasis:names:tc:ebxml-regrep:profile:ws:ProtocolType:Atom','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Atom','urn:oasis:names:tc:ebxml-regrep:profile:ws:classificationScheme:ProtocolType','/urn:oasis:names:tc:ebxml-regrep:profile:ws:classificationScheme:ProtocolType/Atom');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:profile:ws:ProtocolType:SOAP',NULL,'urn:oasis:names:tc:ebxml-regrep:profile:ws:ProtocolType:SOAP','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'SOAP','urn:oasis:names:tc:ebxml-regrep:profile:ws:classificationScheme:ProtocolType','/urn:oasis:names:tc:ebxml-regrep:profile:ws:classificationScheme:ProtocolType/SOAP');

-- SOAPStyleType
INSERT INTO CLASSSCHEME VALUES ('urn:oasis:names:tc:ebxml-regrep:profile:ws:classificationScheme:SOAPStyleType',NULL,'urn:oasis:names:tc:ebxml-regrep:profile:ws:classificationScheme:SOAPStyleType','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1','SOAPStyleType','T','urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:profile:ws:SOAPStyleType:Document',NULL,'urn:oasis:names:tc:ebxml-regrep:profile:ws:SOAPStyleType:Document','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Document','urn:oasis:names:tc:ebxml-regrep:profile:ws:classificationScheme:SOAPStyleType','/urn:oasis:names:tc:ebxml-regrep:profile:ws:classificationScheme:SOAPStyleType/Document');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:profile:ws:SOAPStyleType:RPC',NULL,'urn:oasis:names:tc:ebxml-regrep:profile:ws:SOAPStyleType:RPC','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'RPC','urn:oasis:names:tc:ebxml-regrep:profile:ws:classificationScheme:SOAPStyleType','/urn:oasis:names:tc:ebxml-regrep:profile:ws:classificationScheme:SOAPStyleType/RPC');

-- StabilityType
INSERT INTO CLASSSCHEME VALUES ('urn:oasis:names:tc:ebxml-regrep:classificationScheme:StabilityType',NULL,'urn:oasis:names:tc:ebxml-regrep:classificationScheme:StabilityType','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1','StabilityType','T','urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:StabilityType:Dynamic',NULL,'urn:oasis:names:tc:ebxml-regrep:StabilityType:Dynamic','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Dynamic','urn:oasis:names:tc:ebxml-regrep:classificationScheme:StabilityType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:StabilityType/Dynamic');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:StabilityType:DynamicCompatible',NULL,'urn:oasis:names:tc:ebxml-regrep:StabilityType:DynamicCompatible','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'DynamicCompatible','urn:oasis:names:tc:ebxml-regrep:classificationScheme:StabilityType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:StabilityType/DynamicCompatible');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:StabilityType:Static',NULL,'urn:oasis:names:tc:ebxml-regrep:StabilityType:Static','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Static','urn:oasis:names:tc:ebxml-regrep:classificationScheme:StabilityType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:StabilityType/Static');

-- DeletionScopeType
INSERT INTO CLASSSCHEME VALUES ('urn:oasis:names:tc:ebxml-regrep:classificationScheme:DeletionScopeType',NULL,'urn:oasis:names:tc:ebxml-regrep:classificationScheme:DeletionScopeType','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1','DeletionScopeType','T','urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:DeletionScopeType:DeleteAll',NULL,'urn:oasis:names:tc:ebxml-regrep:DeletionScopeType:DeleteAll','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'DeleteAll','urn:oasis:names:tc:ebxml-regrep:classificationScheme:DeletionScopeType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:DeletionScopeType/DeleteAll');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:DeletionScopeType:DeleteRepositoryItemOnly',NULL,'urn:oasis:names:tc:ebxml-regrep:DeletionScopeType:DeleteRepositoryItemOnly','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'DeleteRepositoryItemOnly','urn:oasis:names:tc:ebxml-regrep:classificationScheme:DeletionScopeType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:DeletionScopeType/DeleteRepositoryItemOnly');

-- ErrorHandlingModel
INSERT INTO CLASSSCHEME VALUES ('urn:oasis:names:tc:ebxml-regrep:classificationScheme:ErrorHandlingModel',NULL,'urn:oasis:names:tc:ebxml-regrep:classificationScheme:ErrorHandlingModel','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1','ErrorHandlingModel','T','urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ErrorHandlingModel:FailOnError',NULL,'urn:oasis:names:tc:ebxml-regrep:ErrorHandlingModel:FailOnError','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'FailOnError','urn:oasis:names:tc:ebxml-regrep:classificationScheme:ErrorHandlingModel','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ErrorHandlingModel/FailOnError');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ErrorHandlingModel:LogErrorAndContinue',NULL,'urn:oasis:names:tc:ebxml-regrep:ErrorHandlingModel:LogErrorAndContinue','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'LogErrorAndContinue','urn:oasis:names:tc:ebxml-regrep:classificationScheme:ErrorHandlingModel','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ErrorHandlingModel/LogErrorAndContinue');

-- ErrorSeverityType
INSERT INTO CLASSSCHEME VALUES ('urn:oasis:names:tc:ebxml-regrep:classificationScheme:ErrorSeverityType',NULL,'urn:oasis:names:tc:ebxml-regrep:classificationScheme:ErrorSeverityType','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1','ErrorSeverityType','T','urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ErrorSeverityType:Error',NULL,'urn:oasis:names:tc:ebxml-regrep:ErrorSeverityType:Error','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Error','urn:oasis:names:tc:ebxml-regrep:classificationScheme:ErrorSeverityType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ErrorSeverityType/Error');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ErrorSeverityType:Warning',NULL,'urn:oasis:names:tc:ebxml-regrep:ErrorSeverityType:Warning','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Warning','urn:oasis:names:tc:ebxml-regrep:classificationScheme:ErrorSeverityType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ErrorSeverityType/Warning');

-- ObjectType
INSERT INTO CLASSSCHEME VALUES ('urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType',NULL,'urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1','ObjectType','T','urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'RegistryObject','urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject');

-- ResponseStatusType
INSERT INTO CLASSSCHEME VALUES ('urn:oasis:names:tc:ebxml-regrep:classificationScheme:ResponseStatusType',NULL,'urn:oasis:names:tc:ebxml-regrep:classificationScheme:ResponseStatusType','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1','ResponseStatusType','T','urn:oasis:names:tc:ebxml-regrep:NodeType:UniqueCode');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ResponseStatusType:Failure',NULL,'urn:oasis:names:tc:ebxml-regrep:ResponseStatusType:Failure','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Failure','urn:oasis:names:tc:ebxml-regrep:classificationScheme:ResponseStatusType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ResponseStatusType/Failure');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ResponseStatusType:Success',NULL,'urn:oasis:names:tc:ebxml-regrep:ResponseStatusType:Success','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Success','urn:oasis:names:tc:ebxml-regrep:classificationScheme:ResponseStatusType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ResponseStatusType/Success');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ResponseStatusType:Unavailable',NULL,'urn:oasis:names:tc:ebxml-regrep:ResponseStatusType:Unavailable','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Unavailable','urn:oasis:names:tc:ebxml-regrep:classificationScheme:ResponseStatusType','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ResponseStatusType/Unavailable');

-- other
INSERT INTO CLASSIFICATION VALUES ('urn:oasis:names:tc:ebxml-regrep:classification:RegistryAdministrator',NULL,'urn:oasis:names:tc:ebxml-regrep:classification:RegistryAdministrator','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Classification','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'urn:oasis:names:tc:ebxml-regrep:SubjectRole:RegistryAdministrator',NULL,'urn:freebxml:registry:predefinedusers:registryoperator',NULL);

-- RegistryObject
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:AdhocQuery',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:AdhocQuery','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'AdhocQuery','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/AdhocQuery');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Association',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Association','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Association','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/Association');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:AuditableEvent',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:AuditableEvent','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'AuditableEvent','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/AuditableEvent');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Classification',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Classification','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Classification','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/Classification');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'ClassificationNode','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/ClassificationNode');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationScheme','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'ClassificationScheme','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/ClassificationScheme');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExternalIdentifier',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExternalIdentifier','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'ExternalIdentifier','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/ExternalIdentifier');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExternalLink',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExternalLink','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'ExternalLink','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/ExternalLink');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'ExtrinsicObject','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/ExtrinsicObject');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:image',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:image','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'image','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/ExtrinsicObject/image');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:image:gif',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:image:gif','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'gif','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:image','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/ExtrinsicObject/image/gif');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:image:jpeg',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:image:jpeg','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'jpeg','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:image','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/ExtrinsicObject/image/jpeg');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:WSDL',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:WSDL','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'WSDL','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/ExtrinsicObject/WSDL');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:WSDL:Binding',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:WSDL:Binding','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Binding','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:WSDL','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/ExtrinsicObject/WSDL/Binding');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:WSDL:Port',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:WSDL:Port','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Port','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:WSDL','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/ExtrinsicObject/WSDL/Port');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:WSDL:PortType',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:WSDL:PortType','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'PortType','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:WSDL','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/ExtrinsicObject/WSDL/PortType');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:WSDL:Service',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:WSDL:Service','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Service','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:WSDL','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/ExtrinsicObject/WSDL/Service');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:XACML',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:XACML','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'XACML','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/ExtrinsicObject/XACML');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:XACML:Policy',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:XACML:Policy','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Policy','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:XACML','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/ExtrinsicObject/XACML/Policy');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:XACML:PolicySet',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:XACML:PolicySet','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'PolicySet','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:XACML','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/ExtrinsicObject/XACML/PolicySet');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:XHTML',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:XHTML','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'XHTML','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/ExtrinsicObject/XHTML');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:XHTML:XForm',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:XHTML:XForm','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'XForm','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:XHTML','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/ExtrinsicObject/XHTML/XForm');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:XML',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:XML','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'XML','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/ExtrinsicObject/XML');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:XML:Schematron',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:XML:Schematron','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Schematron','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/ExtrinsicObject/Schematron');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:XMLSchema',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:XMLSchema','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'XMLSchema','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/ExtrinsicObject/XMLSchema');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:XSLT',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject:XSLT','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'XSLT','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ExtrinsicObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/ExtrinsicObject/XSLT');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Federation',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Federation','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Federation','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/Federation');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Notification',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Notification','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Notification','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/Notification');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Organization',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Organization','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Organization','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/Organization');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Person',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Person','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Person','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/Person');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Person:User',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Person:User','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'User','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Person','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/Person/User');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Registry',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Registry','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Registry','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/Registry');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:RegistryPackage',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:RegistryPackage','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'RegistryPackage','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/RegistryPackage');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Service',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Service','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Service','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/Service');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ServiceBinding',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ServiceBinding','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'ServiceBinding','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/ServiceBinding');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:SpecificationLink',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:SpecificationLink','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'SpecificationLink','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/SpecificationLink');
INSERT INTO CLASSIFICATIONNODE VALUES ('urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Subscription',NULL,'urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:Subscription','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject:ClassificationNode','urn:oasis:names:tc:ebxml-regrep:StatusType:Submitted','1.1',NULL,'Subscription','urn:oasis:names:tc:ebxml-regrep:ObjectType:RegistryObject','/urn:oasis:names:tc:ebxml-regrep:classificationScheme:ObjectType/RegistryObject/Subscription');

