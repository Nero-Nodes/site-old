#!/bin/bash

cd /srv/site
git pull
go build -o /usr/local/bin/site
systemctl restart site