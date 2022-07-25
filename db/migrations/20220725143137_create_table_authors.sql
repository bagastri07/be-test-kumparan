-- migrate:up
CREATE TABLE
  `authors` (
    `id` CHAR(36) NOT NULL,
    `name` varchar(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY (`id`)
  )

-- migrate:down
DROP TABLE `authors`;