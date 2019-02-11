package com.starcloud.model;

public class RealtimeData {
    private String Id;
    private String Type;
    private byte[] Data;

    /**
     * @return the type
     */
    public String getType() {
        return Type;
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
     * @return the data
     */
    public byte[] getData() {
        return Data;
    }

    /**
     * @param data the data to set
     */
    public void setData(byte[] data) {
        this.Data = data;
    }

    /**
     * @param type the type to set
     */
    public void setType(String type) {
        this.Type = type;
    }

}