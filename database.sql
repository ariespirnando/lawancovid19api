-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               5.6.24 - MySQL Community Server (GPL)
-- Server OS:                    Win32
-- HeidiSQL Version:             9.4.0.5125
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- Dumping database structure for covid19
CREATE DATABASE IF NOT EXISTS `covid19` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `covid19`;

-- Dumping structure for table covid19.data_corona
CREATE TABLE IF NOT EXISTS `data_corona` (
  `idata_corona` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `dtgl_update` date NOT NULL DEFAULT '0000-00-00',
  `kprovinsi` int(11) DEFAULT NULL,
  `ikasus_positif` int(11) DEFAULT NULL,
  `ikasus_sehat` int(11) DEFAULT NULL,
  `ikasus_meninggal` int(11) DEFAULT NULL,
  PRIMARY KEY (`idata_corona`),
  KEY `dtgl_update` (`dtgl_update`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=latin1;

-- Data exporting was unselected.
-- Dumping structure for table covid19.data_provinsi
CREATE TABLE IF NOT EXISTS `data_provinsi` (
  `idata_provinsi` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `kprovinsi` int(11) DEFAULT NULL,
  `vprovinsi` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`idata_provinsi`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=latin1;

-- Data exporting was unselected.
-- Dumping structure for procedure covid19.lihat_grafik_perprov
DELIMITER //
CREATE DEFINER=`root`@`localhost` PROCEDURE `lihat_grafik_perprov`(
	IN `kode_provinsi` INT

)
BEGIN
SELECT dtgl_update, ikasus_positif  FROM data_corona WHERE kprovinsi = kode_provinsi ORDER BY dtgl_update ASC LIMIT 10 ; 
END//
DELIMITER ;

-- Dumping structure for procedure covid19.lihat_kasus_perprovinsi
DELIMITER //
CREATE DEFINER=`root`@`localhost` PROCEDURE `lihat_kasus_perprovinsi`(
	IN `kode_provinsi` INT


)
BEGIN
SELECT ikasus_positif, ikasus_sehat ,ikasus_meninggal, data_provinsi.kprovinsi, vprovinsi  FROM data_corona JOIN data_provinsi ON data_provinsi.kprovinsi = data_corona.kprovinsi
WHERE data_provinsi.kprovinsi = kode_provinsi ORDER BY dtgl_update DESC LIMIT 1; 
END//
DELIMITER ;

-- Dumping structure for procedure covid19.lihat_provinsi
DELIMITER //
CREATE DEFINER=`root`@`localhost` PROCEDURE `lihat_provinsi`()
BEGIN
SELECT kprovinsi, vprovinsi FROM data_provinsi ORDER By kprovinsi;
END//
DELIMITER ;

-- Dumping structure for procedure covid19.list_kasus_all
DELIMITER //
CREATE DEFINER=`root`@`localhost` PROCEDURE `list_kasus_all`()
BEGIN
SELECT p.kprovinsi,p.vprovinsi ,
(SELECT c.ikasus_positif FROM data_corona c WHERE c.kprovinsi = p.kprovinsi ORDER BY c.dtgl_update DESC LIMIT 1) as positif,
(SELECT c.ikasus_sehat FROM data_corona c WHERE c.kprovinsi = p.kprovinsi ORDER BY c.dtgl_update DESC LIMIT 1) as sembuh,
(SELECT c.ikasus_meninggal FROM data_corona c WHERE c.kprovinsi = p.kprovinsi ORDER BY c.dtgl_update DESC LIMIT 1) as meninggal
FROM data_provinsi p ORDER BY p.kprovinsi;
END//
DELIMITER ;

-- Dumping structure for procedure covid19.update_data
DELIMITER //
CREATE DEFINER=`root`@`localhost` PROCEDURE `update_data`(
	IN `kodeprovinsi` VARCHAR(50),
	IN `namaprovinsi` VARCHAR(50),
	IN `kasuspositif` INT,
	IN `kasussehat` INT,
	IN `kasusmeninggal` INT



)
BEGIN
DECLARE CekProvinsi,CekCorona int;

SELECT COUNT(idata_provinsi) INTO CekProvinsi FROM data_provinsi WHERE kprovinsi = kodeprovinsi;
IF CekProvinsi = 0 THEN
	INSERT INTO data_provinsi (kprovinsi,vprovinsi) VALUES (kodeprovinsi,namaprovinsi); 
END IF;

SELECT COUNT(idata_corona) INTO CekCorona FROM data_corona WHERE dtgl_update=CURDATE() AND kprovinsi = kodeprovinsi;
IF CekCorona > 0 THEN
	UPDATE data_corona SET ikasus_positif = kasuspositif, ikasus_sehat=kasussehat, ikasus_meninggal = kasusmeninggal WHERE dtgl_update=CURDATE() AND kprovinsi = kodeprovinsi;
ELSE
	INSERT INTO data_corona (kprovinsi,dtgl_update,ikasus_positif,ikasus_sehat,ikasus_meninggal) VALUES (kodeprovinsi,CURDATE(),kasuspositif,kasussehat,kasusmeninggal);
END IF; 

END//
DELIMITER ;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
