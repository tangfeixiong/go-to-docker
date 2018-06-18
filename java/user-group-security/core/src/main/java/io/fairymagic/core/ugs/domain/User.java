package io.fairymagic.core.ugs.domain;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonView;
import org.springframework.boot.autoconfigure.domain.EntityScan;
import org.springframework.format.annotation.DateTimeFormat;

import io.fairymagic.core.ugs.domain.jackson.CommonView;

/**
 * User
 * Table: user
 * 
 * @author tangfx128@gmail.com
 *
 */

@EntityScan
@JsonIgnoreProperties(ignoreUnknown = true)
public class User implements Serializable {

    @JsonView(CommonView.Normal.class)
	private Long id;
    @JsonView(CommonView.Normal.class)
	private String username;
    @JsonView(CommonView.Normal.class)
	private String password;
	private String mobile;
	private String email;
	private Byte status;
    @JsonView(CommonView.Manager.class)
	private Profile profile;
    @JsonView(CommonView.Manager.class)
	private List<RoleBinding> roleBindings;
    @JsonView(CommonView.Manager.class)
	private List<Group> groups;

	
	public User() {
		super();
		profile = new Profile();
		roleBindings = new ArrayList<>();
		groups = new ArrayList<>();
	}
	
	public User(long id, String username) {
		this();
		this.id = Long.valueOf(id);
		this.username = username;
		this.profile.setId(id);
	}

    public User(String username, String password) {
        this();
        this.username = username;
        this.password = password;
    }

	public Long getId() {
		return id;
	}
	public void setId(long id) {
		this.id = Long.valueOf(id);
	}
	public String getUsername() {
		return username;
	}
	public void setUsername(String username) {
		this.username = username;
	}
	public String getPassword() {
		return password;
	}
	public void setPassword(String password) {
		this.password = password;
	}
	public String getMobile() {
		return mobile;
	}
	public void setMobile(String mobile) {
		this.mobile = mobile;
	}
	public String getEmail() {
		return email;
	}
	public void setEmail(String email) {
		this.email = email;
	}
	public Byte getStatus() {
		return status;
	}
	public void setStatus(Byte status) {
		this.status = status;
	}

    public Profile getProfile() {
        return profile;
    }
    public void setProfile(Profile profile) throws UnexpectedException {
	    if ( null == profile) {
	        this.profile = new Profile(id);
        }
	    if (profile != null && profile.getId() != id) {
	        throw new UnexpectedException("Profile must be identical");
        }
        this.profile = profile;
    }

	public List<RoleBinding> getRoleBindings() {
		return roleBindings;
	}
	public void setRoleBindings(List<RoleBinding> roleBindings) {
        this.roleBindings.clear();
	    if ( null != roleBindings) {
            this.roleBindings.addAll(roleBindings);
        }
	}

    public List<Group> getGroups() {
        return groups;
    }
    public void setGroups(List<Group> groups) {
	    this.groups.clear();
	    if ( null != groups ) {
            this.groups.addAll(groups);
        }
    }

//    public Long getPersonScore() {
//		return personScore;
//	}
//	public void setPersonScore(Long personScore) {
//		this.personScore = personScore;
//	}

	@Override
	public String toString() {
		return "User [id=" + id + ", username=" + username + ", mobile=" + mobile
				+ ", email=" + email + ", status=" + status;
	}
	
	
}
