CREATE TABLE master.maker_video (
  id BIGSERIAL PRIMARY KEY, -- bigint AUTO_INCREMENT NOT NULL, PRIMARY KEY (id) は PostgreSQLでは BIGSERIAL PRIMARY KEY
  maker_id BIGINT NOT NULL,
  video_id VARCHAR(255) NOT NULL UNIQUE, -- UNIQUE `video_id_uniq` (`video_id`) はカラム定義で UNIQUE を追加
  title VARCHAR(255) NOT NULL,
  published_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL, -- TIMESTAMP WITH TIME ZONE を推奨
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- KEY `maker_id_idx` (`maker_id`) に相当するインデックスの作成
CREATE INDEX maker_id_idx ON master.maker_video (maker_id);