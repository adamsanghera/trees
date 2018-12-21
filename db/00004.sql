-- Create the initial tables for all of the census data

CREATE TABLE tree_1995_intermediary (
  recordid TEXT,
  address TEXT,
  house_number TEXT,
  street TEXT,
  zip_original TEXT,
  cb_original TEXT,
  site TEXT,
  species TEXT,
  diameter TEXT,
  status TEXT,
  wires TEXT,
  sidewalk_condition TEXT,
  support_structure TEXT,
  borough TEXT,
  x TEXT,
  y TEXT,
  longitude TEXT,
  latitude TEXT,
  cb_new TEXT,
  zip_new TEXT,
  censustract_2010 TEXT,
  censusblock_2010 TEXT,
  nta_2010 TEXT,
  segmentid TEXT,
  spc_common TEXT,
  spc_latin TEXT,
  location TEXT
);

CREATE TABLE tree_2005_intermediary (
  objectid TEXT,
  cen_year TEXT,
  tree_dbh TEXT,
  tree_loc TEXT,
  pit_type TEXT,
  soil_lvl TEXT,
  status TEXT,
  spc_latin TEXT,
  spc_common TEXT,
  vert_other TEXT,
  vert_pgrd TEXT,
  vert_tgrd TEXT,
  vert_wall TEXT,
  horz_blck TEXT,
  horz_grate TEXT,
  horz_plant TEXT,
  horz_other TEXT,
  sidw_crack TEXT,
  sidw_raise TEXT,
  wire_htap TEXT,
  wire_prime TEXT,
  wire_2nd TEXT,
  wire_other TEXT,
  inf_canopy TEXT,
  inf_guard TEXT,
  inf_wires TEXT,
  inf_paving TEXT,
  inf_outlet TEXT,
  inf_shoes TEXT,
  inf_lights TEXT,
  inf_other TEXT,
  trunk_dmg TEXT,
  zipcode TEXT,
  zip_city TEXT,
  cb_num TEXT,
  borocode TEXT,
  boroname TEXT,
  cncldist TEXT,
  st_assem TEXT,
  st_senate TEXT,
  nta TEXT,
  nta_name TEXT,
  boro_ct TEXT,
  x_sp TEXT,
  y_sp TEXT,
  objectid_1 TEXT,
  location_1 TEXT,
  lat TEXT,
  lon TEXT
);

CREATE TABLE tree_2015_intermediary (
  tree_id TEXT,
  block_id TEXT,
  created_at TEXT,
  tree_dbh TEXT,
  stump_diam TEXT,
  curb_loc TEXT,
  status TEXT,
  health TEXT,
  spc_latin TEXT,
  spc_common TEXT,
  steward TEXT,
  guards TEXT,
  sidewalk TEXT,
  user_type TEXT,
  problems TEXT,
  root_stone TEXT,
  root_grate TEXT,
  root_other TEXT,
  trunk_wire TEXT,
  trnk_light TEXT,
  trnk_other TEXT,
  brch_light TEXT,
  brch_shoe TEXT,
  brch_other TEXT,
  address TEXT,
  zipcode TEXT,
  zip_city TEXT,
  cb_num TEXT,
  borocode TEXT,
  boroname TEXT,
  cncldist TEXT,
  st_assem TEXT,
  st_senate TEXT,
  nta TEXT,
  nta_name TEXT,
  boro_ct TEXT,
  state TEXT,
  latitude TEXT,
  longitude TEXT,
  x_sp TEXT,
  y_sp TEXT
);

-- Copy over the values themselves

\copy tree_1995_intermediary FROM './tree-census/new_york_tree_census_1995.csv' DELIMITER ',' CSV;

\copy tree_2005_intermediary FROM './tree-census/new_york_tree_census_2005_cleaned.csv' DELIMITER ',' CSV;

\copy tree_2015_intermediary FROM './tree-census/new_york_tree_census_2015.csv' DELIMITER ',' CSV;

-- Delete the fake data

DELETE FROM tree_1995_intermediary
  WHERE address='address';

DELETE FROM tree_2005_intermediary 
  WHERE cen_year='cen_year';

DELETE FROM tree_2015_intermediary
  WHERE block_id='block_id';


