package rest1;

import com.cdyne.ws.EmailVerNoTestEmail;
import com.cdyne.ws.EmailVerNoTestEmailSoap;
import com.google.api.server.spi.config.Api;
import com.google.api.server.spi.config.ApiMethod;
import com.google.appengine.api.datastore.*;
import com.google.appengine.api.users.User;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import javax.inject.Named;
import javax.xml.namespace.QName;
import javax.xml.ws.BindingProvider;

import common.Id;
import common.Response;

@Api(name = "rest1",
     version = "0",
     scopes = {Id.EMAIL_SCOPE},
     clientIds = {Id.WEB_CLIENT_ID, Id.ANDROID_CLIENT_ID, Id.IOS_CLIENT_ID, Id.EXPLORER_ID},
     audiences = {Id.ANDROID_AUDIENCE})
public class Greetings {

    public static DatastoreService datastore = DatastoreServiceFactory.getDatastoreService();
    public static Key map = KeyFactory.createKey("Greetings", "0");

    public static ArrayList<Response> greetings = new ArrayList<Response>();

    static {
        greetings.add(new Response("hello"));
        greetings.add(new Response("goodbye"));
    }



    public Response getGreeting(@Named("id") Integer id) {
        return greetings.get(id);
    }

    public ArrayList<Response> listGreeting() {
        return greetings;
    }

    @ApiMethod(name = "greetings.multiply", httpMethod = "post")
    public Response insertGreeting(@Named("times") Integer times, Response greeting) {
        Response response = new Response();
        StringBuilder responseBuilder = new StringBuilder();
        for (int i = 0; i < times; i++) {
          responseBuilder.append(greeting.getMessage());
        }
        response.setMessage(responseBuilder.toString());
        return response;
    }

    @ApiMethod(name = "greetings.authed", path = "greetings/authed")
    public Response authedGreeting(User user) {
        try {
            return new Response("hello " + user.getEmail());
        } catch (Exception e) {
            return new Response(e.getMessage());
        }
    }

    @ApiMethod(name = "greetings.soap", path = "greetings/soap")
    public Response getSoap() {
        //Map<String, List<String>> headers = new HashMap<String, List<String>>();
        //headers.put("Username", Collections.singletonList(""));
        //headers.put("Password", Collections.singletonList(""));
        try {
            String url="http://ws.cdyne.com/emailverify/Emailvernotestemail.asmx?wsdl";
            QName qname = new QName("http://ws.cdyne.com/", "EmailVerNoTestEmail");
            EmailVerNoTestEmail service = new EmailVerNoTestEmail(null, qname);
            EmailVerNoTestEmailSoap port = service.getPort(EmailVerNoTestEmailSoap.class);
            Map<String, Object> ctx = ((BindingProvider) port).getRequestContext();
            ctx.put(BindingProvider.ENDPOINT_ADDRESS_PROPERTY, url);
            //ctx.put(MessageContext.HTTP_REQUEST_HEADERS, headers);
            return new Response(port.verifyEmail("gert.cuykens@gmail.com", "123").getResponseText());
        } catch (Exception e) {
            return new Response(e.getMessage());
        }
    }

    @ApiMethod(name = "greetings.datastore", path = "greetings/datastore")
    public Response getData () {
        try {
            Entity greeting = new Entity("Greeting", map);
            greeting.setProperty("message", "hello stored");
            datastore.put(greeting);

            Query query = new Query("Greeting", map);
            FetchOptions options = FetchOptions.Builder.withLimit(5);
            List<Entity> list = datastore.prepare(query).asList(options);

            return new Response(list.get(0).getProperty("message").toString());
        } catch (Exception e) {
            return new Response(e.getMessage());
        }
    }
}
