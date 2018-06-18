
package io.stackdocker.core.cli;

import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class App implements CommandLineRunner {

	@Override
	public void run(String... args) {
		if (args.length > 0 && args[0].equals("exitcode")) {
			throw new ExitException();
		}
		System.out.println(args);
	}

	public static void main(String[] args) {
		SpringApplication.run(App.class, args);
	}

}
