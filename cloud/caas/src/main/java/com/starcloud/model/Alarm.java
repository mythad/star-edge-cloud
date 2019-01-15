package com.starcloud.model;

public class Alarm {
    private String Id;
    private String realtimeDataId;
    private int level;

    /**
     * @return the level
     */
    public int getLevel() {
        return level;
    }

    /**
     * @param level the level to set
     */
    public void setLevel(int level) {
        this.level = level;
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