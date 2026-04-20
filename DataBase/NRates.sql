CREATE EXTENSION postgis;

DROP TABLE IF EXISTS simulations;
DROP TABLE IF EXISTS spatial_data;
DROP TABLE IF EXISTS ON_FARM;

---------------------------
-- Table for simulations --
---------------------------
CREATE TABLE simulations(
    id_cell VARCHAR(64),
    id_within_cell INT,
    year INT,
    nitro_kg_ha INT,
    yield_kg_ha FLOAT,
    );


------------------------------
-- Table for on-farm trials --
------------------------------
CREATE TABLE on_farm(
    id_trial VARCHAR(64),
    id_region VARCHAR(64),
    aonr FLOAT,
    PRIMARY KEY(id_trial)
    );





