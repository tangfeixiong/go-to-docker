package io.stackdocker.iscc.flagserver.cache;

import java.util.ArrayList;
import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.cache.annotation.Cacheable;
import org.springframework.cache.annotation.CacheEvict;
import org.springframework.core.env.Environment;
import org.springframework.stereotype.Service;

import cn.com.isc.entity.Flag;

import io.stackdocker.iscc.flagserver.domain.RefreshFlag;


@Service
public class FlagCacheableService {

	@Autowired
    @Qualifier("FlagRepository1")
	private FlagRepository repo;
    
    @Autowired
    private Environment env;

    @CacheEvict(cacheNames = "initialflags", beforeInvocation = true)
    @Cacheable(cacheNames = "initialflags", key = "#projectId")
    public List<RefreshFlag> search( final Integer projectId) {
        System.out.println("Search flags repository by project: " + projectId);
        List<Flag> items = repo.search(projectId);
        List<RefreshFlag> result = new ArrayList<RefreshFlag>(items.size());
        items.forEach(item->{
            result.add(new RefreshFlag(item.getId(), item.getEnv(), item.getTeamNo(),
                    item.getToken(), item.getRound(), item.getMd5String()));
        });
        return result;
    }

    @CacheEvict(cacheNames = "followingflags", beforeInvocation = true)
    @Cacheable(cacheNames = "followingflags", key = "{#projectId, #round}")
    public List<RefreshFlag> find( final Integer projectId, Integer round) {
        System.out.println("Search flags repository by project: " + projectId + " and round: " + round);
        List<Flag> items = repo.find(projectId, round);
        List<RefreshFlag> result = new ArrayList<RefreshFlag>(items.size());
        items.forEach(item->{
            result.add(new RefreshFlag(item.getId(), item.getEnv(), item.getTeamNo(),
                    item.getToken(), item.getRound(), item.getMd5String()));
        });
        return result;
    }

    public RefreshFlag search( final Integer projectId, final Integer teamId, final Integer round ) {
        String cacheType = env.getProperty("spring.cache.type");
        String redisCluster = env.getProperty("spring.redis.cluster.nodes");
        String redisURL = env.getProperty("spring.redis.url");
        String redisHost = env.getProperty("spring.redis.host");
        String redisSentinel = env.getProperty("spring.redis.sentinel.master");
        if ( cacheType == "" || cacheType == "none" || redisHost == "" ) {
            Flag item = repo.search( projectId, teamId, round );
            if (item != null) return new RefreshFlag(item.getId(), item.getEnv(), item.getTeamNo(),
                    item.getToken(), item.getRound(), item.getMd5String());
        }

        List<RefreshFlag> items = search(projectId);
//        items.forEach( item->{
//        	if (item.getTeamId() == teamId && item.getRound() == round ) return item;
//        });
        for (RefreshFlag item: items) {
            if (item.getTeamId() == teamId && item.getRound() == round ) return item;
        }       

        items = find(projectId, round);
        for (RefreshFlag item: items) {
            if (item.getTeamId() == teamId && item.getRound() == round ) return item;
        }
        
        return null;
    }
}