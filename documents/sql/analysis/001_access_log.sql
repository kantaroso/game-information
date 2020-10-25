CREATE TABLE `access_log` (
  `id` bigint AUTO_INCREMENT NOT NULL,
  `method` varchar(255) NOT NULL,
  `endpoint` varchar(255) NOT NULL,
  `query_string` text NOT NULL,
  `user_agent` varchar(255) NOT NULL,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  PRIMARY KEY (id)
)
