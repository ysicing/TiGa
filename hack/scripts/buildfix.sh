#!/bin/bash

go clean -modcache
go mod download
task local
