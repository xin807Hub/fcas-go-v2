#! /bin/bash

mysql -uroot -p123456 < /root/fcas_service.sql
mysql -uroot -p123456 < /root/fcas_system.sql
mysql -uroot -p123456 < /root/biz_isp.sql
mysql -uroot -p123456 < /root/biz_ip_address.sql

