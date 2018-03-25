package io.stackdocker.curriculum.model;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;

public class TopCoach {
    private long id;

    @JsonProperty
    private  String kind;

    @JsonProperty
    private  String name;

    private  String fullname;

    private  String occupation;

    private String tag;


    @JsonCreator
    public TopCoach(
            @JsonProperty("id")  long id,
            @JsonProperty("name") String name  ) {

        this.id = id;
        this.name = name;

        this.fullname = this.name;
        this.kind = "HotPractice";
        this.occupation = "teacher";
        this.tag = "Tsinghua, cloud, PhD";
    }

    public TopCoach(long id, String name, String tag) {
        this(id, name);
        this.tag = tag;
    }

    public TopCoach(long id, String name, String occupation, String tag) {
        this(id, name, tag);
        this.occupation = occupation;
    }

    public String getOccupation() {
        return this.occupation;
    }

    public void setOccupation(String occupation) {
        this.occupation = occupation;
    }

    public String getTag() {
        return this.tag;
    }

    public void setTag(String tag) {
        this.tag = tag;
    }

}
