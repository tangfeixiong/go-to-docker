package io.stackdocker.webpay.api;

import java.util.UUID;
//import javax.validation.constraints.NotNull;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import lombok.Data;
import org.hibernate.validator.constraints.NotEmpty;

import io.stackdocker.webpay.entity.Provider;

/**
 * 
 */
@Data
@JsonIgnoreProperties(ignoreUnknown = true)
public class Payment {

    private Long id;

    @NotEmpty(message = "Billing is UUID.")
    private UUID billing;

    @NotEmpty(message = "Booking is fee (int32).")
    private Integer booking;

    @NotEmpty(message = "Qrcode is URI.")
    private String qrcode;
    
    private Boolean notified;
    
    public Payment(long id, UUID billing, int booking, String qrcode, boolean notified) {
        this.id = Long.valueOf(id);
        this.billing = billing;
        this.booking = Integer.valueOf(booking);
        this.qrcode = qrcode;
        this.notified = Boolean.valueOf(notified);
    }
    
    public Payment(long id, UUID billing, int booking, String qrcode) {
        this(id, billing, booking, qrcode, false);
    }
}