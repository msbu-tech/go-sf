#!/bin/sh
APP_NAME="{{.AppName}}"

rm -rf output
mkdir -p output/app/$APP_NAME

cp -r conf output/
cp -r service library output/app/$APP_NAME

cd output
find . -type d -name ".svn"|xargs rm -rf
tar cvzf $APP_NAME.tar.gz app conf
rm -rf app conf
