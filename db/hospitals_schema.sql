
CREATE TABLE hospital
(
  id                   INTEGER       NULL    ,
  name                 VARCHAR(255)  NOT NULL,
  owner_name           VARCHAR(255)  NOT NULL,
  address              VARCHAR(255)  NOT NULL,
  postal_code          VARCHAR(6)    NOT NULL,
  contact_phone_number VARCHAR(13)   NULL    ,
  symptoms             VARCHAR(1200) NOT NULL,
  PRIMARY KEY (id AUTOINCREMENT)
);

CREATE TABLE hospital_facility
(
  id          INTEGER NOT NULL,
  parking_lot TINYINT NOT NULL DEFAULT 0,
  PRIMARY KEY (id),
  FOREIGN KEY (id) REFERENCES hospital (id)
);

CREATE TABLE hospital_handle_symptom
(
  id          INTEGER NULL    ,
  hospital_id INTEGER NOT NULL,
  symptom_id  INTEGER NOT NULL,
  PRIMARY KEY (id AUTOINCREMENT),
  FOREIGN KEY (hospital_id) REFERENCES hospital (id),
  FOREIGN KEY (symptom_id) REFERENCES symptom (id)
);

CREATE TABLE hospital_open_date
(
  id          INTEGER NULL    ,
  hospital_id INTEGER NOT NULL,
  dow         INT     NOT NULL,
  start_time  TIME    NULL    ,
  end_time    TIME    NULL    ,
  PRIMARY KEY (id AUTOINCREMENT),
  FOREIGN KEY (hospital_id) REFERENCES hospital (id)
);

CREATE TABLE hospital_review_stat
(
  id             INTEGER NOT NULL,
  average_rating FLOAT   NOT NULL DEFAULT 0.0,
  total_rating   INT     NOT NULL DEFAULT 0,
  review_count   INT     NOT NULL DEFAULT 0,
  dx_rating      FLOAT   NOT NULL DEFAULT 0.0,
  PRIMARY KEY (id),
  FOREIGN KEY (id) REFERENCES hospital (id)
);

CREATE TABLE symptom
(
  id   INTEGER      NULL    ,
  name VARCHAR(255) NOT NULL,
  PRIMARY KEY (id AUTOINCREMENT)
);
