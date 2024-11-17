
CREATE TABLE medical_record
(
  id          INTEGER       NULL    ,
  user_id     INTEGER       NOT NULL,
  hospital_id INT           NOT NULL,
  timestamp   TIMESTAMP     NOT NULL,
  doctor_name VARCHAR(20)   NOT NULL,
  notes       VARCHAR(1200) NULL    ,
  symptom     VARCHAR(1200) NULL    ,
  PRIMARY KEY (id AUTOINCREMENT),
  FOREIGN KEY (user_id) REFERENCES user (id)
);

CREATE TABLE payment_method
(
  id               INTEGER     NULL    ,
  user_id          INTEGER     NOT NULL,
  card_holder_name VARCHAR(30) NOT NULL,
  card_number      VARCHAR(19) NOT NULL,
  exp_date         VARCHAR(5)  NOT NULL,
  cvv              VARCHAR(4)  NOT NULL,
  PRIMARY KEY (id AUTOINCREMENT),
  FOREIGN KEY (user_id) REFERENCES user (id)
);

CREATE TABLE priority_set
(
  id         INTEGER NULL    ,
  priority_1 INT     NULL    ,
  priority_2 INT     NULL    ,
  priority_3 INT     NULL    ,
  priority_4 INT     NULL    ,
  priority_5 INT     NULL    ,
  PRIMARY KEY (id AUTOINCREMENT)
);

CREATE TABLE review
(
  id           INTEGER      NULL    ,
  user_id      INTEGER      NOT NULL,
  timestamp    TIMESTAMP    NOT NULL,
  score        INT          NOT NULL,
  text_content VARCHAR(600) NULL    ,
  PRIMARY KEY (id AUTOINCREMENT),
  FOREIGN KEY (user_id) REFERENCES user (id)
);

CREATE TABLE review_photo
(
  id        INTEGER      NULL    ,
  review_id INTEGER      NOT NULL,
  photo_url VARCHAR(255) NOT NULL,
  PRIMARY KEY (id AUTOINCREMENT),
  FOREIGN KEY (review_id) REFERENCES review (id)
);

CREATE TABLE user
(
  id              INTEGER      NULL    ,
  email_address   VARCHAR(255) NOT NULL,
  hashed_password VARCHAR(255) NOT NULL,
  PRIMARY KEY (id AUTOINCREMENT)
);

CREATE TABLE user_profile
(
  id                INTEGER      NULL    ,
  name              VARCHAR(20)  NOT NULL,
  profile_image_url VARCHAR(255) NULL     DEFAULT NULL,
  phone_number      VARCHAR(13)  NOT NULL,
  home_address      VARCHAR(255) NOT NULL,
  postal_code       VARCHAR(6)   NOT NULL,
  candy             INT          NOT NULL DEFAULT 0,
  priority_id       INTEGER      NOT NULL DEFAULT 0,
  PRIMARY KEY (id AUTOINCREMENT),
  FOREIGN KEY (priority_id) REFERENCES priority_set (id),
  FOREIGN KEY (id) REFERENCES user (id)
);
