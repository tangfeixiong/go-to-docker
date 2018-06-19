package io.stackdocker.ugs.apiserver.service;

import io.fairymagic.core.ugs.domain.User;
import io.stackdocker.ugs.apiserver.dao.UserDao;
import io.stackdocker.ugs.apiserver.help.State;
import io.stackdocker.ugs.apiserver.mapper.UserMapper;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Propagation;
import org.springframework.transaction.annotation.Transactional;

@Service
public class UserService {
    private static final Logger logger =  LoggerFactory.getLogger(UserService.class);

    @Autowired
    private UserMapper userMapper;

    @Autowired
    UserDao userDao;

    @Transactional(propagation= Propagation.REQUIRED,rollbackFor={Exception.class})
    public Long addOne(User user) {
        long id =  userMapper.addOne(user);
        return Long.valueOf(id);
    }

    @Transactional(propagation=Propagation.REQUIRED,rollbackFor={Exception.class})
    public User getOne(long id) {
        return userMapper.getOne(id);
    }

    @Transactional(propagation=Propagation.REQUIRED,rollbackFor={Exception.class})
    public User findOneByName(String username) {
        return userMapper.findOneByName(username);
    }

    @Transactional(propagation=Propagation.REQUIRED,rollbackFor={Exception.class})
    public Boolean updateStautsByName(String username, byte status) {
        int row = userMapper.updateStatusByName(username, status);
        return row != 0;
    }

    @Transactional(propagation=Propagation.REQUIRED,rollbackFor={Exception.class})
    public Boolean withdrawByName(String username) {
        int row = userMapper.withdrawByName(username);
        return row != 0;
    }

    @Transactional(propagation=Propagation.REQUIRED,rollbackFor={Exception.class})
    public Boolean revokeByName(String username) {
        int row = userMapper.revokeByName(username);
        return row != 0;
    }

    public State verifyWithBasicAuth(String username, String password) {
        User user = this.findOneByName(username);
        if (true == user.getPassword().contentEquals(password)) {
            if (true == isActive(user)) {
                return State.AUTHENTICATED_USER;
            }
            return State.BLOCKED_USER;
        }
         return State.UNAUTHENTICATED_USER;
    }

    public boolean isActive(User user) {
        return (user.getStatus().byteValue() & 0x8F) < 0x80;
    }
}
