CREATE TABLE `users` (
  `id` integer PRIMARY KEY,
  `username` varchar(255) UNIQUE,
  `role` varchar(255),
  `created_at` timestamp
);

CREATE TABLE `books` (
  `book_id` integer PRIMARY KEY,
  `collection_id` integer,
  `title` text,
  `author` varchar(255),
  `language` varchar(255),
  `year_published` integer,
  `ISBN` varchar(255),
  `created_at` timestamp
);

CREATE TABLE `collections` (
  `collection_id` integer PRIMARY KEY,
  `user_id` integer,
  `name` varchar(255) UNIQUE,
  `status` integer,
  `created_at` timestamp
);

ALTER TABLE `collections` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `books` ADD FOREIGN KEY (`collection_id`) REFERENCES `collections` (`collection_id`);
