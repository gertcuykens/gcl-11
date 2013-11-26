package my.endpoints;

import com.google.api.client.googleapis.services.json.AbstractGoogleJsonClientRequest;
import com.google.api.client.http.HttpHeaders;
import com.google.api.client.http.UriTemplate;

public abstract class EndpointsRequest<T> extends AbstractGoogleJsonClientRequest<T> {

    /**
     * @param client Google client
     * @param method HTTP Method
     * @param uriTemplate URI template for the path relative to the base URL. If it starts with a "/"
     * the base path from the base URL will be stripped out. The URI template can also be a
     * full URL. URI template expansion is done using
     * {@link UriTemplate#expand(String, String, Object, boolean)}
     * @param content A POJO that can be serialized into JSON or {@code null} for none
     * @param responseClass response class to parse into
     */
    public EndpointsRequest(Endpoints client, String method, String uriTemplate, Object content, Class<T> responseClass) {
        super(client, method, uriTemplate, content, responseClass);
    }

    /** Data format for the response. */
    @com.google.api.client.util.Key
    private String alt;

    /**
     * Data format for the response. [default: json]
     */
    public String getAlt() {
        return alt;
    }

    /** Data format for the response. */
    public EndpointsRequest<T> setAlt(String alt) {
        this.alt = alt;
        return this;
    }

    /** Selector specifying which fields to include in a partial response. */
    @com.google.api.client.util.Key
    private String fields;

    /**
     * Selector specifying which fields to include in a partial response.
     */
    public String getFields() {
        return fields;
    }

    /** Selector specifying which fields to include in a partial response. */
    public EndpointsRequest<T> setFields(String fields) {
        this.fields = fields;
        return this;
    }

    /**
     * API key. Your API key identifies your project and provides you with API access, quota, and
     * reports. Required unless you provide an OAuth 2.0 token.
     */
    @com.google.api.client.util.Key
    private String key;

    /**
     * API key. Your API key identifies your project and provides you with API access, quota, and
     * reports. Required unless you provide an OAuth 2.0 token.
     */
    public String getKey() {
        return key;
    }

    /**
     * API key. Your API key identifies your project and provides you with API access, quota, and
     * reports. Required unless you provide an OAuth 2.0 token.
     */
    public EndpointsRequest<T> setKey(String key) {
        this.key = key;
        return this;
    }

    /** OAuth 2.0 token for the current user. */
    @com.google.api.client.util.Key("oauth_token")
    private String oauthToken;

    /**
     * OAuth 2.0 token for the current user.
     */
    public String getOauthToken() {
        return oauthToken;
    }

    /** OAuth 2.0 token for the current user. */
    public EndpointsRequest<T> setOauthToken(String oauthToken) {
        this.oauthToken = oauthToken;
        return this;
    }

    /** Returns response with indentations and line breaks. */
    @com.google.api.client.util.Key
    private Boolean prettyPrint;

    /**
     * Returns response with indentations and line breaks. [default: true]
     */
    public Boolean getPrettyPrint() {
        return prettyPrint;
    }

    /** Returns response with indentations and line breaks. */
    public EndpointsRequest<T> setPrettyPrint(Boolean prettyPrint) {
        this.prettyPrint = prettyPrint;
        return this;
    }

    /**
     * Available to use for quota purposes for server-side applications. Can be any arbitrary string
     * assigned to a user, but should not exceed 40 characters. Overrides userIp if both are provided.
     */
    @com.google.api.client.util.Key
    private String quotaUser;

    /**
     * Available to use for quota purposes for server-side applications. Can be any arbitrary string
     * assigned to a user, but should not exceed 40 characters. Overrides userIp if both are provided.
     */
    public String getQuotaUser() {
        return quotaUser;
    }

    /**
     * Available to use for quota purposes for server-side applications. Can be any arbitrary string
     * assigned to a user, but should not exceed 40 characters. Overrides userIp if both are provided.
     */
    public EndpointsRequest<T> setQuotaUser(String quotaUser) {
        this.quotaUser = quotaUser;
        return this;
    }

    /**
     * IP address of the site where the request originates. Use this if you want to enforce per-user
     * limits.
     */
    @com.google.api.client.util.Key
    private String userIp;

    /**
     * IP address of the site where the request originates. Use this if you want to enforce per-user
     * limits.
     */
    public String getUserIp() {
        return userIp;
    }

    /**
     * IP address of the site where the request originates. Use this if you want to enforce per-user
     * limits.
     */
    public EndpointsRequest<T> setUserIp(String userIp) {
        this.userIp = userIp;
        return this;
    }

    @Override
    public final Endpoints getAbstractGoogleClient() {
        return (Endpoints) super.getAbstractGoogleClient();
    }

    @Override
    public EndpointsRequest<T> setDisableGZipContent(boolean disableGZipContent) {
        return (EndpointsRequest<T>) super.setDisableGZipContent(disableGZipContent);
    }

    @Override
    public EndpointsRequest<T> setRequestHeaders(HttpHeaders headers) {
        return (EndpointsRequest<T>) super.setRequestHeaders(headers);
    }
}
