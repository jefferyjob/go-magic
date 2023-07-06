#!/bin/bash

cd ../ && go mod tidy && go build -o app cmd/app