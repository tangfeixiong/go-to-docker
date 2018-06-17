package io.stackdocker.ugs.apiserver.domain;

public class NotFoundException extends Exception {

    public NotFoundException(String id) {
        super(id);
    }
}
