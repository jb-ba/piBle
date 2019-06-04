#!/usr/bin/env bash

# Only adds things which are not in .gitignore (unlike "git add .")
git add *
git pull
git commit -m "automatic push from shell"
git push origin master

