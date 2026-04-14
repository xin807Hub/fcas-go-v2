### V1是第1版基于Java开发的大数据项目
### V2是第2版基于Golang开发的大数据项目


#### “fcas_cluster”是本项目 clickhouse 集群的名称, 每新增一台clickhouse服务器，都需要在配置文件：/etc/clickhouse-server/config.xml 中配置好集群。
配置方式如下
<remote_servers>
    <fcas_cluster>
      <shard>
        <replica>
          <host>127.0.0.1</host>
          <port>9000</port>
          <user>default</user>
          <password>123456</password>
        </replica>
      </shard>
    </fcas_cluster>
</remote_servers>

在remote_servers标签中新增一个集群，集群名称就是标签名。
每一个shard代表1台机器，即1个分片。


造数据：
clickhouse-client -udefault --password RootSi314 --query "INSERT INTO bigdata_fcas.ods_gen_traffic (o_start_time, user_id, d_user_id, isp, d_isp, src_ip, dst_ip, src_port, dst_port, protocol, app_type, app_id, bytes_up, bytes_dn, src_area_id, dst_area_id, link_id, is_oversea, dst_province, host)
SELECT now() , user_id, d_user_id, isp, d_isp, src_ip, dst_ip, src_port, dst_port, protocol, app_type, app_id, bytes_up, bytes_dn, src_area_id, dst_area_id, link_id, is_oversea, dst_province, host FROM bigdata_fcas.ods_gen_traffic"



