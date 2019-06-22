CREATE TABLE entry (
    id BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    author_id BIGINT UNSIGNED NOT NULL,
    keyword VARCHAR(191) UNIQUE,
    description MEDIUMTEXT,
    len SMALLINT,
    updated_at DATETIME NOT NULL,
    created_at DATETIME NOT NULL,
    KEY len_idx_on_entry(len)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE user (
    id BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    name VARCHAR(191) UNIQUE,
    salt VARCHAR(20),
    password VARCHAR(40),
    created_at DATETIME NOT NULL
) Engine=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE star (
    id BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    keyword VARCHAR(191) NOT NULL,
    user_name VARCHAR(191) NOT NULL,
    created_at DATETIME
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

UPDATE entry e1, entry e2 SET e1.len = character_length(e2.description) where e1.id = e2.id
