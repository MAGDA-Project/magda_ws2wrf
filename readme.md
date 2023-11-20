# MAGDA Weather stations to WRF


This module can be used to convert weather
stations observations in MAGDA format into ascii
WRF format.


## Installation

The module use go-netcdf to read a netcdf
containing world orography data.

In order to use it, you need the developer version of the
library provided by your distribution installed.

On ubuntu you can install it with:

```bash
sudo apt install libnetcdf-dev
```

The orography data is used to calculate the altitude of every weather station
given their latitude and longitude coordinates.

You can download the orography file from 
https://zenodo.org/record/4607436/files/orog.nc

The file must be saved in path `~/.magda_ws2wrf/orog.nc`

## Usage on CIMA Typhoon

An orography file is already usable by `wrfprod` user:
/data/safe/home/wrfprod/.magda_ws2wrf/orog.nc.

`magda_ws2wrf` is already present in /data/safe/home/wrfprod/bin/magda_ws2wrf

## Command line usage

This module implements a console command
that can be used to convert observations from
CSV to ascii WRF format.

Usage of `mag_ws2wrf`:

```
mag_ws2wrf [options] <input file>...
Options:
  -date string
        date and hour of the stations data to filter [YYYYMMDDHH]
  -domain string
        domain of the stations to filter [MinLat,MaxLat,MinLon,MaxLon]
  -outfile string
        where to save converted file (default "./out")
```

* <input file> can be specified more than once. When multiple files are specified, their data is 
concatenated in a single output ASCII file.
* if date option is specified, only measurement of the date will be converted
* if domain is option is specified, only stations contained in that domain will be converted
* if no date or domain option are used, all data in input files is converted.
