#!/bin/bash

cd /sign-server

if [ ! -e sign-server ]; then
  make deps
  make
fi

if [ ! -e  secring.pgp ]; then
  ./gen-key.sh
fi

./sign-server
