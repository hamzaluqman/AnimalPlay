CREATE TABLE animal
(
    id INT NOT NULL AUTO_INCREMENT,
    dateCreated DATETIME NOT NULL,
    dateUpdated DATETIME NOT NULL,
    animalName VARCHAR(255) NOT NULL,
    animalType VARCHAR(255) NOT NULL, 
    PRIMARY KEY (id)
);