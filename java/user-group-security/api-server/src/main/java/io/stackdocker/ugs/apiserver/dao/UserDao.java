package io.stackdocker.ugs.apiserver.dao;

import io.fairymagic.core.ugs.domain.User;

import org.apache.ibatis.session.SqlSession;
import org.springframework.stereotype.Component;

@Component
public class UserDao {

    private  final SqlSession sqlSession;

    public UserDao(SqlSession sqlSession) {
        this.sqlSession = sqlSession;
    }

    public Long addOne(User user) {
        long id = this.sqlSession.insert("addOne", user);
        return Long.valueOf(id);
    }
}
