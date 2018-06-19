package io.stackdocker.ugs.apiserver.service;

import io.fairymagic.core.ugs.domain.Role;
import io.fairymagic.core.ugs.domain.RoleBinding;
import io.fairymagic.core.ugs.domain.User;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Propagation;
import org.springframework.transaction.annotation.Transactional;

import java.util.Collection;
import java.util.Set;
import java.util.TreeSet;

@Service
public class RoleService {
    private static final Logger logger =  LoggerFactory.getLogger(UserService.class);



    @Transactional(propagation= Propagation.REQUIRED,rollbackFor={Exception.class})
    public Set<Role> getWithBindings(Collection<RoleBinding> bindings) {
        return new TreeSet<>();
    }

}
