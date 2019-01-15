package com.starcloud.model;

public class Result {
    private String Id;
    private byte[] Data;
    private String realtimeDataId;

     /**
     * @return the data
     */
    public byte[] getData() {
        return Data;
    }

    /**
     * @return the id
     */
    public String getId() {
        return Id;
    }

    /**
     * @param id the id to set
     */
    public void setId(String id) {
        this.Id = id;
    }

    /**
     * @param data the data to set
     */
    public void setData(byte[] data) {
        this.Data = data;
    }

    /**
     * @return the realtimeDataId
     */
    public String getRealtimeDataId() {
        return realtimeDataId;
    }

    /**
     * @param realtimeDataId the realtimeDataId to set
     */
    public void setRealtimeDataId(String realtimeDataId) {
        this.realtimeDataId = realtimeDataId;
    }
}