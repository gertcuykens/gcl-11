import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.xml.namespace.QName;
import javax.xml.ws.BindingProvider;
import javax.xml.ws.Service;
import java.io.IOException;
import java.io.PrintWriter;
import java.net.URL;
import java.util.Map;

import com.cdyne.ws.EmailVerNoTestEmail;
import com.cdyne.ws.EmailVerNoTestEmailSoap;

public class SoapServlet2 extends HttpServlet {
    @Override
    public void doGet(HttpServletRequest req, HttpServletResponse resp) throws IOException {
        //Map<String, List<String>> headers = new HashMap<String, List<String>>();
        //headers.put("Username", Collections.singletonList(""));
        //headers.put("Password", Collections.singletonList(""));
        resp.setContentType("text/plain");
        PrintWriter printer = resp.getWriter();
        try {
            String url="http://ws.cdyne.com/emailverify/Emailvernotestemail.asmx?wsdl";
            QName qname = new QName("http://ws.cdyne.com/", "EmailVerNoTestEmail");
            EmailVerNoTestEmail service = new EmailVerNoTestEmail(null, qname);
            EmailVerNoTestEmailSoap port = service.getPort(EmailVerNoTestEmailSoap.class);
            Map<String, Object> ctx = ((BindingProvider) port).getRequestContext();
            ctx.put(BindingProvider.ENDPOINT_ADDRESS_PROPERTY, url);
            //ctx.put(MessageContext.HTTP_REQUEST_HEADERS, headers);
            printer.println(port.verifyEmail("gert.cuykens@gmail.com", "123").getResponseText());
        } catch (Exception e) {
            printer.println("Error occurred while sending SOAP Request to Server");
            e.printStackTrace(printer);
        }
    }

    public static void main(String args[]) {
        try {
            URL url = new URL("http://ws.cdyne.com/emailverify/Emailvernotestemail.asmx?wsdl");
            QName qname = new QName("http://ws.cdyne.com/", "EmailVerNoTestEmail");
            Service service = Service.create(url, qname);
            EmailVerNoTestEmailSoap port = service.getPort(EmailVerNoTestEmailSoap.class);
            System.out.println(port.verifyEmail("gert.cuykens@gmail.com", "123").getResponseText());

            String url2="http://ws.cdyne.com/emailverify/Emailvernotestemail.asmx?wsdl";
            QName qname2 = new QName("http://ws.cdyne.com/", "EmailVerNoTestEmail");
            EmailVerNoTestEmail service2 = new EmailVerNoTestEmail(null, qname2);
            EmailVerNoTestEmailSoap port2 = service2.getPort(EmailVerNoTestEmailSoap.class);
            Map<String, Object> ctx = ((BindingProvider) port).getRequestContext();
            ctx.put(BindingProvider.ENDPOINT_ADDRESS_PROPERTY, url2);
            System.out.println(port2.verifyEmail("gert.cuykens@gmail.com", "123").getResponseText());
        } catch (Exception e) {
            System.err.println("Error occurred while sending SOAP Request to Server");
            e.printStackTrace(System.err);
        }
    }
}
