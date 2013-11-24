package my.endpoints;

import com.google.api.client.json.GenericJson;

public final class EndpointsResponse extends GenericJson {

    @com.google.api.client.util.Key
    public String message;

    public String getMessage() {
        return message;
    }

    public EndpointsResponse setMessage(String message) {
        this.message = message;
        return this;
    }
}
