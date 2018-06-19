package io.stackdocker.ugs.apiserver.help;

public class NotFoundException extends Exception {

    public NotFoundException(String id) {
        super(id);
    }
}
