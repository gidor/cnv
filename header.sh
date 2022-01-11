#!/bin/bash

SOURCE_DIR=`dirname $0`

for i in *.go # or whatever other pattern...
do
  if ! grep -q Copyright $i
  then
    cat $i | awk -f $SOURCE_DIR/remove_header.awk > $i.new && cat $SOURCE_DIR/copyright.txt $i.new > $i 
  fi
done

stripheader(){

mv $SOURCE_DIR/tmp $1

}