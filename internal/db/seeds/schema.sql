-- Create "products" table
-- "mysql://root:my-secret-pw@127.0.0.1:3307/boilerplate_go"
CREATE TABLE `products` (`id` int NOT NULL, `name` varchar(45) NOT NULL, `price` float NULL, PRIMARY KEY (`id`)) CHARSET utf8mb3 COLLATE utf8mb3_general_ci;
