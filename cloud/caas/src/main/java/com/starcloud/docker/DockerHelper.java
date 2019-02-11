package com.starcloud.docker;

import java.util.List;

import com.github.dockerjava.api.DockerClient;
import com.github.dockerjava.api.command.CreateContainerResponse;
import com.github.dockerjava.api.command.InspectVolumeResponse;
import com.github.dockerjava.api.command.ListImagesCmd;
import com.github.dockerjava.api.command.ListVolumesResponse;
import com.github.dockerjava.api.model.Container;
import com.github.dockerjava.api.model.ExposedPort;
import com.github.dockerjava.api.model.Image;
import com.github.dockerjava.api.model.Network;
import com.github.dockerjava.api.model.Ports;
import com.github.dockerjava.core.DockerClientBuilder;

//https://github.com/docker-java/docker-java/wiki
public class DockerHelper {
    static DockerClient dockerClient;
    private static String nodeAddress = "tcp://localhost:2375";

    public static void Initialize() {
        // 使用DockerClientBuilder创建链接
        dockerClient = DockerClientBuilder.getInstance(nodeAddress).build();
    }

    public void CreateContainer(String config) {
        ExposedPort tcp8080 = ExposedPort.tcp(8080);
        // 设置映射到主机的端口
        Ports portBindings = new Ports();
        portBindings.bind(tcp8080, Ports.Binding.bindPort(8089));
        // 创建一个新的Container并且与主机端口号绑定
        CreateContainerResponse container = dockerClient.createContainerCmd("hadoop:latest")
                .withPortBindings(portBindings).exec();
        // 运行一个Container
        dockerClient.startContainerCmd(container.getId()).exec();
    }

    public void StartContainer(String cid) {
        dockerClient.startContainerCmd(cid).exec();
    }

    public void StopContainer(String cid) {
        dockerClient.stopContainerCmd(cid).exec();
    }

    public void removeContainer(String cid) {
        dockerClient.removeContainerCmd(cid).exec();
    }

    public List<Image> GetAllImages() {
        try {
            ListImagesCmd cmd = dockerClient.listImagesCmd();
            List<Image> images = cmd.exec();
            return images;
        } catch (Exception ex) {
            throw ex;
        }
    }

    public List<Container> GetAllContainer() {
        try {
            List<Container> containers = dockerClient.listContainersCmd().exec();
            return containers;
        } catch (Exception ex) {
            return null;
        }
        // dockerClient.listContainersCmd()
        // dockerClient.inspectContainerCmd(containerId)
    }

    public void CreateNetwork() {
        dockerClient.createNetworkCmd().withName("staredgecloud").withDriver("overlay").exec();
    }

    public void RemoveNetwork(String nid) {
        dockerClient.removeNetworkCmd(nid).exec();
    }

    public List<Network> GetNetwork() {
        try {
            List<Network> networks = dockerClient.listNetworksCmd().exec();
            return networks;
        } catch (Exception ex) {
            return null;
        }
    }

    public void CreateVolumn() {
        dockerClient.createVolumeCmd().withName("myNamedVolume").exec();
    }

    public void RemoveVolumn(String vid) {
        dockerClient.removeVolumeCmd(vid).exec();
    }

    public List<InspectVolumeResponse> GetVolume() {
        try {
            ListVolumesResponse vg = dockerClient.listVolumesCmd().exec();
            return vg.getVolumes();
        } catch (Exception ex) {
            return null;
        }
    }
}