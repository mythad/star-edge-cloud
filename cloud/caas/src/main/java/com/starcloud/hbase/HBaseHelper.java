// package com.starcloud.hbase;

// import java.io.IOException;
// import java.util.ArrayList;
// import java.util.List;

// import com.starcloud.model.Family;

// import org.apache.hadoop.conf.Configuration;
// import org.apache.hadoop.hbase.HBaseConfiguration;
// import org.apache.hadoop.hbase.TableName;
// import org.apache.hadoop.hbase.client.Admin;
// import org.apache.hadoop.hbase.client.ColumnFamilyDescriptorBuilder;
// import org.apache.hadoop.hbase.client.Connection;
// import org.apache.hadoop.hbase.client.ConnectionFactory;
// import org.apache.hadoop.hbase.client.Delete;
// import org.apache.hadoop.hbase.client.Get;
// import org.apache.hadoop.hbase.client.Put;
// import org.apache.hadoop.hbase.client.Result;
// import org.apache.hadoop.hbase.client.ResultScanner;
// import org.apache.hadoop.hbase.client.Scan;
// import org.apache.hadoop.hbase.client.Table;
// import org.apache.hadoop.hbase.client.TableDescriptorBuilder;
// import org.apache.hadoop.hbase.filter.FirstKeyOnlyFilter;
// import org.apache.hadoop.hbase.util.Bytes;
// import org.apache.tomcat.jni.Time;

// //https://blog.csdn.net/sinat_39409672/article/details/78403015
// public class HBaseHelper {
//     public static Connection connection;
//     // 用HBaseconfiguration初始化配置信息是会自动加载当前应用的classpath下的hbase-site.xml
//     public static Configuration configuration = HBaseConfiguration.create();
//     // static Configuration conf = null;
//     static {
//         // configuration = HBaseConfiguration.create();
//         configuration.set("hbase.zookeeper.quorum", "myhbase");
//         configuration.set("hbase.zookeeper.property.clientPort", "2181");
//         configuration.set("log4j.logger.org.apache.hadoop.hbase", "WARN");
//     }

//     public static void InitConnection() throws Exception {
//         // 对connection进行初始化、
//         // 当然也可以手动加载配置文件，手动加载配置文件时要调用configuration的addResource方法
//         // configuration.addResource("hbase-site.xml");
//         connection = ConnectionFactory.createConnection(configuration);
//     }

//     public static byte[] getData(String rowkey) throws Exception {
//         TableName tbname = TableName.valueOf("bd14:fromJava");
//         Table table = connection.getTable(tbname);
//         Get get = new Get(Bytes.toBytes(rowkey));
//         Result result = table.get(get);
//         return result.getValue(Bytes.toBytes("i"), Bytes.toBytes("realtime_data"));
//         // 遍历结果对象results
//         // for (Result result : results) {
//         // // 嵌套遍历result获取cell
//         // for (Cell cell : result.listCells()) {
//         // // 使用CellUtil工具类直接获取cell中的数据
//         // String family = Bytes.toString(CellUtil.cloneFamily(cell));
//         // String qualify = Bytes.toString(CellUtil.cloneQualifier(cell));
//         // String rowkey = Bytes.toString(CellUtil.cloneRow(cell));
//         // String value = Bytes.toString(CellUtil.cloneValue(cell));
//         // System.err.println(family + "_" + qualify + "_" + rowkey + "_" + value);
//         // }
//         // }
//     }

//     public static void putData(String tableName, String rowkey, String family, String column) throws Exception {
//         // 通过表名获取tbName
//         TableName tbname = TableName.valueOf(tableName);
//         // 通过connection获取相应的表
//         Table table = connection.getTable(tbname);
//         // 创建Random对象以作为随机参数
//         // Random random = new Random();
//         // hbase支持批量写入数据，创建Put集合来存放批量的数据
//         List<Put> batput = new ArrayList<>();
//         // 实例化put对象，传入行键
//         Put put = new Put(Bytes.toBytes(family.getId()));
//         // 调用addcolum方法，向i簇中添加字段
//         put.addColumn(Bytes.toBytes(family), Bytes.toBytes(column), data);
//         table.put(batput);
//     }

//     public void updateData(String tableName, String rowKey, String family, String columkey, String updatedata)
//             throws Exception {
//         // hbase中更新数据同样采用put方法，在相同的位置put数据，则在查询时只会返回时间戳较新的数据
//         // 且在文件合并时会将时间戳较旧的数据舍弃
//         Put put = new Put(Bytes.toBytes(rowKey));
//         // 将新数据添加到put中
//         put.addColumn(Bytes.toBytes(family), Bytes.toBytes(columkey), Bytes.toBytes(updatedata));
//         Table table = connection.getTable(TableName.valueOf(tableName));
//         // 将put写入HBase
//         table.put(put);
//     }

//     public static void createTable(String tableName, String... cf1) throws Exception {
//         Admin admin = connection.getAdmin();
//         try {
//             TableDescriptorBuilder builder = TableDescriptorBuilder.newBuilder(TableName.valueOf(tableName));
//             for (String colName : cf1) {
//                 builder.setColumnFamily(ColumnFamilyDescriptorBuilder.newBuilder(Bytes.toBytes(colName)).build());
//             }
//             admin.createTable(builder.build());
//         } finally {
//             admin.close();
//         }
//     }

//     public void deleteTable(String tableName) throws Exception {
//         Admin admin = connection.getAdmin();
//         // 通过tableName创建表名
//         TableName tbName = TableName.valueOf(tableName);
//         // 判断表是否存在，若存在就删除，不存在就退出
//         if (admin.tableExists(tbName)) {
//             // 首先将表解除占用，否则无法删除
//             admin.disableTable(tbName);
//             // 调用delete方法
//             admin.deleteTable(tbName);
//             System.err.println("表" + tableName + "已删除");
//         } else {
//             System.err.println("表" + tableName + "不存在！");
//         }
//     }

//     // 删除某条记录
//     public void deleteData(String tableName, String rowKey, String family, String columkey) throws Exception {
//         Table table = connection.getTable(TableName.valueOf(tableName));
//         // 创建delete对象
//         Delete deletData = new Delete(Bytes.toBytes(rowKey));
//         // 将要删除的数据的准确坐标添加到对象中
//         deletData.addColumn(Bytes.toBytes(family), Bytes.toBytes(columkey));
//         // 删除表中数据 
//         table.delete(deletData);
//     }

//     // 删除一行数据
//     public void deleteRow(String tableName, String rowKey) throws Exception {
//         Table table = connection.getTable(TableName.valueOf(tableName));
//         // 通过行键删除一整行的数据
//         Delete deletRow = new Delete(Bytes.toBytes(rowKey));
//         table.delete(deletRow);
//     }

//     /**
//      * 使用Scan与Filter的方式对表行数进行统计
//      */
//     public static long RowCountWithScanAndFilter(Table table) {
//         long rowCount = 0;
//         try {
//             Scan scan = new Scan();
//             scan.setFilter(new FirstKeyOnlyFilter());
//             ResultScanner resultScanner = table.getScanner(scan);
//             for (Result result : resultScanner) {
//                 rowCount += result.size();
//             }
//         } catch (IOException e) {
//             throw new RuntimeException();
//         }
//         return rowCount;
//     }
// }
