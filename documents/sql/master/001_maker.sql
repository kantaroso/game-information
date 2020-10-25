CREATE TABLE `maker` (
  `id` bigint AUTO_INCREMENT NOT NULL,
  `name` varchar(255) NOT NULL,
  `ohp_url` varchar(255) NOT NULL,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  PRIMARY KEY (id)
)
