#! /bin/bash

seed() {
    go run main.go seed
}

if [[ $1 == "seed" ]]
then
    seed
fi