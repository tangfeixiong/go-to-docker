package io.stackdocker.iscc.flagserver.cache;

import java.util.List;

import org.springframework.stereotype.Repository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.CrudRepository;
import org.springframework.data.repository.query.Param;

import cn.com.isc.entity.Flag;

@Repository("FlagRepository1")
public interface  FlagRepository extends CrudRepository<Flag, Integer>{
	@Query("SELECT f FROM Flag f WHERE env = :env")
    List<Flag> search(@Param("env") Integer projectId );

	@Query("SELECT f FROM Flag f WHERE env = :env AND team_no = :team_no")
    List<Flag> search(@Param("env") Integer projectId, @Param("team_no") Integer teamId);

	@Query("SELECT f FROM Flag f WHERE env = :env AND team_no = :team_no AND round = :round")
	Flag search(@Param("env") Integer projectId, @Param("team_no") Integer teamId, @Param("round") Integer round);

	@Query("SELECT f FROM Flag f WHERE env = :env AND round >= :round")
    List<Flag> find(@Param("env") Integer projectId, @Param("round") Integer round);
}