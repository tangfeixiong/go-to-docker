package io.stackdocker.iscc.flagserver.api;

import java.util.List;

public interface RefreshConfigCacheService {
    public void add(Integer projectId, RefreshReqResp refreshContext);
    public List<RefreshReqResp> list(Integer projectId);
}