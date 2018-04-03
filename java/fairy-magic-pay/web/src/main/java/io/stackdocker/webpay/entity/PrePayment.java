package io.stackdocker.webpay.entity;

import java.util.UUID;
//import javax.validation.constraints.Size;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

import lombok.Data;

import lombok.Getter;
import lombok.Setter;
import org.hibernate.validator.constraints.NotEmpty;

/**
 * 
 */
@Data
@JsonIgnoreProperties(ignoreUnknown = true)
public class PrePayment {

    @NotEmpty(message = "ID is required.")
    @Getter
    @Setter
    private Long id;

    @NotEmpty(message = "Billing is required.")
    //@Size(min=4, max=35)
    @Getter
    @Setter
    private UUID billing;

    @NotEmpty(message = "Summary is required.")
    @Getter
    @Setter
    private Integer summary;

    @NotEmpty(message = "Payment provider is required.")
    @Getter
    @Setter
    private Provider provider;
    
    public PrePayment(long id, UUID billing, int summary, Provider provider) {
        this.id = Long.valueOf(id);
        this.billing = billing;
        this.summary = summary;
        this.provider = provider;
    }
}