package io.fairymagic.core.ugs.domain;

public class Rule {

    public Long resourceId;

    public Long roleId;

    public Rule() {

    }

    public Rule(long resourceId, long roleId) {
        this.resourceId = Long.valueOf(resourceId);
        this.roleId = Long.valueOf(roleId);
    }

    public Long getResourceId() {
        return resourceId;
    }

    public void setResourceId(long resourceId) {
        this.resourceId = Long.valueOf(resourceId);
    }

    public Long getRoleId() {
        return roleId;
    }

    public void setRoleId(long roleId) {
        this.roleId = Long.valueOf(roleId);
    }
}
