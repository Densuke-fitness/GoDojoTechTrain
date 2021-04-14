CREATE TABLE IF NOT EXISTS `users`
(
    `id`                  BIGINT(20) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `name`                VARCHAR(256) NOT NULL,
    `created_at`          DATETIME DEFAULT CURRENT_TIMESTAMP, 
    `updated_at`          DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;

