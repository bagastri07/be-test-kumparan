-- migrate:up
ALTER TABLE `articles`
  ADD CONSTRAINT `articles_fk` FOREIGN KEY (`author_id`) REFERENCES `authors` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION;

-- migrate:down
ALTER TABLE `articles` DROP FOREIGN KEY articles_fk;
