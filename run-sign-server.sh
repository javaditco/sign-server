#!/bin/bash

cd /sign-server

if [ ! -e sign-server ]; then
  make deps
  make
fi

./sign-server
