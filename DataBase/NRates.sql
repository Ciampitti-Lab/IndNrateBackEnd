CREATE EXTENSION postgis;

DROP TABLE IF EXISTS simulations;
DROP TABLE IF EXISTS spatial_data;

---------------------------
-- Table for simulations --
---------------------------
CREATE TABLE simulations(
    id_sim VARCHAR(64),
    id_cell INT,
    nitrogen INT,
    yield FLOAT,
    PRIMARY KEY(id_sim)
    );




