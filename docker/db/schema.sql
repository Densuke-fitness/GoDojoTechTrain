CREATE TABLE IF NOT EXISTS `users`
(
    `id`                  BIGINT(20) UNSIGNED AUTO_INCREMENT,
    `name`                VARCHAR(256) NOT NULL,
    `created_at`          DATETIME DEFAULT CURRENT_TIMESTAMP, 
    `updated_at`          DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id)

) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;


CREATE TABLE IF NOT EXISTS `characters_master`
(
    `id`                  BIGINT(20) UNSIGNED,
    `name`                VARCHAR(50) NOT NULL,
    PRIMARY KEY(id)
) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;

-- Insert the default data.
INSERT INTO characters_master (id, name) VALUES (1, "Python");
INSERT INTO characters_master (id, name) VALUES (2, "Golang");
INSERT INTO characters_master (id, name) VALUES (3, "Jupyter");
INSERT INTO characters_master (id, name) VALUES (4, "C#");
INSERT INTO characters_master (id, name) VALUES (5, "JavaScript");

CREATE TABLE IF NOT EXISTS `events_master`
(
    `id` BIGINT(20) UNSIGNED,
    `description` VARCHAR(255),
    PRIMARY KEY(id)
) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;

-- Insert the default data.
INSERT INTO events_master (id, description) VALUES (1, "normal event: This period is indefinite.");

CREATE TABLE IF NOT EXISTS `characters_lottery_rate`
(
    `event_id`            BIGINT(20) UNSIGNED,
    `character_id`        BIGINT(20) UNSIGNED,
    `lottery_rate`        FLOAT NOT NULL DEFAULT 0.0,
    FOREIGN KEY(event_id) REFERENCES events_master(id),
    FOREIGN KEY(character_id) REFERENCES characters_master(id)
) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;

-- Insert the default data.
INSERT INTO characters_lottery_rate (event_id, character_id, lottery_rate) VALUES (1, 1, 0.3);
INSERT INTO characters_lottery_rate (event_id, character_id, lottery_rate) VALUES (1, 2, 0.4);
INSERT INTO characters_lottery_rate (event_id, character_id, lottery_rate) VALUES (1, 3, 0.1);
INSERT INTO characters_lottery_rate (event_id, character_id, lottery_rate) VALUES (1, 4, 0.2);
INSERT INTO characters_lottery_rate (event_id, character_id, lottery_rate) VALUES (1, 5, 0.1);

CREATE TABLE IF NOT EXISTS `possession_characters`
(
    `user_id`       BIGINT(20) UNSIGNED,
    `character_id`  BIGINT(20) UNSIGNED,
    `character_seq` BIGINT(20) UNSIGNED NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(character_id) REFERENCES characters_master(id)
) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;