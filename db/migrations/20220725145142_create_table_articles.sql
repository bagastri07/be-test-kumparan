-- migrate:up
create table
  `articles` (
    `id` char(36) not null,
    `author_id` char(36) not null,
    `title` VARCHAR(255) not null,
    `body` TEXT not null,
    `created_at` TIMESTAMP not null default CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP null,
    primary key (`id`)
  )

-- migrate:down
drop table `articles`;