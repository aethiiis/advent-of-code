#! /bin/bash
# PARAMETERS
year=$1
day=$2

### YEAR ###
# Check if year is provided
if [ -z $year ]
then
    echo "Please provide year as argument"
    exit 1
fi
# Check if year is a number
if ! [[ $year =~ ^[0-9]+$ ]]
then
    echo "Year must be a number"
    exit 1
fi
# Check if year is less than 2015
if [ $year -lt 2015 ]
then
    echo "Year must be greater than 2015"
    exit 1
fi
# Check if year directory exists
if [ ! -d $year ]
then
    mkdir $year
    echo "Directory $year created"
    
fi
# Check if src directory exists
if [ ! -d "$year/src" ]
then
    mkdir "$year/src"
    echo "Directory $year/src created"
fi
# Check if properties file exists
if [ ! -f "$year/properties" ]
then
    touch "$year/properties"
    echo "File $year/properties created"
    echo -n "Specify the file extension (e.g. py, js, java): "
    read EXTENSION
    echo "EXTENSION=$EXTENSION" >> "$year/properties"
    echo "Information saved in $year/properties"
fi
cd "$year/src"

### UTILS ###
# Check if utils directory exists
if [ ! -d "utils" ]
then
    mkdir utils
    echo "Directory utils created"
    exit 0
fi
## Check if template files exist
if ! ls utils/template* utils/test_template* 1> /dev/null 2>&1
then
    echo "You need to create template files in utils directory"
    exit 0
fi

### DAY ###
# Check if day is provided
if [ -z $day ]
then
    echo "Please provide day as argument"
    exit 1
fi
# Check if day is a number
if ! [[ $day =~ ^[0-9]+$ ]]
then
    echo "Day must be a number"
    exit 1
fi
# Check if day is less than 1 or greater than 25
if [ $day -lt 1 ] || [ $day -gt 25 ]
then
    echo "Day must be between 1 and 25"
    exit 1
fi
# Assign directory name
if [ $day -lt 10 ]
then
    dir_name="day0$day"
else
    dir_name="day$day"
fi
# Check if day directory exists
if [ ! -d $dir_name ]
then
    mkdir $dir_name
    echo "Directory $dir_name created"
fi
cd $dir_name
# Check if day directory is empty
if ! ls $dir_name/* 1> /dev/null 2>&1
then
    # Check if properties file exists
    if [ ! -f "../../properties" ]
    then
        echo "Properties file not found"
        exit 1
    fi
    source ../../properties
    cp ../utils/template* "$dir_name.$EXTENSION"
    cp ../utils/test_template* "$dir_name"_test."$EXTENSION"
    sed -i "s#{{DAY}}#$dir_name#g" "$dir_name.$EXTENSION"
    echo "Copied template files to $dir_name"
fi
