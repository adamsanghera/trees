'''
This file cleans the 2005 csv, which is egregiously dirty.
'''

with open('tree-census/new_york_tree_census_2005.csv') as f:
  with open('tree-census/new_york_tree_census_2005_cleaned.csv', 'w') as new_f:
    # Get rid of the header
    header = f.readline().rstrip()

    new_f.write('{},lat,lon\n'.format(header))

    while True:
      first = f.readline().rstrip()
      second = f.readline().rstrip()
      third = f.readline().rstrip().split(' ')

      if second != 'New York':
        break

      new_f.write('{},{},{}\n'.format(first, third[0][1:], third[1][:len(third[1])-1]))
