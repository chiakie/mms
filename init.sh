#!/usr/bin/env bash

if [ -d "./vendor" ]; then
    rm -rf ./vendor
fi

govendor init

# gin related
govendor fetch github.com/gin-gonic/gin
govendor fetch github.com/gin-contrib/sessions

# database related
govendor fetch github.com/mattn/go-sqlite3
govendor fetch github.com/jinzhu/gorm/^

# config related
govendor fetch gopkg.in/yaml.v2

# form validator related
govendor fetch github.com/asaskevich/govalidator
