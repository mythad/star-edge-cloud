package com.starcloud.caas;

import java.util.List;

import com.github.dockerjava.api.model.Container;
import com.github.dockerjava.api.model.Image;
import com.starcloud.docker.DockerHelper;

import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
@EnableAutoConfiguration
public class DockerController {
    @CrossOrigin(origins = "*", maxAge = 3600)
    @RequestMapping(value = "/api/caas/listimage")
    public Object listImage() {
        DockerHelper helper = new DockerHelper();
        List<Image> images = helper.GetAllImages();
        return images;
    }

    @CrossOrigin(origins = "*", maxAge = 3600)
    @RequestMapping(value = "/api/caas/listcontainer")
    @ResponseBody
    public List<Container> listContainer() {
        DockerHelper helper = new DockerHelper();
        return helper.GetAllContainer();
    }

    @RequestMapping(value = "/api/caas/createcontainer")
    @ResponseBody
    public String createContainer(String cid) {
        try {
            DockerHelper helper = new DockerHelper();
            helper.CreateContainer(cid);
            return "";
        } catch (Exception ex) {
            return "error";
        }
    }

    @RequestMapping(value = "/api/caas/startcontainer")
    @ResponseBody
    public String startContainer(String cid) {
        try {
            DockerHelper helper = new DockerHelper();
            helper.StartContainer("cid");
            return "";
        } catch (Exception ex) {
            return "error";
        }
    }

    @RequestMapping(value = "/api/caas/stopcontainer")
    @ResponseBody
    public String stopContainer(String cid) {
        try {
            DockerHelper helper = new DockerHelper();
            helper.StopContainer(cid);
            return "";
        } catch (Exception ex) {
            return "error";
        }
    }

    @RequestMapping(value = "/api/caas/removecontainer")
    @ResponseBody
    public String removeContainer(String cid) {
        try {
            DockerHelper helper = new DockerHelper();
            helper.removeContainer(cid);
            return "success";
        } catch (Exception ex) {
            return "error";
        }
    }

    public String addNetwork() {
        return "";
    }
}