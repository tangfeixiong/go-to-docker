package io.stackdocker.curriculum.rest;


import java.util.concurrent.atomic.AtomicLong;

import io.stackdocker.curriculum.model.*;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class HighlightsController {

    private static final String template = "Hello, %s!";
    private final AtomicLong counter = new AtomicLong();

    @RequestMapping("/api/v1alpha/namespaces/default/top-classes")
    public TopClass[] topClasses(/*@RequestParam(value="name", defaultValue="World") String name*/) {
        return new TopClass[] {
                new TopClass(counter.incrementAndGet(),
                        "Kubernetes 101: Deploy with kubeadm"),
                new TopClass(counter.incrementAndGet(),
                        "Docker 101: build, ship and run",
                        "top1, class"),
                new TopClass(counter.incrementAndGet(),
                        "CNCF 101: KaaS, Kubernetes as a Service",
                        "top1, lesson"),
                new TopClass(counter.incrementAndGet(),
                        "Go 101: Working with concurrency",
                        "top2"),
                new TopClass(counter.incrementAndGet(),
                        "Machine Learning 101: Learning Tensorflow", "" +
                        "top2, class"),
                new TopClass(counter.incrementAndGet(),
                        "Block Chain 101: Implementing decentrialized ledger via ethereum",
                        "lesson")
        };
    }

    @RequestMapping("/api/v1alpha/namespaces/default/top-practices")
    public TopPractice[] topPractices(/*@RequestParam(value="name", defaultValue="World") String name*/) {
        return new TopPractice[] {
                new TopPractice(counter.incrementAndGet(),
                        "Kubernetes hands-on: Deploy with Kubeadm",
                        "kubeadm 1.9, HA, Fedora26"),
                new TopPractice(counter.incrementAndGet(),
                        "Docker hands-on: Build spring boot image",
                        "video",
                        "java1.8, tomcat"),
                new TopPractice(counter.incrementAndGet(),
                        "Go hands-on: channel and sync",
                        "go1.9, MAC")
        };
    }

    @RequestMapping("/api/v1alpha/namespaces/default/top-coaches")
    public TopCoach[] topCoaches(/*@RequestParam(value="name", defaultValue="World") String name*/) {
        return new TopCoach[] {
                new TopCoach(counter.incrementAndGet(),
                        "Tim Cook",
                        "CEO, phone, swift"),
                new TopCoach(counter.incrementAndGet(),
                        "Jeff Bezos",
                        "PKU",
                        "CEO, cloud, c"),
                new TopCoach(counter.incrementAndGet(),
                        "Larry Page",
                        "ZJU",
                        "CEO, AI, go"),
                new TopCoach(counter.incrementAndGet(),
                        "Satya Nadella",
                        "SJTU",
                        "CEO, OS, cs"),
                new TopCoach(counter.incrementAndGet(),
                        "Pony Ma",
                        "SZU",
                        "CEO, SN, js"),
                new TopCoach(counter.incrementAndGet(),
                        "Robbin Li",
                        "PKU",
                        "CEO, SE, cpp"),
                new TopCoach(counter.incrementAndGet(),
                        "Jack Ma",
                        "HZNU",
                        "CEO, b2b, Java"),
                new TopCoach(counter.incrementAndGet(),
                        "Elon Musk",
                        "FDU",
                        "CEO, AP, python")
        };
    }

    @RequestMapping("/api/v1alpha/namespaces/default/top-trainees")
    public TopTrainee[] topTrainees(/*@RequestParam(value="name", defaultValue="World") String name*/) {
        return new TopTrainee[] {
                new TopTrainee(counter.incrementAndGet(),
                        String.format(template, "Steve Jobs"))
        };
    }

    @RequestMapping("/api/v1alpha/namespaces/default/top-hires")
    public TopHire[] topHires(/*@RequestParam(value="name", defaultValue="World") String name*/) {
        return new TopHire[] {
                new TopHire(counter.incrementAndGet(),
                        String.format(template, "Facebook"))
        };
    }

    @RequestMapping("/api/v1alpha/namespaces/default/top-hunters")
    public TopHunter[] topHunters(/*@RequestParam(value="name", defaultValue="World") String name*/) {
        return new TopHunter[] {
                new TopHunter(counter.incrementAndGet(),
                        String.format(template, "Bill Gates"))
        };
    }

    /**
     *  Show category for detailed entry, for example clicking More...
     *  Feixiong Tang (tangfx128@gmail.com)
      */
    @RequestMapping("/api/v1alpha/namespaces/default/headlines")
    public Headline[] headlines(/*@RequestParam(value="name", defaultValue="World") String name*/) {
        return new Headline[] {
                new Headline(counter.incrementAndGet(),
                        String.format(template, "Donald Trump, China surplus, steel and aluminum tariffs, trade war"))
        };
    }

}
