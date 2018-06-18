package io.fairymagic.core.ugs.domain;

public class Resource {

    private Long id;

    private String name;

    public Resource() {

    }

    public Resource(String name) {
        this.name = name;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

}
