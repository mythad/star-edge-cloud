package com.starcloud.hbase;

import java.io.IOException;

import com.alibaba.fastjson.JSON;
import com.starcloud.model.Alarm;
import com.starcloud.model.Family;
import com.starcloud.model.RealtimeData;

import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.hbase.HBaseConfiguration;
import org.apache.hadoop.hbase.TableName;
import org.apache.hadoop.hbase.client.Admin;
import org.apache.hadoop.hbase.client.ColumnFamilyDescriptorBuilder;
import org.apache.hadoop.hbase.client.Connection;
import org.apache.hadoop.hbase.client.ConnectionFactory;
import org.apache.hadoop.hbase.client.Delete;
import org.apache.hadoop.hbase.client.Get;
import org.apache.hadoop.hbase.client.Put;
import org.apache.hadoop.hbase.client.Result;
import org.apache.hadoop.hbase.client.ResultScanner;
import org.apache.hadoop.hbase.client.Scan;
import org.apache.hadoop.hbase.client.Table;
import org.apache.hadoop.hbase.client.TableDescriptorBuilder;
import org.apache.hadoop.hbase.filter.FirstKeyOnlyFilter;
import org.apache.hadoop.hbase.util.Bytes;

public class FamilyManager {
    public static Connection connection;
    
    // 用HBaseconfiguration初始化配置信息是会自动加载当前应用的classpath下的hbase-site.xml
    public static Configuration configuration = HBaseConfiguration.create();

    // static Configuration conf = null;
    static {
        // configuration = HBaseConfiguration.create();
        configuration.set("hbase.zookeeper.quorum", "myhbase");
        configuration.set("hbase.zookeeper.property.clientPort", "2181");
         configuration.set("log4j.logger.org.apache.hadoop.hbase", "WARN");
        configuration.set("hbase.zookeeper.property.dataDir", "/hbase-data");
        // configuration.set("hbase.rootdir", "/apps/root");
        configuration.set("hadoop.home.dir", "/hbase");
    }

    public static void initConnection() throws Exception {
        // 对connection进行初始化、
        // 当然也可以手动加载配置文件，手动加载配置文件时要调用configuration的addResource方法
        // configuration.addResource("hbase-site.xml");
        connection = ConnectionFactory.createConnection(configuration);
    }

    public static void createFamilyTable() throws IOException {
        Admin admin = connection.getAdmin();
        try {
            Admin hAdmin = connection.getAdmin();  
            if (hAdmin.tableExists(TableName.valueOf("monitor"))) {  
                System.out.println("表已存在");  
                return;
            }
            TableDescriptorBuilder builder = TableDescriptorBuilder.newBuilder(TableName.valueOf("monitor"));
            builder.setColumnFamily(ColumnFamilyDescriptorBuilder.newBuilder(Bytes.toBytes("realtime")).build())
                    .setColumnFamily(ColumnFamilyDescriptorBuilder.newBuilder(Bytes.toBytes("result")).build())
                    .setColumnFamily(ColumnFamilyDescriptorBuilder.newBuilder(Bytes.toBytes("alarm")).build());
            admin.createTable(builder.build());
        } finally {
            admin.close();
        }
    }

    public static void putRealtimeData(RealtimeData data) throws IOException {
        Put put = new Put(Bytes.toBytes(data.getId()));
        String jsonString = JSON.toJSONString(data);
        // 将数据添加到put中
        put.addColumn(Bytes.toBytes("realtime"), Bytes.toBytes("data"), Bytes.toBytes(jsonString));
        put.addColumn(Bytes.toBytes("realtime"), Bytes.toBytes("id"), Bytes.toBytes(data.getId()));
        Table table = connection.getTable(TableName.valueOf("monitor"));
        // 将put写入HBase
        table.put(put);
    }

    public static void putResult(com.starcloud.model.Result data) throws IOException {
        Put put = new Put(Bytes.toBytes(data.getRealtimeDataId()));
        String jsonString = JSON.toJSONString(data);
        // 将数据添加到put中
        put.addColumn(Bytes.toBytes("result"), Bytes.toBytes("data"), Bytes.toBytes(jsonString));
        put.addColumn(Bytes.toBytes("result"), Bytes.toBytes("id"), Bytes.toBytes(data.getId()));
        Table table = connection.getTable(TableName.valueOf("monitor"));
        // 将put写入HBase
        table.put(put);
    }

    public static void putAlarm(Alarm data) throws IOException {
        Put put = new Put(Bytes.toBytes(data.getRealtimeDataId()));
        String jsonString = JSON.toJSONString(data);
        // 将数据添加到put中
        put.addColumn(Bytes.toBytes("alarm"), Bytes.toBytes("data"), Bytes.toBytes(jsonString));
        put.addColumn(Bytes.toBytes("alarm"), Bytes.toBytes("id"), Bytes.toBytes(data.getId()));
        Table table = connection.getTable(TableName.valueOf("monitor"));
        // 将put写入HBase
        table.put(put);
    }

    public static Family getFamily(String rowkey) throws IOException {
        TableName tbname = TableName.valueOf("monitor");
        Table table = connection.getTable(tbname);
        Get get = new Get(Bytes.toBytes(rowkey));
        Result result = table.get(get);

        Family family = new Family();
        family.setId(rowkey);

        byte[] d1 = result.getValue(Bytes.toBytes("realtime"), Bytes.toBytes("data"));
        String s1 = Bytes.toString(d1);
        family.setRealtimeData(JSON.parseObject(s1, RealtimeData.class));

        byte[] d2 = result.getValue(Bytes.toBytes("result"), Bytes.toBytes("data"));
        String s2 = Bytes.toString(d1);
        family.setRealtimeData(JSON.parseObject(s2, RealtimeData.class));

        byte[] d3 = result.getValue(Bytes.toBytes("alarm"), Bytes.toBytes("data"));
        String s3 = Bytes.toString(d1);
        family.setRealtimeData(JSON.parseObject(s3, RealtimeData.class));

        return family;
    }

    // 删除一行数据
    public static void deleteFamily(String familyId) throws Exception {
        Table table = connection.getTable(TableName.valueOf("monitor"));
        // 通过行键删除一整行的数据
        Delete deletRow = new Delete(Bytes.toBytes(familyId));
        table.delete(deletRow);
    }

    /**
     * 使用Scan与Filter的方式对表行数进行统计
     * 
     * @throws IOException
     */
    public static long RowCountWithScanAndFilter() throws IOException {
        Table table = connection.getTable(TableName.valueOf("monitor"));
        long rowCount = 0;
        try {
            Scan scan = new Scan();
            scan.setFilter(new FirstKeyOnlyFilter());
            ResultScanner resultScanner = table.getScanner(scan);
            for (Result result : resultScanner) {
                rowCount += result.size();
            }
        } catch (IOException e) {
            throw new RuntimeException();
        }
        return rowCount;
    }
}