CREATE DATABASE IF NOT EXISTS webbeer;
use webbeer;

DROP TABLE IF EXISTS beer;

CREATE TABLE beer (
   id INTEGER PRIMARY KEY AUTO_INCREMENT,
   name text NOT NULL,
   type integer NOT NULL,
   style integer not null
) ENGINE=InnoDB;