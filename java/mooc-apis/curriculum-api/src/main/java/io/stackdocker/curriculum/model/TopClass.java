package io.stackdocker.curriculum.model;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonProperty;

public class TopClass {
    private long id;

    @JsonProperty
    private  String kind;

    @JsonProperty
    private  String name;

    @JsonIgnore
    private  String topic;

    private  String tag;

    @JsonCreator
    public TopClass(
            @JsonProperty("id")  long id,
            @JsonProperty("name") String name  ) {

        this.id = id;
        this.name = name;

        this.topic = this.name;
        this.kind = "HotClass";
        this.tag = "top1, lesson";
    }

    public TopClass(long id, String name, String tag) {
        this(id, name);
        this.tag = tag;
    }

    public String getTag() {
        return this.tag;
    }

    public void setTag(String tag) {
        this.tag = tag;
    }
}
