package io.stackdocker.iscc.flagserver.domain;

import java.io.Serializable;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonGetter;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.annotation.JsonSetter;
import lombok.Data;
import org.hibernate.validator.constraints.NotEmpty;

import cn.com.isc.entity.Flag;

@Data
// @Entity(name="flag")
@JsonIgnoreProperties({ "serialVersionUID", "imageId" })
public class RefreshConfig implements Serializable {

	private static final long serialVersionUID = 1L;

	// @Id
	// @SequenceGenerator(name="city_generator", sequenceName="city_sequence", initialValue = 23)
	// @GeneratedValue(generator = "city_generator")
	// @GeneratedValue(strategy=GenerationType.AUTO)
	// private Integer id;
	
	// @Column(nullable = true)
	// private Integer env;

	// @Column(nullable = true)
	// private Integer TeamNo;

	// @Column(nullable = true)
	// private String Token;

	// @Column(nullable = true)
	// private Integer Round;

	// @Column(nullable = true)
	// private String md5String;

    @JsonProperty("container_id")	
    private Integer containerId;
    
    @JsonProperty("refresh_config_id")
    private Integer configId;
	
    @JsonProperty("team_id")	
    private Integer teamId;

    @JsonProperty("name")	
    private String name;	
    
    @JsonProperty("sub_path")
    private String subPath;

    @JsonProperty("state_code")	
	private Integer stateCode;

    @JsonProperty("state_message")	
	private String stateMessage;

    // @NotEmpty(message = "Flag is required.")
    @JsonProperty("flag")	
	private Flag flag;
    
    private Integer projectId;
    private Integer imageId;
	
    @JsonCreator
	public RefreshConfig() {	
        containerId = 0;
        configId = 0;
        teamId = 0;
        name = "";
        subPath = "";
        stateCode = 0;
        stateMessage = "";
    }
	
	public RefreshConfig(Integer env, String token, Integer teamNumber, String md5, Integer round) {
		super();
        flag = new Flag(env, token, teamNumber, md5, round);
        imageId = env;
        teamId = teamNumber;
	}
    
    @JsonGetter("container_id")
    public Integer getContainerId() {
        return containerId;
    }
    
    @JsonSetter("container_id")
    public void setContainerId(Integer id) {
        this.containerId = id;
    }
    
    @JsonGetter("refresh_config_id")
    public Integer getConfigId() {
        return configId;
    }
    
    @JsonSetter("refresh_config_id")
    public void setConfigId(Integer id) {
        this.configId = id;
    }
    
    @JsonGetter("team_id")
    public Integer getTeamId() {
        return teamId;
    }
    
    @JsonSetter("team_id")
    public void setTeamId(Integer id) {
        this.teamId = id;
    }
	
    @JsonGetter("name")
	public String getName() {
		return name;
	}
	
    @JsonSetter("name")
	public void setName(String name) {
		this.name = name;
	}
    
    @JsonGetter("sub_path")
    public String getSubPath() {
        return subPath;
    }
    
    @JsonSetter("sub_path")
    public void setSubPath(String subPath) {
        this.subPath = subPath;
    }	
	
    @JsonGetter("state_code")
	public Integer getStateCode() {
		return stateCode;
	}

    @JsonSetter("state_code")
	public void setStateCode(Integer stateCode) {
		this.stateCode = stateCode;
	}

    @JsonGetter("state_message")
	public String getStateMessage() {
		return stateMessage;
	}

    @JsonSetter("state_message")
	public void setStateMessage(String stateMessage) {
		this.stateMessage = stateMessage;
	}
    
    public Flag getFlag() {
        return flag;
    }
    
    public void setFlag(Flag flag) {
        this.flag = flag;
    }	
    
    public void setRefreshFlag(RefreshFlag flag) {
        if ( this.flag != null && flag != null) {
            // this.flag.setId(flag.getId());
            // this.flag.setEnv(flag.getProjectId());
            // this.flag.setTeamNo(flag.getTeamId());
            this.flag.setToken(flag.getToken());
            this.flag.setRound(flag.getRound());
            this.flag.setMd5String(flag.getMd5String());
        }
    }
    
    public Integer getProjectId() {
        return projectId;
    }
    
    public void setProjectId(Integer id) {
        this.projectId = id;
    }	
    
    public Integer getImageId() {
        return imageId;
    }
    
    public void setImageId(Integer id) {
        this.imageId = id;
    }	
	
	@Override
	public String toString() {
		return "container_id:" + getContainerId() 
                + ",config_id:" + getConfigId() + ",team_id:" + getTeamId() 
                + ",name:" + getName() + ",sub_path:" + getSubPath();
	}
}
