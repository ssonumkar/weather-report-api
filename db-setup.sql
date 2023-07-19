CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    dob DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS WeatherHistory (
    id INT AUTO_INCREMENT PRIMARY KEY,
	user_id INT NOT NULL,
    city VARCHAR(255) NOT NULL,
    temperature_min FLOAT,
	temperature_max FLOAT,
    feels_like FLOAT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Users Registration ----------------------------
DROP PROCEDURE IF EXISTS RegisterUser;
delimiter //
CREATE PROCEDURE RegisterUser(
    username VARCHAR(255),
    password VARCHAR(255),
    dateOfBirth DATE
)
BEGIN
    INSERT INTO Users (username, password, dob)
    VALUES (username, password, dateOfBirth);
END//

-- Users Login: ----------------------------------
DROP PROCEDURE IF EXISTS LoginUser;
delimiter //
CREATE PROCEDURE LoginUser(
    username VARCHAR(255),
    password VARCHAR(255)
)
BEGIN
    SELECT id
    FROM Users
    WHERE username = username AND password = password;
END//

-- Adding Weather History: ------------------
DROP PROCEDURE IF EXISTS AddWeatherHistory;
delimiter //
CREATE PROCEDURE AddWeatherHistory(
    userId INT,
    city VARCHAR(255),
    temperature_min DECIMAL(10,2),
	temperature_max DECIMAL(10,2),
    feels_like INT,
    created_at DATETIME
)
BEGIN
    INSERT INTO WeatherHistory (user_id, city, temperature_min, temperature_max, feels_like, created_at)
    VALUES (userId, city, temperature_min, temperature_max, feels_like, created_at);
END//

-- Deleting Weather History: ----------------------
DROP PROCEDURE IF EXISTS DeleteWeatherHistory;
delimiter //
CREATE PROCEDURE DeleteWeatherHistory(
    historyId INT
)
BEGIN
    DELETE FROM WeatherHistory
    WHERE id = historyId;
END//

-- Bulk Delete Weather history --------------------------
DROP PROCEDURE IF EXISTS BulkDeleteWeatherHistory;
delimiter //
CREATE PROCEDURE BulkDeleteWeatherHistory(
    historyIds VARCHAR(100)
)
BEGIN
    DECLARE sqlQuery NVARCHAR(1000);
    SET sqlQuery = 'DELETE FROM WeatherHistory WHERE id IN (' + historyIds + ')';
    EXECUTE sqlQuery;
END//

 