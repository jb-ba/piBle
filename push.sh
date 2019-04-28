#!/usr/bin/env bash
git add *
git commit -m "automatic push from shell"
printf "jonas27\n" | git push origin master

