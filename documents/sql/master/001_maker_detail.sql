CREATE TABLE `maker_detail` (
  `maker_id` bigint NOT NULL,
  `ohp_url` varchar(255) NOT NULL,
  `twitter_name` varchar(255) NOT NULL,
  `youtube_channel_id` varchar(255) NOT NULL,
  `youtube_keywords` varchar(255) NOT NULL,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  PRIMARY KEY (maker_id)
)
