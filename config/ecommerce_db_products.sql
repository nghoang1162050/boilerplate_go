-- MySQL dump 10.13  Distrib 8.0.41, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: ecommerce_db
-- ------------------------------------------------------
-- Server version	9.2.0

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
-- Table structure for table `products`
--
use ecommerce_db;
DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `brand` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `description` text COLLATE utf8mb4_unicode_ci,
  `price` decimal(10,2) NOT NULL,
  `stock_quantity` int DEFAULT '0',
  `cpu` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `ram` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `storage` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `gpu` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES ('','Gaming Laptop X1','Alienware','High performance gaming laptop',2200.00,10,'Intel Core i9','32GB','1TB SSD','NVIDIA RTX 3080','2025-04-09 02:10:44'),('2cdcfcc6-a158-4b7a-8402-137a7641137c','UltraBook Slim Z','Dell','Slim and lightweight ultrabook',1350.00,15,'Intel Core i7','16GB','512GB SSD','Intel Iris Xe','2025-04-09 02:10:44'),('a53b9d69-7848-46f0-a89b-7f872697f654','Gaming Laptop X1','Alienware','High performance gaming laptop',2200.00,10,'Intel Core i9','32GB','1TB SSD','NVIDIA RTX 3080','2025-04-09 02:10:44'),('a53b9d69-7848-46f0-a89b-7f872697f655','Gaming Laptop X3','Alienware','High performance gaming laptop',2200.00,10,'Intel Core i9','32GB','1TB SSD','NVIDIA RTX 3080','2025-04-09 02:10:44'),('b48fa7a3-150f-11f0-a903-164fa3fc2eef','Laptop Pro 14','Apple','A powerful MacBook Pro with M1 chip.',1999.99,10,'Apple M1','16GB','512GB SSD','Integrated','2025-04-09 06:56:04'),('b490be45-150f-11f0-a903-164fa3fc2eef','Inspiron 15','Dell','Dell Inspiron laptop for everyday use.',649.99,25,'Intel Core i5-1135G7','8GB','256GB SSD','Intel Iris Xe','2025-04-09 06:56:04'),('b4911cc8-150f-11f0-a903-164fa3fc2eef','ThinkPad X1 Carbon','Lenovo','Premium business ultrabook.',1499.00,15,'Intel Core i7-1165G7','16GB','1TB SSD','Integrated','2025-04-09 06:56:04'),('b4911e1a-150f-11f0-a903-164fa3fc2eef','ROG Zephyrus G14','Asus','High-performance gaming laptop.',1799.00,8,'AMD Ryzen 9 6900HS','32GB','1TB SSD','NVIDIA RTX 3060','2025-04-09 06:56:04'),('b4911e67-150f-11f0-a903-164fa3fc2eef','Pavilion 14','HP','Affordable laptop with solid performance.',599.99,30,'Intel Core i5-1240P','8GB','512GB SSD','Intel Iris Xe','2025-04-09 06:56:04'),('b4911f61-150f-11f0-a903-164fa3fc2eef','Aspire 5','Acer','Versatile everyday laptop.',549.00,20,'Intel Core i3-1115G4','8GB','256GB SSD','Intel UHD','2025-04-09 06:56:04'),('b4911fbe-150f-11f0-a903-164fa3fc2eef','MacBook Air','Apple','Lightweight and fast.',1199.00,12,'Apple M2','8GB','256GB SSD','Integrated','2025-04-09 06:56:04'),('b49143da-150f-11f0-a903-164fa3fc2eef','Surface Laptop 5','Microsoft','Stylish productivity laptop.',1299.00,18,'Intel Core i5-1235U','8GB','512GB SSD','Intel Iris Xe','2025-04-09 06:56:04'),('b49144a6-150f-11f0-a903-164fa3fc2eef','Legion 5','Lenovo','Gaming laptop with solid specs.',1399.00,10,'AMD Ryzen 7 5800H','16GB','1TB SSD','NVIDIA RTX 3050 Ti','2025-04-09 06:56:04'),('b49144f0-150f-11f0-a903-164fa3fc2eef','ZenBook 13','Asus','Slim ultrabook with OLED display.',899.00,22,'Intel Core i7-1165G7','16GB','512GB SSD','Integrated','2025-04-09 06:56:04'),('b4914530-150f-11f0-a903-164fa3fc2eef','Omen 16','HP','Gaming laptop with powerful specs.',1599.00,9,'Intel Core i7-12700H','32GB','1TB SSD','NVIDIA RTX 3070','2025-04-09 06:56:04'),('b491456e-150f-11f0-a903-164fa3fc2eef','Nitro 5','Acer','Affordable gaming machine.',999.00,14,'AMD Ryzen 5 5600H','16GB','512GB SSD','NVIDIA GTX 1650','2025-04-09 06:56:04'),('b49145a9-150f-11f0-a903-164fa3fc2eef','MacBook Pro 16','Apple','Professional MacBook with high specs.',2499.00,7,'Apple M2 Pro','32GB','1TB SSD','Integrated','2025-04-09 06:56:04'),('b491471b-150f-11f0-a903-164fa3fc2eef','Latitude 5420','Dell','Business laptop with strong security.',1099.00,13,'Intel Core i5-1145G7','16GB','512GB SSD','Intel Iris Xe','2025-04-09 06:56:04'),('b4914a7c-150f-11f0-a903-164fa3fc2eef','EliteBook 840','HP','Enterprise-grade business laptop.',1249.00,10,'Intel Core i7-1255U','16GB','1TB SSD','Integrated','2025-04-09 06:56:04'),('b4916825-150f-11f0-a903-164fa3fc2eef','IdeaPad 3','Lenovo','Budget-friendly notebook.',399.00,40,'Intel Core i3-10110U','8GB','256GB SSD','Intel UHD','2025-04-09 06:56:04'),('b49168c5-150f-11f0-a903-164fa3fc2eef','TUF Dash F15','Asus','Gaming laptop with durable design.',1099.00,11,'Intel Core i7-11370H','16GB','512GB SSD','NVIDIA RTX 3060','2025-04-09 06:56:04'),('b4916905-150f-11f0-a903-164fa3fc2eef','Predator Helios 300','Acer','High-performance gaming laptop.',1499.00,6,'Intel Core i7-11800H','16GB','1TB SSD','NVIDIA RTX 3070','2025-04-09 06:56:04'),('b4916943-150f-11f0-a903-164fa3fc2eef','XPS 13','Dell','Premium ultrabook with InfinityEdge display.',1299.00,10,'Intel Core i5-1240P','16GB','512GB SSD','Intel Iris Xe','2025-04-09 06:56:04'),('b4916b4d-150f-11f0-a903-164fa3fc2eef','MateBook X Pro','Huawei','Sleek and stylish ultrabook.',1399.00,8,'Intel Core i7-1165G7','16GB','1TB SSD','Integrated','2025-04-09 06:56:04'),('b4916baa-150f-11f0-a903-164fa3fc2eef','Galaxy Book 3','Samsung','Portable and modern design.',999.00,17,'Intel Core i5-1240P','8GB','512GB SSD','Integrated','2025-04-09 06:56:04'),('b4916c73-150f-11f0-a903-164fa3fc2eef','Gram 16','LG','Ultra-lightweight 16-inch laptop.',1199.00,12,'Intel Core i7-1260P','16GB','1TB SSD','Intel Iris Xe','2025-04-09 06:56:04'),('b4916cd2-150f-11f0-a903-164fa3fc2eef','Swift X','Acer','Content creation laptop.',899.00,14,'AMD Ryzen 7 5800U','16GB','512GB SSD','NVIDIA RTX 3050','2025-04-09 06:56:04'),('b4916d14-150f-11f0-a903-164fa3fc2eef','Vivobook 15','Asus','Affordable and feature-rich.',599.00,26,'Intel Core i5-1135G7','8GB','512GB SSD','Integrated','2025-04-09 06:56:04'),('b4916d4c-150f-11f0-a903-164fa3fc2eef','Vostro 14','Dell','Small business laptop.',749.00,15,'Intel Core i3-1215U','8GB','256GB SSD','Intel UHD','2025-04-09 06:56:04'),('b4916eb8-150f-11f0-a903-164fa3fc2eef','ProBook 450','HP','Durable business laptop.',899.00,20,'Intel Core i5-1235U','16GB','512GB SSD','Intel Iris Xe','2025-04-09 06:56:04'),('b4916f1c-150f-11f0-a903-164fa3fc2eef','ThinkBook 14s','Lenovo','Modern business device.',949.00,17,'Intel Core i5-1135G7','16GB','512GB SSD','Integrated','2025-04-09 06:56:04'),('b4917cfa-150f-11f0-a903-164fa3fc2eef','ROG Strix G15','Asus','RGB gaming beast.',1699.00,9,'AMD Ryzen 9 5900HX','32GB','1TB SSD','NVIDIA RTX 3080','2025-04-09 06:56:04'),('b4917d92-150f-11f0-a903-164fa3fc2eef','MacBook Air 15','Apple','Larger screen MacBook.',1399.00,13,'Apple M2','8GB','512GB SSD','Integrated','2025-04-09 06:56:04'),('b4917ddc-150f-11f0-a903-164fa3fc2eef','Surface Laptop Studio','Microsoft','2-in-1 creative laptop.',1899.00,6,'Intel Core i7-11370H','32GB','1TB SSD','NVIDIA RTX A2000','2025-04-09 06:56:04'),('b4917e1a-150f-11f0-a903-164fa3fc2eef','Envy x360','HP','Convertible touch-screen laptop.',1099.00,11,'AMD Ryzen 7 5700U','16GB','1TB SSD','Integrated','2025-04-09 06:56:04'),('b4917e58-150f-11f0-a903-164fa3fc2eef','Yoga Slim 7','Lenovo','Slim and fast laptop.',1099.00,15,'AMD Ryzen 7 5800U','16GB','512GB SSD','Integrated','2025-04-09 06:56:04'),('b4917e91-150f-11f0-a903-164fa3fc2eef','Zenbook Flip','Asus','2-in-1 convertible ultrabook.',999.00,18,'Intel Core i5-1240P','8GB','512GB SSD','Integrated','2025-04-09 06:56:04'),('b4917ec8-150f-11f0-a903-164fa3fc2eef','Aspire Vero','Acer','Eco-friendly laptop.',699.00,22,'Intel Core i7-1195G7','16GB','512GB SSD','Intel Iris Xe','2025-04-09 06:56:04'),('b4917eff-150f-11f0-a903-164fa3fc2eef','Latitude 9430','Dell','High-end business device.',1799.00,7,'Intel Core i7-1265U','32GB','1TB SSD','Integrated','2025-04-09 06:56:04'),('b4917f36-150f-11f0-a903-164fa3fc2eef','ZBook Firefly','HP','Mobile workstation.',1999.00,5,'Intel Core i7-1185G7','32GB','1TB SSD','NVIDIA T500','2025-04-09 06:56:04'),('b4917f6d-150f-11f0-a903-164fa3fc2eef','Legion Slim 7','Lenovo','Slim gaming performance.',1699.00,8,'AMD Ryzen 9 5900HX','32GB','1TB SSD','NVIDIA RTX 3060','2025-04-09 06:56:04'),('b4918031-150f-11f0-a903-164fa3fc2eef','ROG Flow Z13','Asus','Gaming tablet with keyboard.',1899.00,6,'Intel Core i9-12900H','16GB','1TB SSD','NVIDIA RTX 3050 Ti','2025-04-09 06:56:04'),('b491808a-150f-11f0-a903-164fa3fc2eef','Swift 3 OLED','Acer','Affordable OLED screen.',799.00,20,'Intel Core i5-1240P','8GB','512GB SSD','Intel Iris Xe','2025-04-09 06:56:04'),('b49180c7-150f-11f0-a903-164fa3fc2eef','XPS 17','Dell','Large screen ultrabook.',2199.00,5,'Intel Core i7-12700H','32GB','1TB SSD','NVIDIA RTX 3050','2025-04-09 06:56:04'),('efdcd7de-cb80-416b-8c47-995e3681cccb','Office Laptop B2','HP','Budget laptop for office tasks',750.00,20,'Intel Core i5','8GB','512GB SSD','Intel UHD Graphics','2025-04-09 02:10:44');
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-04-11 15:43:51
