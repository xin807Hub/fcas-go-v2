#!/bin/bash

workdir=$(pwd)

\cp -f $workdir/config.yaml /home/etl/
\cp -rf $workdir/dict /home/etl/
\cp -rf $workdir/etl.service /usr/lib/systemd/system/
# \cp -rf $workdir/libzmq.so.5.0.0 /usr/lib64/
# ln -s /usr/lib64/libzmq.so.5.0.0 /usr/lib64/libzmq.so.5

systemctl daemon-reload
systemctl enable etl
systemctl restart etl
systemctl status etl
