package io.stackdocker.curriculum.model;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;

public class TopTrainee {
    private long id;

    @JsonProperty
    private  String kind;

    @JsonProperty
    private  String name;

    @JsonCreator
    public TopTrainee(
            @JsonProperty("id")  long id,
            @JsonProperty("name") String name  ) {

        this.id = id;
        this.name = name;
    }
}
