package com.starcloud.caas;

import com.starcloud.docker.DockerHelper;
import com.starcloud.hbase.FamilyManager;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class DemoApplication {

	// docker run -p 8888:8080 -v
	// ./star-edge-cloud/display/web:/usr/local/tomcat/webapps/web tomcat
	public static void main(String[] args) {
		try {
			System.setProperty("hadoop.home.dir", "/");
			DockerHelper.Initialize();
			FamilyManager.initConnection();
			FamilyManager.createFamilyTable();
			SpringApplication.run(DemoApplication.class, args);
		} catch (Exception ex) {

		}
	}
}
