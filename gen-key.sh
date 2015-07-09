#!/bin/bash
cat >signer-key.txt <<EOF
     %echo Generating a basic OpenPGP key
     %no-ask-passphrase
     Key-Type: dsa
     Key-Length: 2048
     Key-Usage: sign
     Name-Real: signature server
     Name-Comment: keys used for sign files
     Name-Email: tarball-signer@example.org
     Expire-Date: 0
     %pubring pubring.gpg
     %secring secring.gpg
     %commit
     %echo done
EOF
gpg2 --batch --gen-key signer-key.txt
rm signer-key.txt
