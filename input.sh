#! /bin/bash
# PARAMETERS
year=$1
day=$2

### YEAR ###
if [ -z "$year" ]
then
    echo "Please provide year as argument"
    exit 1
fi
if ! [[ $year =~ ^[0-9]+$ ]]
then
    echo "Year must be a number"
    exit 1
fi
if [ "$year" -lt 2015 ]
then
    echo "Year must be greater than 2015"
    exit 1
fi
if [ ! -d "$year" ]
then
    echo "Directory $year does not exist"
fi

### DAY ###
if [ -z "$day" ]
then
    echo "Please provide day as argument"
    exit 1
fi
if ! [[ $day =~ ^[0-9]+$ ]]
then
    echo "Day must be a number"
    exit 1
fi
if [ "$day" -lt 1 ] || [ "$day" -gt 25 ]
then
    echo "Day must be between 1 and 25"
    exit 1
fi
if [ "$day" -lt 10 ]
then
    dir_name="day0$day"
else
    dir_name="day$day"
fi
if [ ! -d "$year/src/$dir_name" ]
then
    echo "Directory $year/src/$dir_name does not exist"
fi

### INPUT ###
export "$(grep -v '^#' .env | grep '^AOC_COOKIE' | xargs)"
echo "$AOC_COOKIE"
curl --cookie session="$AOC_COOKIE" https://adventofcode.com/"$year"/day/"$day"/input > "$year"/src/"$dir_name"/input.txt