package io.stackdocker.curriculum.model;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;

public class TopPractice {
    private long id;

    @JsonProperty
    private  String kind;

    @JsonProperty
    private  String name;

    private  String topic;

    private  String way;

    private String tag;


    @JsonCreator
    public TopPractice(
            @JsonProperty("id")  long id,
            @JsonProperty("name") String name  ) {

        this.id = id;
        this.name = name;

        this.topic = this.name;
        this.kind = "HotPractice";
        this.way = "cook";
        this.tag = "java, web, micro-service";
    }

    public TopPractice(long id, String name, String tag) {
        this(id, name);
        this.tag = tag;
    }

    public TopPractice(long id, String name, String way, String tag) {
        this(id, name, tag);
        this.way = way;
    }

    public String getWay() {
        return this.way;
    }

    public void setWay(String way) {
        this.way = way;
    }

    public String getTag() {
        return this.tag;
    }

    public void setTag(String tag) {
        this.tag = tag;
    }

}
