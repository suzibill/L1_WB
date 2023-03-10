#!/bin/bash

for i in {01..26}; do
  mkdir "ex$i"
  touch ex$i/main.go
done
