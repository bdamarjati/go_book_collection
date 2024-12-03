CREATE TABLE `users` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(255) UNIQUE,
  `role` varchar(255),
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `books` (
  `book_id` integer PRIMARY KEY AUTO_INCREMENT,
  `collection_id` integer NOT NULL,
  `title` text,
  `author` varchar(255),
  `language` varchar(255),
  `year_published` integer,
  `ISBN` varchar(255),
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `collections` (
  `collection_id` integer PRIMARY KEY AUTO_INCREMENT,
  `user_id` integer NOT NULL,
  `name` varchar(255) UNIQUE,
  `status` integer,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE `collections` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `books` ADD FOREIGN KEY (`collection_id`) REFERENCES `collections` (`collection_id`);
