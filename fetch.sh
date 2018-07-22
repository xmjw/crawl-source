#! /bin/bash

# Fetch the relevant CSV and store it...
curl "http://downloads.majestic.com/majestic_million.csv" -o $(date +%Y-%m-%d).csv
