#!/usr/bin/env bash

for i in `ls`;do mv $i ${echo $i | awk -F '\@' '{print $1}'};done
