-- Create "todo_table" table
CREATE TABLE `todo_table` (
  `id` integer NULL,
  `title` varchar NOT NULL,
  PRIMARY KEY (`id`)
);
-- Create "todo_entry" table
CREATE TABLE `todo_entry` (
  `id` integer NULL,
  `content` varchar NOT NULL,
  `is_done` boolean NULL,
  `table_id` integer NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `0` FOREIGN KEY (`table_id`) REFERENCES `todo_table` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
);
