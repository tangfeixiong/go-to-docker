/*
  Inspired by
    thymeleafexamples.stsm.business.entities.Type
 */
package /* thymeleafexamples.stsm.business.entities */ io.stackdocker.webpay.entity;


public enum Provider {
    
    ALIPAY("ALIPAY", "via Alibaba Alipay"),
    UNIONPAY("UNIONPAY", "via UnionLink Unionpay"),
    WECHATPAY("WECHATPAY", "via WeChat WxPay");
    
    
    public static final Provider[] ALL = { ALIPAY, UNIONPAY, WECHATPAY };
    
    
    private final String name;
    private final String label;

    Provider(String name, String label) {
        this.name = name;
        this.label = label;
    }
    
    public static Provider forName(final String name) {
        if (name == null) {
            throw new IllegalArgumentException("Name cannot be null for payment provider");
        }
        if (name.toUpperCase().equals("ALIPAY")) {
            return ALIPAY;
        } else if (name.toUpperCase().equals("UNIONPAY")) {
            return UNIONPAY;
        } else if (name.toUpperCase().equals("WECHATPAY")) {
            return WECHATPAY;
        }
        throw new IllegalArgumentException("Name \"" + name + "\" does not correspond to any payment provider");
    }    
    
    public String getName() {
        return this.name;
    }

    public String getLabel() { return this.label; }
    
    @Override
    public String toString() {
        return getName();
    }
    
    
}
