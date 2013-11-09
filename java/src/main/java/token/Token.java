package token;
import com.google.api.client.auth.oauth2.AuthorizationCodeFlow;
import com.google.api.client.auth.oauth2.BearerToken;
import com.google.api.client.auth.oauth2.ClientParametersAuthentication;
import com.google.api.client.auth.oauth2.Credential;
import com.google.api.client.auth.oauth2.TokenResponse;
import com.google.api.client.googleapis.auth.oauth2.GoogleAuthorizationCodeFlow;
import com.google.api.client.googleapis.auth.oauth2.GoogleAuthorizationCodeRequestUrl;
import com.google.api.client.googleapis.auth.oauth2.GoogleClientSecrets;
import com.google.api.client.googleapis.auth.oauth2.GoogleTokenResponse;
import com.google.api.client.http.GenericUrl;
import com.google.api.client.http.javanet.NetHttpTransport;
import com.google.api.client.json.jackson2.JacksonFactory;
//import com.google.api.client.extensions.appengine.http.UrlFetchTransport;

import java.io.IOException;
import java.io.Reader;
import java.util.Arrays;
import java.util.Collection;
import java.util.Map;

public class Token {
    public Token(String key) throws Exception {
        clientSecrets=key;
        initializeFlow();
    }

    public Token(Reader key) throws Exception {
        googleClientSecrets=GoogleClientSecrets.load(new JacksonFactory(),key);
        googleInitializeFlow();
    }

    public AuthorizationCodeFlow flow;

    public GoogleAuthorizationCodeFlow googleFlow;

    public GoogleClientSecrets googleClientSecrets;

    public String clientSecrets;

    public Credential credential;

    public static Collection<String> SCOPE = Arrays.asList(
        "https://www.googleapis.com/auth/userinfo.profile",
        "https://www.googleapis.com/auth/userinfo.email");

    public static String REDIRECT = "http://localhost:8080/oauth2callback";

    public String googleAuthorizeUrl(String state) throws Exception {
        return new GoogleAuthorizationCodeRequestUrl(
            googleClientSecrets,
            REDIRECT,
            SCOPE).setState(state).build();
    }

    public String authorizeUrl(String state) throws Exception {
        return flow.newAuthorizationUrl().setState(state)
            .setRedirectUri(REDIRECT)
            .build();
    }

    public void googleExchangeCode(String authorizationCode) throws Exception {
        GoogleTokenResponse response = googleFlow.newTokenRequest(authorizationCode).setRedirectUri(REDIRECT).execute();
        credential = googleFlow.createAndStoreCredential(response, null);
    }

    public void exchangeCode(String authorizationCode) throws Exception {
        TokenResponse response = flow.newTokenRequest(authorizationCode).setRedirectUri(REDIRECT).execute();
        credential = flow.createAndStoreCredential(response, null);
    }

    private void googleInitializeFlow() throws IOException {
        googleFlow=new GoogleAuthorizationCodeFlow.Builder(
            new NetHttpTransport(),
            new JacksonFactory(),
            googleClientSecrets,
            SCOPE)
            .setAccessType("offline")
            .setApprovalPrompt("force")
            .build();
    }

    private void initializeFlow() throws IOException {
        flow = new AuthorizationCodeFlow.Builder(
            BearerToken.authorizationHeaderAccessMethod(),
            new NetHttpTransport(), //new UrlFetchTransport(),
            new JacksonFactory(),
            new GenericUrl("https://accounts.google.com/o/oauth2/token"),
            new ClientParametersAuthentication("522156758812-09f5qv0e4gqjdjqfocerqcud5m5jutau.apps.googleusercontent.com", clientSecrets),
            "522156758812-09f5qv0e4gqjdjqfocerqcud5m5jutau.apps.googleusercontent.com",
            "https://accounts.google.com/o/oauth2/auth")
            .setScopes(SCOPE) //.setCredentialStore(new AppEngineCredentialStore())
            .build();
    }
}
