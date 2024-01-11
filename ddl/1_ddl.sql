DROP TABLE IF EXISTS user;

DROP TABLE IF EXISTS task;

CREATE TABLE task (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `title` VARCHAR(128) NOT NULL,
    `status` VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL
) Engine = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE user (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(20) NOT NULL,
    `password` VARCHAR(80) NOT NULL,
    `role` VARCHAR(80) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL
) Engine = InnoDB DEFAULT CHARSET = utf8mb4;