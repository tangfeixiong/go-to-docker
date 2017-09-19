package io.stackdocker.iscc.flagserver;

import com.fasterxml.classmate.TypeResolver;
import org.joda.time.LocalDate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.context.request.async.DeferredResult;
import org.springframework.web.servlet.config.annotation.EnableWebMvc;

import springfox.documentation.annotations.ApiIgnore;
import springfox.documentation.builders.ApiInfoBuilder;
import springfox.documentation.builders.AuthorizationScopeBuilder;
import springfox.documentation.builders.ImplicitGrantBuilder;
import springfox.documentation.builders.OAuthBuilder;
import springfox.documentation.builders.ParameterBuilder;
import springfox.documentation.builders.PathSelectors;
import springfox.documentation.builders.RequestHandlerSelectors;
import springfox.documentation.builders.ResponseMessageBuilder;
import springfox.documentation.schema.ModelRef;
import springfox.documentation.schema.WildcardType;
import springfox.documentation.service.ApiInfo;
import springfox.documentation.service.ApiKey;
import springfox.documentation.service.AuthorizationScope;
import springfox.documentation.service.BasicAuth;
import springfox.documentation.service.GrantType;
import springfox.documentation.service.LoginEndpoint;
import springfox.documentation.service.SecurityReference;
import springfox.documentation.service.SecurityScheme;
import springfox.documentation.service.Tag;
import springfox.documentation.spi.DocumentationType;
import springfox.documentation.spi.service.contexts.SecurityContext;
import springfox.documentation.spring.web.plugins.Docket;
import springfox.documentation.swagger.web.ApiKeyVehicle;
import springfox.documentation.swagger.web.SecurityConfiguration;
import springfox.documentation.swagger.web.UiConfiguration;
import springfox.documentation.swagger2.annotations.EnableSwagger2;

import java.time.*;
import java.util.List;

import io.stackdocker.iscc.flagserver.controller.RefreshController;

import static com.google.common.base.Predicates.*;
import static com.google.common.collect.Lists.*;
import static springfox.documentation.schema.AlternateTypeRules.newRule;

@Configuration
//@EnableWebMvc
@EnableSwagger2
@ComponentScan(basePackageClasses = {
	RefreshController.class,
})
public class SwaggerConfig {

    @Autowired
    private TypeResolver typeResolver;

    private ApiKey apiKey() {
        return new ApiKey("mykey", "api_key", "header");
    }

    @Bean
    public Docket documentation() {
        return new Docket(DocumentationType.SWAGGER_2)
            .select()
                .apis(RequestHandlerSelectors.any())
                .paths(PathSelectors.any())
                .build()
            .pathMapping("/")
            .directModelSubstitute(LocalDate.class,
			    String.class)
            .genericModelSubstitutes(ResponseEntity.class)
            .alternateTypeRules(
                newRule(typeResolver.resolve(DeferredResult.class,
                    typeResolver.resolve(ResponseEntity.class, WildcardType.class)),
                typeResolver.resolve(WildcardType.class)))
            .useDefaultResponseMessages(false)
		    .globalResponseMessage(RequestMethod.GET,
                newArrayList(new ResponseMessageBuilder()
                    .code(500)
                    .message("500 message")
                    .responseModel(new ModelRef("Error"))
                    .build()))
            .securitySchemes(newArrayList(apiKey()))
	        // .securityContexts(newArrayList(securityContext()))
	        .enableUrlTemplating(true)
	        // .globalOperationParameters(
	        //     newArrayList(new ParameterBuilder()
	        //         .name("someGlobalParameter")
	        //         .description("Description of someGlobalParameter")
	        //         .modelRef(new ModelRef("string"))
	        //         .parameterType("query")
	        //         .required(true)
	        //         .build()))
	        .tags(new Tag("Flag Service", "All apis relating to Flag")) 
	        // .additionalModels(typeResolver.resolve(AdditionalModel.class)) 
	        ;
    }

	private SecurityContext securityContext() {
	  return SecurityContext.builder()
	      .securityReferences(defaultAuth())
	      .forPaths(PathSelectors.regex("/anyPath.*"))
	      .build();
	}
	
	List<SecurityReference> defaultAuth() {
	  AuthorizationScope authorizationScope
	      = new AuthorizationScope("global", "accessEverything");
	  AuthorizationScope[] authorizationScopes = new AuthorizationScope[1];
	  authorizationScopes[0] = authorizationScope;
	  return newArrayList(
	      new SecurityReference("mykey", authorizationScopes));
	}
	
	@Bean
	SecurityConfiguration security() {
	  return new SecurityConfiguration(
	      "test-app-client-id",
	      "test-app-client-secret",
	      "test-app-realm",
	      "test-app",
	      "apiKey",
	      ApiKeyVehicle.HEADER, 
	      "api_key", 
	      "," /*scope separator*/);
	}
	
	@Bean
	UiConfiguration uiConfig() {
	  return new UiConfiguration(
	      "validatorUrl",// url
	      "none",       // docExpansion          => none | list
	      "alpha",      // apiSorter             => alpha
	      "schema",     // defaultModelRendering => schema
	      UiConfiguration.Constants.DEFAULT_SUBMIT_METHODS,
	      false,        // enableJsonEditor      => true | false
	      true,         // showRequestHeaders    => true | false
	      60000L);      // requestTimeout => in milliseconds, defaults to null (uses jquery xh timeout)
	}

    private ApiInfo apiInfo() {
        return new ApiInfoBuilder()
                .title("Springfox petstore API")
                .description("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum " +
                        "has been the industry's standard dummy text ever since the 1500s, when an unknown printer "
                        + "took a " +
                        "galley of type and scrambled it to make a type specimen book. It has survived not only five " +
                        "centuries, but also the leap into electronic typesetting, remaining essentially unchanged. " +
                        "It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum " +
                        "passages, and more recently with desktop publishing software like Aldus PageMaker including " +
                        "versions of Lorem Ipsum.")
                .termsOfServiceUrl("http://springfox.io")
                .contact("springfox")
                .license("Apache License Version 2.0")
                .licenseUrl("https://github.com/springfox/springfox/blob/master/LICENSE")
                .version("2.0")
                .build();
    }
}
