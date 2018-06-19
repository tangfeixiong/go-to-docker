package io.stackdocker.ugs.apiserver.mapper;

import io.fairymagic.core.ugs.domain.User;

import org.apache.ibatis.annotations.Param;

public interface UserMapper {

    /**
     * add a new user
     * @param user @see User
     * @return id java.lang.Long
     */
    public long addOne(User user);

    /**
     * verify an user
     * @param username java.lang.String
     * @param  password java.lang.String
     * @return count, 0 or 1
     */
    public int verifySecureByName(@Param("username") String username, @Param("password") String password);

    /**
     * verify an user
     * @param email java.lang.String
     * @param  password java.lang.String
     * @return count, 0 or 1
     */
    public int verifySecureByEmail(@Param("email") String email, @Param("password") String password);
}
