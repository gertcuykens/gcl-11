package rest2;

import javax.inject.Inject;
import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;

@Path("/")
public class RestServlet {
    @Inject
    RestService restService;

    @GET
    @Path("/json")
    @Produces(MediaType.APPLICATION_JSON)
    public String getHelloWorldJSON() {
        return "{\"result\":\"" + restService.createHelloMessage("World") + "\"}";
    }
}
