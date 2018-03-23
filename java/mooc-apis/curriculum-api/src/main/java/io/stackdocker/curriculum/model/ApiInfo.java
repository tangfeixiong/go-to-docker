package io.stackdocker.curriculum.model;

public class ApiInfo {

    private final long id;
    private final String content;

    public ApiInfo(long id, String content) {
        this.id = id;
        this.content = content;
    }

    public long getId() {
        return id;
    }

    public String getContent() {
        return content;
    }
}
