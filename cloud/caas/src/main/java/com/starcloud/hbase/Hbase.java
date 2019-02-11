package com.starcloud.hbase;

import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.hbase.*;
import org.apache.hadoop.hbase.client.*;
import org.apache.hadoop.hbase.filter.*;
import org.apache.hadoop.hbase.util.Bytes;
import java.io.IOException;

/**
 * @author allen
 */
public class Hbase {
    static Configuration conf = null;
    static {
        conf = HBaseConfiguration.create();
        conf.set("hbase.zookeeper.quorum", "myhbase");
        conf.set("hbase.zookeeper.property.clientPort", "2181");
        conf.set("log4j.logger.org.apache.hadoop.hbase", "WARN");
    }

    public static void createTable(String tableName, String... families) throws Exception {
        HTableDescriptor tableDescriptor = new HTableDescriptor(TableName.valueOf(tableName));
        try {
            Connection connection = ConnectionFactory.createConnection(conf);
            Admin admin = connection.getAdmin();
            for (String family : families) {
                tableDescriptor.addFamily(new HColumnDescriptor(family));
            }
            if (admin.tableExists(TableName.valueOf(tableName))) {
                System.out.println("Table Exists");
                System.exit(0);
            } else {
                admin.createTable(tableDescriptor);
                System.out.println("Create table Success!!!Table Name:[" + tableName + "]");
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public static void deleteTable(String tableName) {
        try {
            Connection connection = ConnectionFactory.createConnection(conf);
            Admin admin = connection.getAdmin();
            TableName table = TableName.valueOf(tableName);
            admin.disableTable(table);
            admin.deleteTable(table);
            System.out.println("delete table " + tableName + " ok!");
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public static void updateTable(String tableName, String rowKey, String familyName, String columnName, String value)
            throws Exception {
        try {
            Connection connection = ConnectionFactory.createConnection(conf);
            Table table = connection.getTable(TableName.valueOf(tableName));
            Put put = new Put(Bytes.toBytes(rowKey));
            put.addColumn(Bytes.toBytes(familyName), Bytes.toBytes(columnName), Bytes.toBytes(value));
            table.put(put);
            System.out.println("Update table success");
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public static void addData(String rowKey, String tableName, String[] column, String[] value) {
        try {
            Connection connection = ConnectionFactory.createConnection(conf);
            Table table = connection.getTable(TableName.valueOf(tableName));
            Put put = new Put(Bytes.toBytes(rowKey));
            HColumnDescriptor[] columnFamilies = table.getTableDescriptor().getColumnFamilies();
            for (int i = 0; i < columnFamilies.length; i++) {
                String familyName = columnFamilies[i].getNameAsString();
                if (familyName.equals("version")) {
                    for (int j = 0; j < column.length; j++) {
                        put.addColumn(Bytes.toBytes(familyName), Bytes.toBytes(column[j]), Bytes.toBytes(value[j]));
                    }
                    table.put(put);
                    System.out.println("Add Data Success!");
                }
            }
        } catch (IOException e) {
        }
    }

    public static void deleteAllColumn(String tableName, String rowKey) {
        try {
            Connection connection = ConnectionFactory.createConnection(conf);
            Table table = connection.getTable(TableName.valueOf(tableName));
            Delete delAllColumn = new Delete(Bytes.toBytes(rowKey));
            table.delete(delAllColumn);
            System.out.println("Delete AllColumn Success");
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public static void deleteColumn(String tableName, String rowKey, String familyName, String columnName) {
        try {
            Connection connection = ConnectionFactory.createConnection(conf);
            Table table = connection.getTable(TableName.valueOf(tableName));
            Delete delColumn = new Delete(Bytes.toBytes(rowKey));
            delColumn.addColumn(Bytes.toBytes(familyName), Bytes.toBytes(columnName));
            table.delete(delColumn);
            System.out.println("Delete Column Success");
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public static void getResultByVersion(String tableName, String rowKey, String familyName, String columnName) {
        try {
            Connection connection = ConnectionFactory.createConnection(conf);
            Table table = connection.getTable(TableName.valueOf(tableName));
            Get get = new Get(Bytes.toBytes(rowKey));
            get.addColumn(Bytes.toBytes(familyName), Bytes.toBytes(columnName));
            get.setMaxVersions(3);
            Result result = table.get(get);
            for (Cell cell : result.listCells()) {
                System.out.println("family:"
                        + Bytes.toString(cell.getFamilyArray(), cell.getFamilyOffset(), cell.getFamilyLength()));
                System.out.println("qualifier:" + Bytes.toString(cell.getQualifierArray(), cell.getQualifierOffset(),
                        cell.getQualifierLength()));
                System.out.println(
                        "value:" + Bytes.toString(cell.getValueArray(), cell.getValueOffset(), cell.getValueLength()));
                System.out.println("Timestamp:" + cell.getTimestamp());
                System.out.println("---------------");
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    // public static void main(String[] args) throws Exception {
    //     System.setProperty("hadoop.home.dir", "/usr/local/hadoop");
    //     Hbase.createTable("ota_pre_record", "version"); // Hbase.deleteTable("ota_pre_record"); //
    //                                                     // Hbase.updateTable("ota_pre_record","324b1f27c982ea87XyZf+1502176018+20900","version","check_time","2017-11-08
    //                                                     // 10:46:34"); //
    //                                                     // Hbase.updateTable("ota_pre_record","324b1f27c982ea87XyZf+1502176018+20900","version","download_time","2017-11-08
    //                                                     // 14:37:34"); //
    //                                                     // Hbase.updateTable("ota_pre_record","24b1f27c982ea87XyZf+1502176018+20900","version","upgrade_time","2017-11-08
    //                                                     // 15:35:34"); //
    //                                                     // Hbase.deleteColumn("ota_pre_record","24b1f27c982ea87XyZf+1502176018+20900","version","upgrade_time");
    //                                                     // //
    //                                                     // Hbase.deleteAllColumn("ota_pre_record","24b1f27c982ea87XyZf+1502176018+20900");
    //                                                     // //
    //                                                     // Hbase.getResultByVersion("ota_pre_record","24b1f27c982ea87XyZf+1502176018+20900","version","check_time");
    //                                                     // } }
    // }
}