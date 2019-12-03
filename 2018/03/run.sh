psql -U postgres -h aoc2018day3 -c 'CREATE DATABASE aoc2018day3;'
psql -U postgres -h aoc2018day3 aoc2018day3 -f table.sql
psql -U postgres -h aoc2018day3 aoc2018day3 -f input.sql
psql -U postgres -h aoc2018day3 aoc2018day3 -f $SCRIPT
