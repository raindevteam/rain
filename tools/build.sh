#!/usr/bin/env bash

cd internal/handler && go generate
cd ../..
go install
