CREATE TABLE `maker_video` (
  `id` bigint AUTO_INCREMENT NOT NULL,
  `maker_id` bigint NOT NULL,
  `video_id` varchar(255) NOT NULL,
  `title` varchar(255) NOT NULL,
  `published_at` timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  PRIMARY KEY (id),
  KEY `maker_id_idx` (`maker_id`),
  UNIQUE `video_id_uniq` (`video_id`)
)
