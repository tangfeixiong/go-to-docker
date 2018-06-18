package io.fairymagic.core.ugs.domain;

public class Member {
    private Long groupId;
    private Long userId;
    private Byte status;

    public Member() {
        super();
    }
    public Member(int groupId, long userId, byte status) {
        this();
        this.groupId = Long.valueOf(groupId);
        this.userId = Long.valueOf(userId);
        this.status = Byte.valueOf(status);
    }

    public Long getGroupId() {
        return groupId;
    }
    public void setGroupId(long groupId) {
        this.groupId = Long.valueOf(groupId);
    }

    public Long getUserId() {
        return userId;
    }
    public void setUserId(long userId) {
        this.userId = Long.valueOf(userId);
    }

    public Byte getStatus() {
        return status;
    }
    public void setStatus(byte status) {
        this.status = Byte.valueOf(status);
    }
}
