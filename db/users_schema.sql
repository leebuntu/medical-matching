
CREATE TABLE medical_record
(
  id          INTEGER       NULL    ,
  user_id     INTEGER       NOT NULL,
  hospital_id INTEGER       NOT NULL,
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

CREATE TABLE priority
(
  id          INTEGER NULL    ,
  user_id     INTEGER NOT NULL,
  priority_id INT     NOT NULL,
  rank        INT     NOT NULL,
  PRIMARY KEY (id AUTOINCREMENT),
  FOREIGN KEY (user_id) REFERENCES user (id)
);

CREATE TABLE user
(
  id              INTEGER      NULL    ,
  email           VARCHAR(255) NOT NULL,
  hashed_password VARCHAR(255) NOT NULL,
  PRIMARY KEY (id AUTOINCREMENT)
);

CREATE TABLE user_profile
(
  id                INTEGER      NOT NULL,
  name              VARCHAR(20)  NOT NULL,
  profile_image_url VARCHAR(255) NULL     DEFAULT NULL,
  phone_number      VARCHAR(20)  NOT NULL,
  home_address      VARCHAR(255) NOT NULL,
  postal_code       VARCHAR(6)   NOT NULL,
  candy             INT          NOT NULL DEFAULT 0,
  card_id           INTEGER      NULL    ,
  PRIMARY KEY (id),
  FOREIGN KEY (id) REFERENCES user (id),
  FOREIGN KEY (card_id) REFERENCES payment_method (id)
);
