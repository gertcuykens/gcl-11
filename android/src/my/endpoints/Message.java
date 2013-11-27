package my.endpoints;

import com.google.api.client.json.GenericJson;
import com.google.api.client.util.Key;

import java.util.Collection;

public final class Message extends GenericJson {
    @Key public String message;
    public String getMessage() {
        return message;
    }
    public Message setMessage(String message) {
        this.message = message;
        return this;
    }
    @Key public String kind;
    public String getKind() {
        return kind;
    }
    public Message setKind(String kind) {
        this.kind = kind;
        return this;
    }
    @Key public String etag;
    public String getEtag() {
        return etag;
    }
    public Message setEtag(String etag) {
        this.etag = etag;
        return this;
    }
    @Key public Collection<Message> items;
    public Collection<Message> getItems() {
        return items;
    }
    public Message setItems(Collection<Message> items) {
        this.items = items;
        return this;
    }
}
