#!/usr/bin/env bash

#  Eecute on pi after restart to make bluetooth work
sudo hciconfig hci0 down
sudo service bluetooth stop
