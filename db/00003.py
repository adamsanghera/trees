'''
This file creates the create table statements, which are used in 00004.sql
'''

import csv

print("CREATE TABLE tree_1995_intermediary")
with open('tree-census/new_york_tree_census_1995.csv') as f:
  for field in f.readline().split(','):
    print("{} TEXT,".format(field.rstrip()))

print("CREATE TABLE tree_2005_intermediary")
with open('tree-census/new_york_tree_census_2005_cleaned.csv') as f:
  for field in f.readline().split(','):
    print("{} TEXT,".format(field.rstrip()))

print("CREATE TABLE tree_2015_intermediary")
with open('tree-census/new_york_tree_census_2015.csv') as f:
  for field in f.readline().split(','):
    print("{} TEXT,".format(field.rstrip()))
