#!/bin/sh

rsync --archive --compress --delete ./bin/ipcr root@111.111.111.111:/var/www/api.example.com/bin/