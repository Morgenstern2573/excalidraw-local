DROP TABLE IF EXISTS Drawings;
DROP TABLE IF EXISTS Collections;

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