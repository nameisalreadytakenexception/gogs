#!/bin/sh
sudo service gogs start
sudo chown -R git /vat/app/current/custom/conf/app.ini
