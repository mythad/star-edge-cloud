package com.starcloud.model;

public class Family {
    private String Id;
    private RealtimeData realtimeData;
    private Alarm alarm;
    private Result result;

    /**
     * @return the id
     */
    public String getId() {
        return Id;
    }

    /**
     * @return the result
     */
    public Result getResult() {
        return result;
    }

    /**
     * @param result the result to set
     */
    public void setResult(Result result) {
        this.result = result;
    }

    /**
     * @return the alarm
     */
    public Alarm getAlarm() {
        return alarm;
    }

    /**
     * @param alarm the alarm to set
     */
    public void setAlarm(Alarm alarm) {
        this.alarm = alarm;
    }

    /**
     * @return the realtimeData
     */
    public RealtimeData getRealtimeData() {
        return realtimeData;
    }

    /**
     * @param realtimeData the realtimeData to set
     */
    public void setRealtimeData(RealtimeData realtimeData) {
        this.realtimeData = realtimeData;
    }

    /**
     * @param id the id to set
     */
    public void setId(String id) {
        this.Id = id;
    }

}