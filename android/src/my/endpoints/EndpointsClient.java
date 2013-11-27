package my.endpoints;

import com.google.api.client.googleapis.services.json.AbstractGoogleJsonClient;
import com.google.api.client.http.HttpRequestInitializer;
import com.google.api.client.http.HttpTransport;
import com.google.api.client.json.JsonFactory;

public class EndpointsClient extends AbstractGoogleJsonClient {
    public static final String ROOT_URL = "https://gcl-11.appspot.com/_ah/api/";
    public static final String SERVICE_PATH = "rest1/0/";

    public EndpointsClient(HttpTransport transport, JsonFactory jsonFactory, HttpRequestInitializer httpRequestInitializer) {
        super (new Builder(transport, jsonFactory, httpRequestInitializer).setApplicationName("gcl-11"));
    }

    public static final class Builder extends AbstractGoogleJsonClient.Builder {
        public Builder(HttpTransport transport, JsonFactory jsonFactory, HttpRequestInitializer httpRequestInitializer) {
            super(transport, jsonFactory, ROOT_URL, SERVICE_PATH, httpRequestInitializer, false);
        }
        @Override
        public EndpointsClient build() {
            return new EndpointsClient(getTransport(), getJsonFactory(), getHttpRequestInitializer());
        }
    }

    public class Get extends EndpointsClientRequest<Message> {
        protected Get(String get) {
            super(EndpointsClient.this, "GET", get, null, Message.class);
        }
    }

    public class Post extends EndpointsClientRequest<Message> {
        protected Post(String get, Message post) {
            super(EndpointsClient.this, "POST", get, post, Message.class);
        }
    }

    public Get get(String get) throws java.io.IOException {
        Get response = new Get(get);
        initialize(response);
        return response;
    }

    public Post post(String get, Message post) throws java.io.IOException {
        Post response = new Post(get, post);
        initialize(response);
        return response;
    }

}
