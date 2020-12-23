#!/bin/bash
[ -d "~/.config/fileOrganizer" ] && mkdir ~/.config/fileOrganizer

ln -f ./config.json ~/.config/fileOrganizer/config.json

sudo ln -f ./fileOrganizer.service ~/.config/systemd/user/fileOrganizer.service

sudo cp -f ./main /bin/fileOrganizer
