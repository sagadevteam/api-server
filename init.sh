#!/bin/bash
go get -u github.com/pressly/goose/cmd/goose
echo "Successfully installed goose"
echo "Use \"goose -dir migrations/ mysql 'user:password@/database?parseTime=true' up\" to complete migrations."