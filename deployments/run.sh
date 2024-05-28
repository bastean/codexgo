#!/bin/bash

logs=logs

today=$(date -u +%d-%m-%Y)

now=$(date -u +%H_%M_%S)

log=$logs/$today/$now.log

./codexgo |& tee $log
