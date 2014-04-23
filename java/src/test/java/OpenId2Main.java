import org.openid4java.consumer.ConsumerException;
import org.openid4java.consumer.ConsumerManager;
import org.openid4java.consumer.VerificationResult;
import org.openid4java.discovery.DiscoveryException;
import org.openid4java.discovery.DiscoveryInformation;
import org.openid4java.discovery.Identifier;
import org.openid4java.message.AuthRequest;
import org.openid4java.message.ParameterList;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.UnsupportedEncodingException;
import java.net.URLDecoder;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class OpenId2Main {
    public static String URL = "https://www.google.com/accounts/o8/id";

    public static Map<String, List<String>> getParameterMap(String url) {
        try {
            Map<String, List<String>> params = new HashMap<String, List<String>>();
            String[] urlParts = url.split("\\?");
            if (urlParts.length > 1) {
                String query = urlParts[1];
                for (String param : query.split("&")) {
                    String[] pair = param.split("=");
                    String key = URLDecoder.decode(pair[0], "UTF-8");
                    String value = "";
                    if (pair.length > 1) {
                        value = URLDecoder.decode(pair[1], "UTF-8");
                    }
                    List<String> values = params.get(key);
                    if (values == null) {
                        values = new ArrayList<String>();
                        params.put(key, values);
                    }
                    values.add(value);
                }
            }
            return params;
        } catch (UnsupportedEncodingException ex) {
            throw new AssertionError(ex);
        }
    }

    public static Map<String, String[]> cMap(Map<String, List<String>> params) {
            Map<String, String[]> paraMap = new HashMap<String, String[]>();
            for (Map.Entry<String, List<String>> entry : params.entrySet()) {
                String key = entry.getKey();
                List<String> list = entry.getValue();
                String[] array=new String[list.size()];
                for(int i = 0, n = list.size(); i < n; i++) {
                    array[i]=list.get(i);
                }
                paraMap.put(key,array);
            }
            return paraMap;
    }

    public static void main(String[] args) throws Exception {
        ConsumerManager manager = new ConsumerManager();
        List discoveries = manager.discover(URL);
        DiscoveryInformation discovered = manager.associate(discoveries);
        AuthRequest authReq = manager.authenticate(discovered, "http://localhost:8080/openid");
        String destination = authReq.getDestinationUrl(true);
        System.out.println(destination);
        //resp.sendRedirect(authReq.getDestinationUrl(true));
        //ParameterList openidResp = new ParameterList(request.getParameterMap());
        /*
        Map<String, String[]> map = new HashMap<String, String[]>();
        String[] value = new String[1];
        value[0]="test";
        map.put("test",value);
        */
        //StringBuffer receivingURL = request.getRequestURL();
        //String queryString = request.getQueryString();
        //if (queryString != null && queryString.length() > 0) receivingURL.append("?").append(request.getQueryString());
        System.out.print("Code:");
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        String code = br.readLine();

        ParameterList openidMap = new ParameterList(cMap(getParameterMap(code)));
        VerificationResult verification = manager.verify(code, openidMap, discovered);
        Identifier verified = verification.getVerifiedId();
        System.out.println(verified);
        System.exit(0);
    }
}