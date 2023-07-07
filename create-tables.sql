CREATE TABLE Rides (
    form_id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100),
    dep_city VARCHAR(100),
    dep_date DATE,
    dep_time TIME,
    arr_date VARCHAR(100),
    arr_time TIME,
    flight_num VARCHAR(100)
);
