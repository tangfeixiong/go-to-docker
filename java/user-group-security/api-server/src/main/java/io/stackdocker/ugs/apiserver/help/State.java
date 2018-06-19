package io.stackdocker.ugs.apiserver.help;

public enum State {
    AUTHENTICATED_USER("User is authenticated"),
    UNAUTHENTICATED_USER("Invalid credentials"),
    BLOCKED_USER("User is blocked");

    private  final  String message;

    State(String message) {
        this.message = message;
    }

    public String message() {
        return message;
    }
}
