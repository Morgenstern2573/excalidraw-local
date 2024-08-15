DROP TABLE IF EXISTS Drawings;
DROP TABLE IF EXISTS Collections;
DROP TABLE IF EXISTS Users;

CREATE TABLE Collections(
  Name TEXT UNIQUE, 
  ID TEXT PRIMARY KEY
);

CREATE TABLE Drawings(
  Name TEXT, 
  CollectionID TEXT, 
  ID TEXT,
  Data TEXT, 
  FOREIGN KEY(CollectionID) REFERENCES Collections(ID),
  PRIMARY KEY(Name, CollectionID)
);

CREATE TABLE Users(
  Name TEXT,
  Email TEXT PRIMARY KEY,
  PasswordHash TEXT NOT NULL
);