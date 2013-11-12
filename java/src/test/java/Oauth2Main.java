import com.google.api.client.http.GenericUrl;
import com.google.api.client.http.HttpRequest;
import com.google.api.client.http.HttpRequestFactory;
import com.google.api.client.http.HttpResponse;
import com.google.api.client.http.HttpTransport;
import com.google.api.client.http.javanet.NetHttpTransport;
import com.google.api.client.json.jackson2.JacksonFactory;
import com.google.api.services.oauth2.Oauth2;
import com.google.api.services.oauth2.model.Tokeninfo;
import com.google.api.services.oauth2.model.Userinfo;
import token.Token;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.net.URL;

public class Oauth2Main {
    private static BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

    public static void main(String[] args) throws Exception {
        Token googleToken = new Token(new FileReader(System.getProperty("user.home") + "/.google.json"));
        System.out.println(googleToken.googleAuthorizeUrl("test"));
        System.out.print("Enter Code: ");
        googleToken.googleExchangeCode(br.readLine());
        userinfo(googleToken);

        Token token = new Token(System.getenv().get("API_KEY"));
        System.out.println(token.authorizeUrl("test"));
        System.out.print("Enter Code: ");
        token.exchangeCode(br.readLine());
        //String c = code(token);
        //System.out.println(c);
        //token.exchangeCode(c);
        userinfo(token);

        System.exit(0);
    }

    public static void userinfo(Token token) throws Exception {
        Oauth2 oauth2 = new Oauth2.Builder(new NetHttpTransport(), new JacksonFactory(), token.credential).setApplicationName("gcl-11").build();
        Tokeninfo tokeninfo = oauth2.tokeninfo().setAccessToken(token.credential.getAccessToken()).execute();
        System.out.println(tokeninfo.toPrettyString());
        /*if (!tokeninfo.getAudience().equals(token.googleClientSecrets.getDetails().getClientId())) {
            System.err.println("ERROR: audience does not match our client ID!");
        }*/
        Userinfo userinfo = oauth2.userinfo().get().execute();
        System.out.println(userinfo.toPrettyString());
    }

    public static String code(Token token) throws Exception {
        HttpTransport HTTP_TRANSPORT = new NetHttpTransport();
        HttpRequestFactory requestFactory = HTTP_TRANSPORT.createRequestFactory();
        URL url = new URL(token.authorizeUrl("test"));
        //InputStream is = ...;
        //String contentType = "plain/text";
        //HttpRequest httpRequest = requestFactory.buildPostRequest(new GenericUrl(url), new InputStreamContent(contentType,is));
        HttpRequest httpRequest = requestFactory.buildGetRequest(new GenericUrl(url));
        HttpResponse response = httpRequest.execute();
        return response.getRequest().getUrl().getFirst("code").toString();
    }
}
