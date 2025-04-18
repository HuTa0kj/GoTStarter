#!/bin/bash

# input package name
read -p "Origin Module Name（Such as gotstarter）: " OLD_PREFIX
read -p "New Module Name（Such as myproject）: " NEW_MODULE

echo "Update go.mod ->: $NEW_MODULE"
go mod edit -module "$NEW_MODULE"

echo "Replace: ${OLD_PREFIX}/ → ${NEW_MODULE}/"
find . -name "*.go" -exec sed -i '' "s|${OLD_PREFIX}/|${NEW_MODULE}/|g" {} +

echo "go mod tidy..."
go mod tidy

echo "Init Successful！"