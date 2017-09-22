package io.stackdocker.iscc.flagserver.domain;

import java.io.Serializable;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;

import lombok.Data;

@Data
// @Entity
public class RefreshFlag implements Serializable {

	private static final long serialVersionUID = 1L;

	// @Id
	// @GeneratedValue(strategy=GenerationType.AUTO)
	private Integer id;
	
	private Integer projectId;
	private Integer teamId;
	private String token;
	private Integer round;
	private String md5String;
	
	private RefreshFlag() {
		super();
	}
	
	public RefreshFlag(Integer id, Integer projectId, Integer teamId, String token, Integer round, String md5) {
		this();
        this.id = id;
		this.projectId = projectId;
		this.teamId = teamId;
		this.token = token;
		this.round = round;
		this.md5String = md5;
	}
}
