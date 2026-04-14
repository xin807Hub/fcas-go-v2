#!/bin/bash

install -D monitor_disk.sh /home/monitor_disk/
chmod +x /home/monitor_disk/monitor_disk.sh
\cp -rf monitor_disk.service /usr/lib/systemd/system/

systemctl daemon-reload
systemctl enable monitor_disk
systemctl restart monitor_disk

