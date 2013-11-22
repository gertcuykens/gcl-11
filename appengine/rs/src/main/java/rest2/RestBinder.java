package rest2;

import org.glassfish.hk2.utilities.binding.AbstractBinder;

public class RestBinder extends AbstractBinder {
    @Override
    protected void configure() {
        bind(RestService.class).to(RestService.class);
    }
}
