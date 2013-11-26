package my.endpoints;

import com.google.api.client.googleapis.services.AbstractGoogleClientRequest;
import com.google.api.client.googleapis.services.GoogleClientRequestInitializer;
import com.google.api.client.googleapis.services.json.AbstractGoogleJsonClient;
import com.google.api.client.http.HttpRequestInitializer;
import com.google.api.client.http.HttpTransport;
import com.google.api.client.json.JsonFactory;

public class Endpoints extends AbstractGoogleJsonClient {
    public static final String ROOT_URL = "https://gcl-11.appspot.com/_ah/api/";
    public static final String SERVICE_PATH = "rest1/0/";

    public Endpoints(HttpTransport transport, JsonFactory jsonFactory, HttpRequestInitializer httpRequestInitializer) {
        super (new Builder(transport, jsonFactory, httpRequestInitializer).setApplicationName("gcl-11"));
    }

    @Override
    protected void initialize(AbstractGoogleClientRequest<?> httpClientRequest) throws java.io.IOException {
        super.initialize(httpClientRequest);
    }

    public static final class Builder extends AbstractGoogleJsonClient.Builder {
        public Builder(HttpTransport transport, JsonFactory jsonFactory, HttpRequestInitializer httpRequestInitializer) {
            super(transport, jsonFactory, ROOT_URL, SERVICE_PATH, httpRequestInitializer, false);
        }

        @Override
        public Endpoints build() {
            return new Endpoints(getTransport(), getJsonFactory(), getHttpRequestInitializer());
        }

        @Override
        public Builder setRootUrl(String rootUrl) {
            return (Builder) super.setRootUrl(rootUrl);
        }

        @Override
        public Builder setServicePath(String servicePath) {
            return (Builder) super.setServicePath(servicePath);
        }

        @Override
        public Builder setHttpRequestInitializer(HttpRequestInitializer httpRequestInitializer) {
            return (Builder) super.setHttpRequestInitializer(httpRequestInitializer);
        }

        @Override
        public Builder setApplicationName(String applicationName) {
            return (Builder) super.setApplicationName(applicationName);
        }

        @Override
        public Builder setSuppressPatternChecks(boolean suppressPatternChecks) {
            return (Builder) super.setSuppressPatternChecks(suppressPatternChecks);
        }

        @Override
        public Builder setGoogleClientRequestInitializer(GoogleClientRequestInitializer googleClientRequestInitializer) {
            return (Builder) super.setGoogleClientRequestInitializer(googleClientRequestInitializer);
        }
    }

/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////

    public Rest rest() {
        return new Rest();
    }

    public class Rest {

        public GetGreeting getGreeting(String get) throws java.io.IOException {
            GetGreeting response = new GetGreeting(get);
            initialize(response);
            return response;
        }

        public class GetGreeting extends EndpointsRequest<EndpointsResponse> {
            private static final String REST_PATH = "response/";

            protected GetGreeting(String get) {
                super(Endpoints.this, "GET", REST_PATH+get, null, EndpointsResponse.class);
            }

            @Override
            public GetGreeting setAlt(String alt) {
                return (GetGreeting) super.setAlt(alt);
            }

            @Override
            public GetGreeting setFields(String fields) {
                return (GetGreeting) super.setFields(fields);
            }

            @Override
            public GetGreeting setKey(String key) {
                return (GetGreeting) super.setKey(key);
            }

            @Override
            public GetGreeting setOauthToken(String oauthToken) {
                return (GetGreeting) super.setOauthToken(oauthToken);
            }

            @Override
            public GetGreeting setPrettyPrint(Boolean prettyPrint) {
                return (GetGreeting) super.setPrettyPrint(prettyPrint);
            }

            @Override
            public GetGreeting setQuotaUser(String quotaUser) {
                return (GetGreeting) super.setQuotaUser(quotaUser);
            }

            @Override
            public GetGreeting setUserIp(String userIp) {
                return (GetGreeting) super.setUserIp(userIp);
            }

        }

    }

/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////

}


