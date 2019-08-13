-- +goose Up
CREATE TABLE tag_article (
  id int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  tag_id int(10) UNSIGNED NOT NULL,
  article_id int(10) UNSIGNED NOT NULL,
  ctime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  utime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE tag_article;