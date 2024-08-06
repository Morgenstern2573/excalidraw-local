DROP TABLE IF EXISTS Collections;
DROP TABLE IF EXISTS Scenes;

CREATE TABLE Collections(
  Name TEXT, 
  ID TEXT PRIMARY KEY
);

CREATE TABLE Scenes(
  Name TEXT, 
  CollectionID TEXT, 
  ID TEXT,
  Data TEXT, 
  FOREIGN KEY(CollectionID) REFERENCES Collections(ID),
  PRIMARY KEY(Name, CollectionID)
);



