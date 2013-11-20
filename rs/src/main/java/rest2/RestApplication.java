package rest2;

import org.glassfish.jersey.server.ResourceConfig;
import org.glassfish.jersey.server.ServerProperties;

public class RestApplication extends ResourceConfig {
    public RestApplication() {
        register(new RestBinder());
        register(RestServlet.class);
        //register(rest2.RestService.class);
        property(ServerProperties.METAINF_SERVICES_LOOKUP_DISABLE, true);
        //packages(true, "");
    }
}


/*
@ApplicationPath("/")
public class rest2.RestApplication extends Application {
    @Override
    public Set<Class<?>> getClasses() {
        final Set<Class<?>> classes = new HashSet<Class<?>>();
        classes.add(rest2.RestServlet.class);
        classes.add(rest2.RestService.class);
        return classes;
    }
}
*/
