DROP DATABASE IF EXISTS cosmetics_store;
CREATE DATABASE IF NOT EXISTS cosmetics_store;

USE cosmetics_store;

DROP TABLE IF EXISTS `products`;
CREATE TABLE `products`(
	`id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `product_name` VARCHAR(100) NOT NULL,
    `description` TEXT NOT NULL,
    `price` BIGINT NOT NULL,
    `image` VARCHAR(200) NOT NULL,
    `is_hot` BOOLEAN NOT NULL,
    `category_id` INT NOT NULL,
    `brand_id` INT NOT NULL,
    `number_available` INT NOT NULL
);

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`(
	`id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `full_name` VARCHAR(100) NOT NULL,
    `email` VARCHAR(191) NOT NULL,
    `password` LONGTEXT NOT NULL,
    `phone` VARCHAR(12) NOT NULL,
    `address` LONGTEXT NOT NULL,
    `date_of_birth` TEXT NOT NULL,
    `gender` ENUM('Male', 'Female'),
    `avatar` VARCHAR(200) NOT NULL,
    `role_id` INT NOT NULL
);

DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles`(
	`id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(100) NOT NULL
);

DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories`(
	`id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `category_name` VARCHAR(100) NOT NULL
);

DROP TABLE IF EXISTS `brands`;
CREATE TABLE `brands`(
	`id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `brand_name` VARCHAR(100) NOT NULL
);

DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders`(
	`id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT NOT NULL,
    `total` INT NOT NULL
);

DROP TABLE IF EXISTS `order_items`;
CREATE TABLE `order_items`(
	`id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `order_id` INT NOT NULL,
    `product_id` INT NOT NULL,
    `quantity` INT NOT NULL
);

ALTER TABLE `products` ADD FOREIGN KEY (`category_id`) REFERENCES `categories`(`id`);
ALTER TABLE `products` ADD FOREIGN KEY (`brand_id`) REFERENCES `brands`(`id`);
ALTER TABLE `users` ADD FOREIGN KEY (`role_id`) REFERENCES `roles`(`id`);
ALTER TABLE `orders` ADD FOREIGN KEY (`user_id`) REFERENCES `users`(`id`);
ALTER TABLE `order_items` ADD FOREIGN KEY (`order_id`) REFERENCES `orders`(`id`);
ALTER TABLE `order_items` ADD FOREIGN KEY (`product_id`) REFERENCES `products`(`id`);






