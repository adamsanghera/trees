# Trees

An application for documenting, researching, querying, and maintaining the trees of NYC's streets.

## Todo

Database side

1. Evaluate Graphql
1. Give up on graphql
1. Use postgresql w/ postGIS, for using our #geospatialData
1. Figure out how to programmatically load the tree census data into a postregsql instance that has postGIS enabled.  We'll probably need to transform the geospatial data into some specially named columns.

Backend side

1. Create an api server wrapping postgresql, to answer basic queries:
   1. Geospatial queries
      1. What are the closest trees, given this address?
      1. What are the closest trees, given this latitude/longitude?
      1. What's the closest tree of X specie?
   1. Aggregation queries
      1. How many trees are there of X specie?
      1. What is the most commonly-occurring tree specie in NYC?
      1. What is the most commonly-occurring tree specie in X borough?
      1. What is the most commonly-occurring tree specie on X street?
   1. Zoom+enhance queries
      1. Tell me more about this specific tree

Frontend side

1. Evaluate flutter
1. Just write the app!

