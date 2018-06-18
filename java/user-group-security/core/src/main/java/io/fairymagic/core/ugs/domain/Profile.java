package io.fairymagic.core.ugs.domain;

import java.io.Serializable;

/**
 * Profile of User
 * Table: profile
 *
 * @author tangfx128@gmail.com
 *
 */

public class Profile implements Serializable {
	
	private Long id;
	private String fullname;
	private String userName;
	private String userRow;
	private String uerMobile;
	private String userRemark;

	public Profile() {
	    super();
    }
    public Profile(long id) {
	    this();
	    this.id = Long.valueOf(id);
    }

	public Long getProfileId() {
		return id;
	}
	public void setProfileId(Long id) {
        if ( null != id && 0 < id.intValue()) {
            this.id = id;
        }
	}

	public Long getId() {
		return id;
	}
	public void setId(long id) {
		this.id = Long.valueOf(id);
	}

	public String getFullname() {
		return fullname;
	}
	public void setFullname(String fullname) {
		this.fullname = fullname;
	}

	public String getUserName() {
		return userName;
	}
	public void setUserName(String userName) {
		this.userName = userName;
	}
	public String getUserRow() {
		return userRow;
	}
	public void setUserRow(String userRow) {
		this.userRow = userRow;
	}
	public String getUerMobile() {
		return uerMobile;
	}
	public void setUerMobile(String uerMobile) {
		this.uerMobile = uerMobile;
	}
	public String getUserRemark() {
		return userRemark;
	}
	public void setUserRemark(String userRemark) {
		this.userRemark = userRemark;
	}
	@Override
	public String toString() {
		return "Profile [profileId=" + id + ", fullname=" + fullname + ", userName=" + userName + ", userRow="
				+ userRow + ", uerMobile=" + uerMobile + ", userRemark=" + userRemark + "]";
	}
	
	
	
	
}
