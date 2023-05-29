#!/bin/sh

option="${1}"
echo ${option}
case ${option} in 
   -u)
      docker-compose -f ${PWD}/docker-compose.yml up -d
      ;; 
   -d)
      docker-compose -f ${PWD}/docker-compose.yml down
      ;; 
   *)  
      echo "`basename ${0}`:usage: [-u up] | [-d down]" 
      exit 1 # Command to come out of the program with status 1
      ;; 
esac 