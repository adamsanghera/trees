-- Set appropriate columns to NULL
-- UPDATE tree_2015_intermediary SET problems = NULL
--   WHERE problems='None' OR problems='';
-- 
-- UPDATE tree_2015_intermediary SET stump_diam = NULL
--   WHERE stump_diam='0';
-- 
-- UPDATE tree_2015_intermediary SET tree_dbh = NULL
--   WHERE tree_dbh='0';
-- 
-- UPDATE tree_2015_intermediary SET health = NULL
--   WHERE health='';
-- 
-- UPDATE tree_2015_intermediary SET spc_latin = NULL
--   WHERE spc_latin='';
-- 
-- UPDATE tree_2015_intermediary SET spc_common = NULL
--   WHERE spc_common='';
-- 
-- UPDATE tree_2015_intermediary SET steward = NULL
--   WHERE steward='' OR steward='None';
-- 
-- UPDATE tree_2015_intermediary SET guards = NULL
--   WHERE guards='';
-- 
-- UPDATE tree_2015_intermediary SET sidewalk = NULL
--   WHERE sidewalk='';
-- 

CREATE EXTENSION postgis;

CREATE TABLE Trees
AS (
  SELECT 
    cast(tree_id as integer), cast(created_at as timestamp), cast(tree_dbh as integer) as tree_diameter,
    cast(stump_diam as integer) as stump_diameter, status, health,
    spc_latin, spc_common, steward, curb_loc as curb_location, guards,
    sidewalk, user_type, problems,
    root_stone, root_grate, root_other,
    trnk_other as trunk_other, trunk_wire, trnk_light as trunk_light,
    brch_light as branch_light, brch_shoe as branch_shoe, brch_other as branch_other,
    address, cast(zipcode as integer), zip_city, boroname as borough_name,
    ST_SetSRID(ST_MakePoint(
      cast(latitude as double precision), cast(longitude as double precision)), 4326) as location
  FROM tree_2015_intermediary
);

select tree_id, address, spc_common, status, (ST_DistanceSphere(ST_SetSRID(St_MakePoint(40.697266, -73.992886), 4326), location) / 1609.344) as m
from trees where ST_DistanceSphere(ST_SetSRID(St_MakePoint(40.697266, -73.992886), 4326), location) / 1609.344 <= 0.5 order by m limit 15;