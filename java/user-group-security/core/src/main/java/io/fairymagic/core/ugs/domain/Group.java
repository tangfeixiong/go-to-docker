package io.fairymagic.core.ugs.domain;

public class Group {
    private Long id;
    private Long ownerId;
    private String name;
    private Byte classifier;
    private Byte status;

    public Group() {
        super();
    }

    public Group(long ownerId, String name) {
        this();
        this.ownerId = Long.valueOf(ownerId);
        this.name = name;
    }

    public Long getId() {
        return id;
    }
    public void setId(Long id) {
        this.id = Long.valueOf(id);
    }

    public Long getOwnerId() {
        return ownerId;
    }
    public void setOwnerId(long ownerId) {
        this.ownerId = Long.valueOf(ownerId);
    }

    public String getName() {
        return name;
    }
    public void setName(String name) {
        this.name = name;
    }

    public Byte getClassifier() {
        return classifier;
    }
    public void setClassifier(byte classifier) {
        this.classifier=Byte.valueOf(classifier);
    }

    public Byte getStatus() {
        return status;
    }
    public void setStatus(byte status) {
        this.status = Byte.valueOf(status);
    }
}
