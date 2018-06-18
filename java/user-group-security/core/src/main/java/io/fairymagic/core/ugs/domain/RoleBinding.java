package io.fairymagic.core.ugs.domain;

public class RoleBinding {

    private Long roleId;

    private Long userId;

    private Long groupId;

    public RoleBinding(){

    }

    public RoleBinding(long roleId, long userId) {
        this.roleId = Long.valueOf(roleId);
        this.userId = Long.valueOf(userId);
    }

    public RoleBinding(long roleId, long userId, long groupId) {
        this.roleId = Long.valueOf(roleId);
        this.userId = Long.valueOf(userId);
        this.groupId = Long.valueOf(groupId);
    }

    public Long getRoleId() {
        return roleId;
    }

    public void setRoleId(long roleId) {
        this.roleId = Long.valueOf(roleId);
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

}
