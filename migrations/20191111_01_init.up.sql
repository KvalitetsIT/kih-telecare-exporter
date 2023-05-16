CREATE TABLE IF NOT EXISTS measurements (
  id varchar(100) UNIQUE NOT NULL,
  measurement varchar(256) UNIQUE NOT NULL,
  patient varchar(256) NOT NULL,
  status int,
  backend_status int,
  backend_reply varchar(256),
  created_at datetime,
  updated_at datetime,

  PRIMARY KEY(id),
  INDEX(measurement),
  INDEX(patient)
);

CREATE TABLE IF NOT EXISTS runstatus (
  id varchar(100) UNIQUE NOT NULL,
  lastrun datetime,
  status int,
  created_at datetime,

  PRIMARY KEY(id),
  INDEX(id)
);
