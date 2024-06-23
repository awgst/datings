CREATE TABLE `users` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `password_hash` varchar(255) NOT NULL,
  `created_at` timestamp,
  `updated_at` timestamp,
  `last_login_at` timestamp
);

CREATE TABLE `profiles` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` int,
  `name` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `premiums` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` int,
  `feature` ENUM ('verified_label', 'no_swipe_quota'),
  `created_at` timestamp
);

CREATE TABLE `swipes` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `profile_id` int,
  `user_id` int,
  `type` ENUM ('pass', 'like'),
  `created_at` timestamp
);

ALTER TABLE `profiles` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `premiums` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `swipes` ADD FOREIGN KEY (`profile_id`) REFERENCES `profiles` (`id`);

ALTER TABLE `swipes` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
