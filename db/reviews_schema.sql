
CREATE TABLE photo
(
  id        INTEGER      NULL    ,
  review_id INTEGER      NOT NULL,
  photo_url VARCHAR(255) NOT NULL,
  PRIMARY KEY (id AUTOINCREMENT),
  FOREIGN KEY (review_id) REFERENCES review (id)
);

CREATE TABLE review
(
  id          INTEGER      NULL    ,
  user_id     INTEGER      NOT NULL,
  hospital_id INTEGER      NOT NULL,
  timestamp   TIMESTAMP    NOT NULL,
  score       TINYINT      NOT NULL,
  context     VARCHAR(600) NULL    ,
  PRIMARY KEY (id AUTOINCREMENT)
);
