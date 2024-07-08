#!/bin/bash

logs=logs

today=$(date -u +%d-%m-%Y)

mkdir -p $logs/$today

now=$(date -u +%H_%M_%S)

log=$logs/$today/$now.log

./codexgo |& tee $log
