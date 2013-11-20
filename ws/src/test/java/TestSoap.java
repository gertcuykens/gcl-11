import com.cdyne.ws.EmailVerNoTestEmail;
import com.cdyne.ws.EmailVerNoTestEmailSoap;
import org.testng.annotations.Test;

import javax.xml.namespace.QName;
import javax.xml.soap.MessageFactory;
import javax.xml.soap.MimeHeaders;
import javax.xml.soap.SOAPBody;
import javax.xml.soap.SOAPConnection;
import javax.xml.soap.SOAPConnectionFactory;
import javax.xml.soap.SOAPElement;
import javax.xml.soap.SOAPEnvelope;
import javax.xml.soap.SOAPMessage;
import javax.xml.soap.SOAPPart;
import javax.xml.transform.Source;
import javax.xml.transform.Transformer;
import javax.xml.transform.TransformerFactory;
import javax.xml.transform.stream.StreamResult;
import javax.xml.ws.BindingProvider;
import javax.xml.ws.Service;

import java.io.PrintStream;
import java.io.PrintWriter;
import java.net.URL;
import java.util.Map;

public class TestSoap {

    private static SOAPMessage createSOAPRequest() throws Exception {
        MessageFactory messageFactory = MessageFactory.newInstance();
        SOAPMessage soapMessage = messageFactory.createMessage();
        SOAPPart soapPart = soapMessage.getSOAPPart();
        String serverURI = "http://ws.cdyne.com/";
        SOAPEnvelope envelope = soapPart.getEnvelope();
        envelope.addNamespaceDeclaration("example", serverURI);

    /*
    Constructed SOAP Request Message:
    <SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:example="http://ws.cdyne.com/">
        <SOAP-ENV:Header/>
        <SOAP-ENV:Body>
            <example:VerifyEmail>
                <example:email>gert.cuykens@gmail.com</example:email>
                <example:LicenseKey>123</example:LicenseKey>
            </example:VerifyEmail>
        </SOAP-ENV:Body>
    </SOAP-ENV:Envelope>
     */

        SOAPBody soapBody = envelope.getBody();
        SOAPElement soapBodyElem = soapBody.addChildElement("VerifyEmail", "example");
        SOAPElement soapBodyElem1 = soapBodyElem.addChildElement("email", "example");
        soapBodyElem1.addTextNode("gert.cuykens@gmail.com");
        SOAPElement soapBodyElem2 = soapBodyElem.addChildElement("LicenseKey", "example");
        soapBodyElem2.addTextNode("123");
        MimeHeaders headers = soapMessage.getMimeHeaders();
        headers.addHeader("SOAPAction", serverURI  + "VerifyEmail");
        soapMessage.saveChanges();
        return soapMessage;
    }

    private static void printSOAPMessage(SOAPMessage soapResponse, PrintWriter printer) throws Exception {
        TransformerFactory transformerFactory = TransformerFactory.newInstance();
        Transformer transformer = transformerFactory.newTransformer();
        Source sourceContent = soapResponse.getSOAPPart().getContent();
        StreamResult result = new StreamResult(printer);
        transformer.transform(sourceContent, result);
    }

    @Test
    public static void test1() {
        PrintWriter printer = new PrintWriter(System.out);
        try {
            SOAPConnectionFactory soapConnectionFactory = SOAPConnectionFactory.newInstance();
            SOAPConnection soapConnection = soapConnectionFactory.createConnection();
            String url = "http://ws.cdyne.com/emailverify/Emailvernotestemail.asmx?wsdl";
            System.out.println("Request SOAP Message = ");
            SOAPMessage soapRequest = createSOAPRequest();
            printSOAPMessage(soapRequest, printer);
            System.out.println();
            System.out.println("Response SOAP Message = ");
            SOAPMessage soapResponse = soapConnection.call(soapRequest, url);
            printSOAPMessage(soapResponse, printer);
            soapConnection.close();
        } catch (Exception e) {
            System.err.println("Error occurred while sending SOAP Request to Server");
            e.printStackTrace(System.err);
        }
    }

    @Test
    public static void test2() {
        try {
            URL url = new URL("http://ws.cdyne.com/emailverify/Emailvernotestemail.asmx?wsdl");
            QName qname = new QName("http://ws.cdyne.com/", "EmailVerNoTestEmail");
            Service service = Service.create(url, qname);
            EmailVerNoTestEmailSoap port = service.getPort(EmailVerNoTestEmailSoap.class);
            System.out.println(port.verifyEmail("gert.cuykens@gmail.com", "123").getResponseText());
        } catch (Exception e) {
            System.err.println("Error occurred while sending SOAP Request to Server");
            e.printStackTrace(System.err);
        }
    }

    @Test
    public static void test3() {
        try {
            String url="http://ws.cdyne.com/emailverify/Emailvernotestemail.asmx?wsdl";
            QName qname = new QName("http://ws.cdyne.com/", "EmailVerNoTestEmail");
            EmailVerNoTestEmail service = new EmailVerNoTestEmail(null, qname);
            EmailVerNoTestEmailSoap port = service.getPort(EmailVerNoTestEmailSoap.class);
            Map<String, Object> ctx = ((BindingProvider) port).getRequestContext();
            ctx.put(BindingProvider.ENDPOINT_ADDRESS_PROPERTY, url);
            System.out.println(port.verifyEmail("gert.cuykens@gmail.com", "123").getResponseText());
        } catch (Exception e) {
            System.err.println("Error occurred while sending SOAP Request to Server");
            e.printStackTrace(System.err);
        }
    }

}
