package io.stackdocker.iscc.flagserver.api;

import java.io.Serializable;
import java.text.ParsePosition;
import java.text.SimpleDateFormat;
import java.util.Calendar;
import java.util.Date;
import java.util.GregorianCalendar;
import java.util.Map;
import java.util.HashMap;

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
import org.springframework.beans.factory.annotation.Value;

import cn.com.isc.entity.Config;
import io.stackdocker.iscc.flagserver.domain.RefreshConfig;

@Data
// @Entity(name="config")
@JsonIgnoreProperties({ "serialVersionUID", "minusSeconds" })
// @JsonPropertyOrder({"periodic", "refresh_settings", "state_code", "state_message"})
public class RefreshReqResp implements Serializable {

	private static final long serialVersionUID = 1L;

	// @Id
	// @SequenceGenerator(name="city_generator", sequenceName="city_sequence", initialValue = 23)
	// @GeneratedValue(generator = "city_generator")
	// @GeneratedValue(strategy=GenerationType.AUTO)
	private Integer id;
	
	// @Column(nullable = true)
    @JsonProperty("image_id")	
	private Integer imageId;

	// @Column(nullable = true)
    @JsonProperty("battlefield_id")	
	private Integer projectId;

	// @Column(nullable = true)
    @JsonProperty("name")	
	private String name;

    @JsonProperty("periodic")	
	private Integer periodic;

    @JsonProperty("refreshing_rfc3339")	
    private Date refreshingAt;

    @JsonProperty("rounds")	
	private Integer rounds;

    @JsonProperty("count")	
	private Integer count;
    
    @JsonProperty("begin")	    
    private String begin;
    
    @JsonProperty("elapsed")	    
    private Long elapsed;
    
    @JsonProperty("data_store")	
	private String dataStore;

    @JsonProperty("state_code")	
	private Integer stateCode;

    @JsonProperty("state_message")	
	private String stateMessage;
    
    @JsonProperty("refreshing_info")
    private Map<String, RefreshConfig> info;	
    
    @JsonProperty("config")	
    // @NotEmpty(message = "Config is required.")
    private Config config;
    
    @Value("${minus.seconds}")
    private Integer minusSeconds;
	
    @JsonCreator
	public RefreshReqResp() {
        this.id = 0;
        this.name = "";
        this.imageId = 0;
        this.projectId = 0;
        this.periodic = 0;
        this.refreshingAt = new GregorianCalendar(1999, 11, 30).getTime();
        this.minusSeconds = 10;
        this.rounds = 1;
        this.count = 0;
        this.begin = "";
        this.elapsed = 0L;
        this.dataStore = "";
        this.stateCode = 0;
        this.stateMessage = "";
        this.info = new HashMap<String, RefreshConfig>();
    }
	
	public RefreshReqResp(Integer periodic, Map<String, RefreshConfig> info) {
		super();
		this.setPeriodic(periodic);
		this.setInfo(info);
	}

    @JsonGetter("id")
	public Integer getId() {
		return id;
	}

    @JsonSetter("id")
	public void setId(Integer id) {
		this.id = id;
	}

    @JsonGetter("image_id")
	public Integer getImageId() {
		return imageId;
	}

    @JsonSetter("image_id")
	public void setImageId(Integer id) {
		this.imageId = id;
	}

    @JsonGetter("battlefield_id")
	public Integer getProjectId() {
		return projectId;
	}

    @JsonSetter("battlefield_id")
	public void setProjectId(Integer id) {
		this.projectId = id;
	}
    
	public String getName() {
		return name;
	}
    
	public void setName(String name) {
		this.name = name;
	}
	
	public Integer getPeriodic() {
		return periodic;
	}
	
	public void setPeriodic(Integer periodic) {
        if ( periodic <= 0 ) return;
		this.periodic = periodic;
	}
	
    @JsonGetter("refreshing_rfc3339")
	public String getRefreshingAt() {
        if (refreshingAt != null) {
            String datetime = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss").format(refreshingAt);
            return datetime;
            // return datetime.replaceAll("(\\d\\d)(\\d\\d)$", "$1:$2");
        }
		return "";
	}
	
    @JsonSetter("refreshing_rfc3339")
	public void setRefreshingAt(String datetime) {
        if ( datetime != "" ) {
            try {
		        this.refreshingAt = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss").parse(datetime, new ParsePosition(0));
            } catch (Exception ex) {
                ex.printStackTrace();
            }
        }
	}
    
    public Date getRefreshingDatetime() {
        return refreshingAt;
    }
    
    public void setRefreshingDatetime(Date datetime) {
        this.refreshingAt = datetime;
    }
	
	public Integer getRounds() {
		return rounds;
	}

	public void setRounds(Integer rounds) {
		this.rounds = rounds;
	}
	
    @JsonGetter("count")
	public Integer getCount() {
		return count;
	}
    
    public void incrementsCount() {
        this.count += 1;
    }

    @JsonSetter("count")
	public void setCount(Integer count) {
		this.count = count;
	}

    @JsonGetter("begin")
    public String getBegin() {
        return begin;
    }
    
    @JsonSetter("begin")
    public void setBegin(String begin) {
        this.begin = begin;
    }
    
    @JsonGetter("elapsed")
    public Long getElapsed() {
        return elapsed;
    }
    
    @JsonSetter("elapsed")
    public void setElapsed(Long elapsed) {
        this.elapsed = elapsed;
    }
        
    @JsonGetter("data_store")
	public String getDataStore() {
		return dataStore;
	}
    
    @JsonGetter("data_store")
	public void setDataStore(String dataStore) {
		this.dataStore = dataStore;
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
	
    @JsonGetter("refreshing_info")
    public Map<String, RefreshConfig> getInfo() {
        return info;
    }
    
    public void addRefreshConfig(String key, RefreshConfig value) {
        if ( key.trim().length() != 0 && null != value) {
            info.put(key, value);
        }
    }
	
    @JsonSetter("refreshing_info")
	public void setInfo(Map<String, RefreshConfig> info) {
        if ( null == info ) return;
        
        //Java 8 only, forEach and Lambda
        info.forEach((k,v)->this.info.put(k, v));
	}
    
    public Config getConfig() {
        return config;
    }
    
    public void setConfig(Config config) {
        this.config = config;
    }
	
	public Integer getMinusSeconds() {
		return minusSeconds;
	}

	public void setMinusSeconds(Integer minusSeconds) {
		this.minusSeconds = minusSeconds;
	}
	
	@Override
	public String toString() {
		return "id:" + getId() + ",image_id:" + getImageId() + ",project_id:" + getProjectId() + ",name:" + getName();
	}
}
