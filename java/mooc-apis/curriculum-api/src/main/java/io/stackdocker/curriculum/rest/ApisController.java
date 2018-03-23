package io.stackdocker.curriculum.rest;


import java.util.concurrent.atomic.AtomicLong;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import io.stackdocker.curriculum.model.ApiInfo;

@RestController
public class ApisController {

    private static final String template = "Hello, %s!";
    private final AtomicLong counter = new AtomicLong();

    @RequestMapping("/apis")
    public ApiInfo[] supportedApis(/*@RequestParam(value="name", defaultValue="World") String name*/) {
        return new ApiInfo[] { 
            new ApiInfo(counter.incrementAndGet(),
                        String.format(template, "demo"))
            };
    }
}
