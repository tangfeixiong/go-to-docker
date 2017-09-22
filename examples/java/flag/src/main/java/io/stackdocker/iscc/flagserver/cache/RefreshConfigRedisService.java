package io.stackdocker.iscc.flagserver.cache;

import org.springframework.data.redis.core.ListOperations;
import org.springframework.data.redis.core.RedisOperations;
import org.springframework.data.redis.core.SetOperations;
import org.springframework.stereotype.Service;
import javax.annotation.Resource;
import java.time.ZonedDateTime;
import java.time.temporal.ChronoUnit;
import java.util.Calendar;
import java.util.Date;
import java.util.List;

import io.stackdocker.iscc.flagserver.api.RefreshConfigCacheService;
import io.stackdocker.iscc.flagserver.api.RefreshReqResp;

@Service("refreshConfigCacheService")
public class RefreshConfigRedisService implements RefreshConfigCacheService {
    
    @Resource(name = "redisTemplate")
    private ListOperations<Integer, RefreshReqResp> messageList;
    
    @Resource(name = "redisTemplate")
    private RedisOperations<Integer, RefreshReqResp> latestMessageExpiration;
    
    @Override
    public void add(Integer projectId, RefreshReqResp message) {
        messageList.leftPush(projectId, message);
        
        ZonedDateTime zonedDateTime = ZonedDateTime.now();
        Date date = Date.from(zonedDateTime.plus(1, ChronoUnit.MINUTES).toInstant());
        Calendar calendar = Calendar.getInstance();
        calendar.setTime(message.getRefreshingDatetime());
        calendar.add(Calendar.SECOND, message.getPeriodic() * (message.getRounds() - message.getCount()));
        date = calendar.getTime();
        latestMessageExpiration.expireAt(projectId, date);
    }
    
    @Override
    public List<RefreshReqResp> list(Integer projectId) {
        return messageList.range(projectId, 0, -1);
    }
}

