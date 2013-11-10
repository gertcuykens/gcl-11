import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
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
import java.io.IOException;
//import java.io.PrintStream;
import java.io.PrintWriter;

public class SoapServlet extends HttpServlet {
    @Override
    public void doGet(HttpServletRequest req, HttpServletResponse resp) throws IOException {
        //resp.setContentType("application/soap+xml");
        resp.setContentType("text/plain");
        //PrintStream printer = new PrintStream (resp.getOutputStream());
        PrintWriter printer = resp.getWriter();
        try {
            SOAPConnectionFactory soapConnectionFactory = SOAPConnectionFactory.newInstance();
            SOAPConnection soapConnection = soapConnectionFactory.createConnection();
            String url = "http://ws.cdyne.com/emailverify/Emailvernotestemail.asmx?wsdl";
            printer.println("Request SOAP Message = ");
            SOAPMessage soapRequest = createSOAPRequest();
            printSOAPMessage(soapRequest, printer);
            printer.println();
            printer.println("Response SOAP Message = ");
            SOAPMessage soapResponse = soapConnection.call(soapRequest, url);
            printSOAPMessage(soapResponse, printer);
            soapConnection.close();
        } catch (Exception e) {
            printer.println("Error occurred while sending SOAP Request to Server");
            e.printStackTrace(printer);
        }
    }

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

    public static void main(String args[]) {
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

}
