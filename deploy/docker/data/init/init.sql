-- MySQL dump 10.13  Distrib 5.7.24, for osx11.1 (x86_64)
--
-- Host: 127.0.0.1    Database: engine
-- ------------------------------------------------------
-- Server version	5.7.36

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
-- Table structure for table `data_source`
--

DROP TABLE IF EXISTS `data_source`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `data_source` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `confType` varchar(50) DEFAULT NULL,
  `type` varchar(50) DEFAULT NULL,
  `content` text,
  `status` int(11) DEFAULT '1',
  `createTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updateTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `data_source`
--

LOCK TABLES `data_source` WRITE;
/*!40000 ALTER TABLE `data_source` DISABLE KEYS */;
INSERT INTO `data_source` VALUES (1,'global','global','global','basic:\n  # debug | info | warn | error | fatal | panic\n  logLevel: info\n  # true|false, with debug level, it prints more debug info\n  debug: false\n  # true|false, if it\'s set to true, then the log will be print to console\n  consoleLog: false\n  # true|false, if it\'s set to true, then the log will be print to log file\n  fileLog: true\n  # Whether to disable the log timestamp, useful when output is redirected to logging\n  #	system like syslog that already adds timestamps.\n  logDisableTimestamp: false\n  # syslog settings\n  syslog:\n    # true|false, if it\'s set to true, then the log will be print to syslog\n    enable: false\n    # The syslog protocol, tcp or udp; Leave empty if no remote syslog server is used\n    network: udp\n    # The syslog server address; Leave empty if no remote syslog server is used\n    address: localhost:514\n    # The syslog level, supports debug, info, warn, error\n    level: info\n    # The syslog tag; Leave empty if no tag is used\n    tag: kuiper\n  # How many hours to split the file\n  rotateTime: 24\n  # Maximum file storage hours\n  maxAge: 72\n  # Maximum file size in bytes, if this is set, maxAge will be ignored\n  rotateSize: 10485760 # 10 MB\n  # Maximum log file count\n  rotateCount: 3\n  # CLI ip\n  ip: 0.0.0.0\n  # CLI port\n  port: 20498\n  # REST service ip\n  restIp: 0.0.0.0\n  # REST service port\n  restPort: 9082\n  # The global time zone from the IANA time zone database, or Local if not set.\n  timezone: Local\n  # true|false, when true, will check the RSA jwt token for rest api\n  authentication: false\n  #  restTls:\n  #    certfile: /var/https-server.crt\n  #    keyfile: /var/https-server.key\n  # Prometheus settings\n  prometheus: false\n  prometheusPort: 20499\n  # The URL where hosts all of pre-build plugins. By default, it\'s at packages.emqx.net\n  pluginHosts: https://packages.emqx.net\n  # Whether to ignore case in SQL processing. Note that, the name of customized function by plugins are case-sensitive.\n  ignoreCase: false\n  sql:\n    # maxConnections indicates the max connections for the certain database instance group by driver and dsn sharing between the sources/sinks\n    # 0 indicates unlimited\n    maxConnections: 0\n  rulePatrolInterval: 10s\n  # cfgStorageType indicates the storage type to store the config, support \"file\",\"sqlite\" and \"fdb\"\n  cfgStorageType: file\n  # enableOpenZiti indicates whether to enable OpenZiti for eKuiper REST service. Currently, it is only supported to work with EdgeX secure mode.\n  enableOpenZiti: false\n\n# The default options for all rules. Each rule can override this setting by defining its own option\nrule:\n  # The qos of the rule. The values can be 0: At most once; 1: At least once; 2: Exactly once\n  # If qos is bigger than 0, the checkpoint mechanism will launch to save states so that they can be\n  # restored for unintended interrupt or planned restart of the rule. The performance may be affected\n  # to enable the checkpoint mechanism\n  qos: 0\n  # The interval in millisecond to run the checkpoint mechanism.\n  checkpointInterval: 300000\n  # Whether to send errors to sinks\n  sendError: true\n  # The strategy to retry for rule errors.\n  restartStrategy:\n    # The maximum retry times\n    attempts: 0\n    # The interval in millisecond to retry\n    delay: 1000\n    # The maximum interval in millisecond to retry\n    maxDelay: 30000\n    # The exponential to increase the interval. It can be a float value.\n    multiplier: 2\n    # How large random value will be added or subtracted to the delay to prevent restarting multiple rules at the same time.\n    jitterFactor: 0.1\nsink:\n  # Control to enable cache or not. If it\'s set to true, then the cache will be enabled, otherwise, it will be disabled.\n  enableCache: false\n\n  # The maximum number of messages to be cached in memory.\n  memoryCacheThreshold: 1024\n\n  # The maximum number of messages to be cached in the disk.\n  maxDiskCache: 1024000\n\n  # The number of messages for a buffer page which is the unit to read/write to disk in batch to prevent frequent IO\n  bufferPageSize: 256\n\n  # The interval in millisecond to resend the cached messages\n  resendInterval: 0\n\n  # Whether to clean the cache when the rule stops\n  cleanCacheAtStop: false\n\nsource:\n  ## Configurations for the global http data server for httppush source\n  # HTTP data service ip\n  httpServerIp: 0.0.0.0\n  # HTTP data service port\n  httpServerPort: 10081\n  # httpServerTls:\n  #    certfile: /var/https-server.crt\n  #    keyfile: /var/https-server.key\n\nstore:\n  #Type of store that will be used for keeping state of the application\n  type: sqlite\n  extStateType: sqlite\n  redis:\n    host: localhost\n    port: 6379\n    password: kuiper\n    #Timeout in ms\n    timeout: 1000\n  sqlite:\n    #Sqlite file name, if left empty name of db will be sqliteKV.db\n    name:\n\n# The settings for portable plugin\nportable:\n  # The executable of python. Specify this if you have multiple python instances in your system\n  # or other circumstance where the python executable cannot be successfully invoked through the default command.\n  pythonBin: python\n  # control init timeout in ms. If the init time is longer than this value, the plugin will be terminated.\n  initTimeout: 6000',1,'2024-06-06 13:06:58','2024-06-21 14:40:30'),(2,'mqtt','source','mqtt','#Global MQTT configurations\ndefault:\n  qos: 1\n  server: \"tcp://bifromq-server:1883\"\n  protocolVersion: \"3.1.1\"\n  insecureSkipVerify: false\n  #decompression: zlib\n  #username: user1\n  #password: password\n  #certificationPath: /var/kuiper/xyz-certificate.pem\n  #privateKeyPath: /var/kuiper/xyz-private.pem.key\n  #rootCaPath: /var/kuiper/xyz-rootca.pem\n  #insecureSkipVerify: false\n  #connectionSelector: mqtt.mqtt_conf1\n  #kubeedgeVersion: \n  #kubeedgeModelFile: \"\"\n\ndemo_conf: #Conf_key\n qos: 1\n server: \"tcp://10.211.55.6:1883\"\n',1,'2024-06-06 13:08:15','2024-06-26 09:41:04'),(3,'redis','source','redis','default:\n  # the redis host address\n  addr: \"redis-server:6379\"\n  # currently supports string and list only\n  datatype: \"string\"\n#  username: \"\"\n#  password: \"\"',1,'2024-06-06 13:09:03','2024-06-26 09:41:14'),(4,'kafka','source','kafka','default:\n  brokers: \"kafka-server:9092\"\n  partition: 0\n  maxBytes: 1000000\n  topic:\n  groupID:\n  maxAttempts: 1',1,'2024-06-21 15:11:18','2024-06-26 09:41:22');
/*!40000 ALTER TABLE `data_source` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `permission`
--

DROP TABLE IF EXISTS `permission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `permission` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `code` varchar(50) NOT NULL,
  `type` varchar(255) NOT NULL,
  `parentId` int(11) DEFAULT NULL,
  `path` varchar(255) DEFAULT NULL,
  `redirect` varchar(255) DEFAULT NULL,
  `icon` varchar(255) DEFAULT NULL,
  `component` varchar(255) DEFAULT NULL,
  `layout` varchar(255) DEFAULT NULL,
  `keepAlive` tinyint(4) DEFAULT NULL,
  `method` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `show` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否展示在页面菜单',
  `enable` tinyint(4) NOT NULL DEFAULT '1',
  `order` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `IDX_30e166e8c6359970755c5727a2` (`code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `permission`
--

LOCK TABLES `permission` WRITE;
/*!40000 ALTER TABLE `permission` DISABLE KEYS */;
INSERT INTO `permission` VALUES (1,'资源管理','Resource_Mgt','MENU',2,'/pms/resource',NULL,'i-fe:list','/src/views/pms/resource/index.vue',NULL,NULL,NULL,NULL,1,1,1),(2,'系统管理','SysMgt','MENU',NULL,NULL,NULL,'i-fe:grid',NULL,NULL,NULL,NULL,NULL,1,1,10),(3,'角色管理','RoleMgt','MENU',2,'/pms/role',NULL,'i-fe:user-check','/src/views/pms/role/index.vue',NULL,NULL,NULL,NULL,1,1,2),(4,'用户管理','UserMgt','MENU',2,'/pms/user',NULL,'i-fe:user','/src/views/pms/user/index.vue',NULL,1,NULL,NULL,1,1,3),(5,'分配用户','RoleUser','MENU',3,'/pms/role/user/:roleId',NULL,'i-fe:user-plus','/src/views/pms/role/role-user.vue',NULL,NULL,NULL,NULL,0,1,1),(6,'全局管理','GlobalMgt','MENU',NULL,NULL,NULL,'i-fe:grid',NULL,NULL,NULL,NULL,NULL,1,1,1),(7,'节点管理','NodeMgt','MENU',6,'/node/list',NULL,'i-fe:folder','/src/views/node/index.vue',NULL,1,'/src/views/node/index.vue',NULL,1,1,2),(8,'个人资料','UserProfile','MENU',NULL,'/profile',NULL,'i-fe:user','/src/views/profile/index.vue',NULL,NULL,NULL,NULL,0,1,99),(9,'基础功能','Base','MENU',NULL,'/base',NULL,'i-fe:grid',NULL,NULL,NULL,NULL,NULL,0,1,0),(10,'基础组件','BaseComponents','MENU',9,'/base/components',NULL,'i-me:awesome','/src/views/base/index.vue',NULL,NULL,NULL,NULL,1,1,1),(11,'Unocss','Unocss','MENU',9,'/base/unocss',NULL,'i-me:awesome','/src/views/base/unocss.vue',NULL,NULL,NULL,NULL,1,1,2),(12,'KeepAlive','KeepAlive','MENU',9,'/base/keep-alive',NULL,'i-me:awesome','/src/views/base/keep-alive.vue',NULL,1,NULL,NULL,1,1,3),(13,'创建新用户','AddUser','BUTTON',4,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,1,1,1),(14,'图标 Icon','Icon','MENU',9,'/base/icon',NULL,'i-fe:feather','/src/views/base/unocss-icon.vue',NULL,NULL,NULL,NULL,1,1,0),(18,'全局配置','GlobalConfig','MENU',6,'/global/config','','i-fe:package','/src/views/configuration/global.vue','',0,'/src/views/configuration/global.vue','',1,1,3),(19,'数据源管理','Datasource','MENU',NULL,'','','i-fe:trello','','',0,'','',1,0,2),(20,'规则集管理','RuleSet','MENU',NULL,'','','i-fe:volume-2','','',0,'','',1,1,4),(21,'Mqtt配置','MqttMgt','MENU',19,'/source/mqtt','','i-simple-icons:juejin','/src/views/configuration/mqtt.vue','',0,'/src/views/configuration/mqtt.vue','',1,1,1),(22,'kafka配置','KafkaMgt','MENU',19,'/source/kafka','','i-fe:truck','/src/views/configuration/kafka.vue','',0,'/src/views/configuration/kafka.vue','',1,1,2),(23,'Redis配置','RedisMgt','MENU',19,'/source/redis','','i-fe:wind','/src/views/configuration/redis.vue','',0,'/src/views/configuration/redis.vue','',1,1,3),(24,'规则集列表','RuleSetList','MENU',20,'/ruleset/index','','i-fe:tool','/src/views/ruleset/index.vue','',0,'/src/views/ruleset/index.vue','',1,1,1),(25,'添加新节点','AddNode','BUTTON',7,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,1,1,1),(26,'编辑规则集','RuleSetEdit','MENU',24,'/ruleset/:ruleSetId','','i-fe:youtube','/src/views/ruleset/ruleset.vue','',0,'','',0,1,1);
/*!40000 ALTER TABLE `permission` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `profile`
--

DROP TABLE IF EXISTS `profile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `profile` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `gender` int(11) DEFAULT NULL,
  `avatar` varchar(255) NOT NULL DEFAULT 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif?imageView2/1/w/80/h/80',
  `address` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `userId` int(11) NOT NULL,
  `nickName` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `IDX_a24972ebd73b106250713dcddd` (`userId`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `profile`
--

LOCK TABLES `profile` WRITE;
/*!40000 ALTER TABLE `profile` DISABLE KEYS */;
INSERT INTO `profile` VALUES (1,NULL,'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif?imageView2/1/w/80/h/80',NULL,NULL,1,'admin');
/*!40000 ALTER TABLE `profile` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role`
--

DROP TABLE IF EXISTS `role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `code` varchar(50) NOT NULL,
  `name` varchar(50) NOT NULL,
  `enable` tinyint(4) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `IDX_ee999bb389d7ac0fd967172c41` (`code`) USING BTREE,
  UNIQUE KEY `IDX_ae4578dcaed5adff96595e6166` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role`
--

LOCK TABLES `role` WRITE;
/*!40000 ALTER TABLE `role` DISABLE KEYS */;
INSERT INTO `role` VALUES (1,'SUPER_ADMIN','超级管理员',1),(2,'ROLE_QA','质检员',1);
/*!40000 ALTER TABLE `role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role_permissions_permission`
--

DROP TABLE IF EXISTS `role_permissions_permission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role_permissions_permission` (
  `roleId` int(11) NOT NULL,
  `permissionId` int(11) NOT NULL,
  PRIMARY KEY (`roleId`,`permissionId`) USING BTREE,
  KEY `IDX_b36cb2e04bc353ca4ede00d87b` (`roleId`) USING BTREE,
  KEY `IDX_bfbc9e263d4cea6d7a8c9eb3ad` (`permissionId`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_permissions_permission`
--

LOCK TABLES `role_permissions_permission` WRITE;
/*!40000 ALTER TABLE `role_permissions_permission` DISABLE KEYS */;
INSERT INTO `role_permissions_permission` VALUES (2,1),(2,2),(2,3),(2,4),(2,5),(2,9),(2,10),(2,11),(2,12),(2,14);
/*!40000 ALTER TABLE `role_permissions_permission` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rule`
--

DROP TABLE IF EXISTS `rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `rule` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ruleSetID` int(11) NOT NULL,
  `name` varchar(50) NOT NULL,
  `ruleType` varchar(50) NOT NULL,
  `statement` varchar(1000) DEFAULT NULL,
  `actions` varchar(1000) DEFAULT NULL,
  `status` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) DEFAULT '0',
  `statusCheck` int(11) DEFAULT '0',
  `statusCheckTime` datetime DEFAULT NULL,
  `statusCheckText` text,
  `createTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updateTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rule`
--

LOCK TABLES `rule` WRITE;
/*!40000 ALTER TABLE `rule` DISABLE KEYS */;
INSERT INTO `rule` VALUES (1,1,'demo','sql','select * from demo where temperature;','[{\n    \"log\":  {}\n}]',2,0,3,'2024-06-21 14:23:06','Get \"http://192.168.31.72:9082/rules/demo/status\": dial tcp 192.168.31.72:9082: connect: connection refused','2024-06-04 14:10:38','2024-06-21 06:23:05');
/*!40000 ALTER TABLE `rule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rule_set`
--

DROP TABLE IF EXISTS `rule_set`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `rule_set` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `status` int(11) DEFAULT '0',
  `deleted` tinyint(4) DEFAULT '0',
  `workerID` int(11) DEFAULT '0',
  `statusCheck` int(11) DEFAULT NULL,
  `statusCheckTime` datetime DEFAULT NULL,
  `scheduleTime` datetime DEFAULT NULL,
  `createTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updateTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rule_set`
--

LOCK TABLES `rule_set` WRITE;
/*!40000 ALTER TABLE `rule_set` DISABLE KEYS */;
INSERT INTO `rule_set` VALUES (1,'test',2,0,16,3,'2024-06-21 14:23:06','2024-06-07 14:37:58','2024-06-04 14:07:46','2024-06-21 06:23:05'),(2,'hello',0,0,0,0,NULL,NULL,'2024-06-21 17:56:05','2024-06-21 17:56:05');
/*!40000 ALTER TABLE `rule_set` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `stream`
--

DROP TABLE IF EXISTS `stream`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `stream` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ruleSetID` int(11) NOT NULL,
  `name` varchar(50) DEFAULT NULL,
  `statement` varchar(800) DEFAULT NULL,
  `status` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) DEFAULT '0',
  `createTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updateTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `stream`
--

LOCK TABLES `stream` WRITE;
/*!40000 ALTER TABLE `stream` DISABLE KEYS */;
INSERT INTO `stream` VALUES (1,1,'demo','create stream demo (temperature float, humidity bigint) WITH (FORMAT=\"JSON\", DATASOURCE=\"devices/+/messages\")',2,0,'2024-06-04 14:08:53','2024-06-04 09:57:29');
/*!40000 ALTER TABLE `stream` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  `enable` tinyint(4) NOT NULL DEFAULT '1',
  `createTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updateTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `IDX_78a916df40e02a9deb1c4b75ed` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'admin','21232f297a57a5a743894a0e4a801fc3',1,'2023-11-18 16:18:59.150632','2024-05-25 11:48:22.929913');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_roles_role`
--

DROP TABLE IF EXISTS `user_roles_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_roles_role` (
  `userId` int(11) NOT NULL,
  `roleId` int(11) NOT NULL,
  PRIMARY KEY (`userId`,`roleId`) USING BTREE,
  KEY `IDX_5f9286e6c25594c6b88c108db7` (`userId`) USING BTREE,
  KEY `IDX_4be2f7adf862634f5f803d246b` (`roleId`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_roles_role`
--

LOCK TABLES `user_roles_role` WRITE;
/*!40000 ALTER TABLE `user_roles_role` DISABLE KEYS */;
INSERT INTO `user_roles_role` VALUES (1,1),(1,2);
/*!40000 ALTER TABLE `user_roles_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `worker`
--

DROP TABLE IF EXISTS `worker`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `worker` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `tag` varchar(255) DEFAULT NULL,
  `ip` varchar(255) DEFAULT NULL,
  `status` int(11) DEFAULT '0',
  `port` int(11) DEFAULT NULL,
  `heartbeatMisses` int(11) DEFAULT '0',
  `heartbeatTime` datetime DEFAULT NULL,
  `lastSourcesTime` datetime DEFAULT NULL,
  `createTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updateTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `worker`
--

LOCK TABLES `worker` WRITE;
/*!40000 ALTER TABLE `worker` DISABLE KEYS */;
/*!40000 ALTER TABLE `worker` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-06-27  9:51:29
