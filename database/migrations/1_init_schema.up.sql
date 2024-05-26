CREATE TABLE animal
(
    id INT NOT NULL AUTO_INCREMENT,
    dateCreated DATETIME NOT NULL,
    dateUpdated DATETIME NOT NULL,
    animalName VARCHAR(255) NOT NULL,
    animalType VARCHAR(255) NOT NULL, 
    PRIMARY KEY (id)
);

INSERT INTO animal(dateCreated, dateUpdated, animalName, animalType)
VALUES 
(UTC_TIMESTAMP(), UTC_TIMESTAMP(), "BouncyBunny", "Rabbit"),
(UTC_TIMESTAMP(), UTC_TIMESTAMP(), "GiddyGoat", "Goat"),
(UTC_TIMESTAMP(), UTC_TIMESTAMP(), "BaBaBlackSheep", "Sheep");