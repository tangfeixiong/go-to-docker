package io.stackdocker.ugs.apiserver.service;

import io.fairymagic.core.ugs.domain.User;
import io.stackdocker.ugs.apiserver.dao.UserDao;
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
    public Boolean verifySecureByName(String username, String password) {
        int row = userMapper.verifySecureByName(username, password);
        return Boolean.valueOf(row != 0);
    }

    @Transactional(propagation=Propagation.REQUIRED,rollbackFor={Exception.class})
    public User getById(long id) {
        return new User();
    }

    @Transactional(propagation=Propagation.REQUIRED,rollbackFor={Exception.class})
    public User findByName(String username) {
        return new User();
    }

    public boolean isActiveState(User user) {
        return (user.getStatus().byteValue() & 0x8F) < 0x80;
    }
}
