#!/bin/sh
PRODUCT_NAME="{{.Productname}}"
APP_NAME="{{.Appname}}"

rm -rf output
mkdir -p output/app/$APP_NAME
mkdir -p output/conf/app/$APP_NAME

cd output
find . -type d -name ".svn"|xargs rm -rf
tar cvzf $APP_NAME.tar.gz app conf
rm -rf app conf
