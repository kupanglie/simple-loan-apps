CREATE TABLE `bank`.`loans` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id` int UNSIGNED NOT NULL,
    `amount` INT UNSIGNED NOT NULL,
    `period` INT UNSIGNED NOT NULL,
    `purpose` TEXT NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NULL,
    `deleted_at` DATETIME NULL
);

