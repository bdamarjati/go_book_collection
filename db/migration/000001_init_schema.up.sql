CREATE TABLE `users` (
  `username` varchar(255) PRIMARY KEY,
  `password` varchar(255) NOT NULL,
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
  `user` varchar(255) NOT NULL,
  `name` varchar(255) UNIQUE,
  `status` integer,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE `collections` ADD FOREIGN KEY (`user`) REFERENCES `users` (`username`);

ALTER TABLE `books` ADD FOREIGN KEY (`collection_id`) REFERENCES `collections` (`collection_id`);
