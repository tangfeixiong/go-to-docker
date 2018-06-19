package io.stackdocker.ugs.apiserver.mapper;

import java.util.List;

import org.apache.ibatis.annotations.Mapper;
import org.springframework.stereotype.Repository;

import io.fairymagic.core.ugs.domain.Profile;
import io.fairymagic.core.ugs.domain.User;

/**
 *
 * Profile of User
 * Table: profile
 *
 * @author tangfx128@gmail.com
 *
 */
@Mapper
public interface ProfileMapper {

    /**
     * add a new profile for user
     * @param profile @see Profile
     * @return id
     */
    public int addOne(Profile profile);

    /**
     * find an profile
     * @param user @see User
     * @return Profile, 0 or 1
     */
    public Profile findOne(User user);

}
