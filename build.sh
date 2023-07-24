#!/bin/bash

echo "building project to ./out/vroom.exe"

rm ./out/vroom.exe

go build -o ./out/
