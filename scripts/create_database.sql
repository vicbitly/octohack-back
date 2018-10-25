CREATE DATABASE IF NOT EXISTS `octohack` DEFAULT CHARACTER SET utf8 */;

USE `octohack`;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) COLLATE utf8_bin NOT NULL,
  `email` varchar(255) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username_UNIQUE` (`username`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


INSERT INTO `users` VALUES (1,'amelia','amelia@bit.ly'),(2,'vic','ve@bit.ly');
