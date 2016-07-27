package com.protocol10.taskapp.creating;

public class TrackModel {
    public String id;
    public String trackName;
    public String trackPath;

    public TrackModel(String id, String trackName, String trackPath) {
        this.id = id;
        this.trackName = trackName;
        this.trackPath = trackPath;
    }

    @Override
    public String toString() {
        return "TrackId\t=" + id + "/n Track name\t=" + trackName;
    }
}
