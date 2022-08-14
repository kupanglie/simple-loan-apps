CREATE TABLE `bank`.`users` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `identity_number` VARCHAR(20) NOT NULL UNIQUE,
    `date_of_birth` DATE NOT NULL,
    `sex` VARCHAR(6) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NULL,
    `deleted_at` DATETIME NULL
);

