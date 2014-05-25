package util
/*
import (
    "encoding/xml"
    "bytes"
    "encoding/xml"
    "io/ioutil"
    "net/http"
    "appengine/urlfetch"
)

const SERVER = "http://ws.cdyne.com/emailverify/Emailvernotestemail.asmx"
const QUERY =
`<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:example="http://ws.cdyne.com/">
    <SOAP-ENV:Header/>
    <SOAP-ENV:Body>
        <example:VerifyEmail>
            <example:email>gert.cuykens@gmail.com</example:email>
            <example:LicenseKey>123</example:LicenseKey>
        </example:VerifyEmail>
    </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`

type SoapResult struct {
    ResponseText string
    ResponseCode int
    LastMailServer string
    GoodEmail bool
}

type SoapResponse struct {
    VerifyEmailResult SoapResult
}

type SoapBody struct {
    VerifyEmailResponse SoapResponse
}

type SoapEnvelope struct {
    XMLName xml.Name
    Body    SoapBody
}

func (gs *Service) Soap(r *http.Request, req *NoRequest, resp *Response) error {
	c := endpoints.NewContext(r)
	httpClient := urlfetch.Client(c)
	respx, err := httpClient.Post(SERVER, "text/xml; charset=utf-8", bytes.NewBufferString(QUERY))
	if err != nil {return err}
	b, err := ioutil.ReadAll(respx.Body)
	if err != nil {return err}
	in := string(b)
	var envelope SoapEnvelope
	parser := xml.NewDecoder(bytes.NewBufferString(in))
	err = parser.DecodeElement(&envelope, nil)
	if err != nil {return err}
	resp.Message = envelope.Body.VerifyEmailResponse.VerifyEmailResult.ResponseText
	respx.Body.Close()
	return nil
}

func GetSoapEnvelope() (envelope *SoapEnvelope) {
    httpClient := new(http.Client)
    resp, err := httpClient.Post(SERVER, "text/xml; charset=utf-8", bytes.NewBufferString(QUERY))
    if err != nil {}
    b, err := ioutil.ReadAll(resp.Body)
    if err != nil {}
    in := string(b)
    envelope = new(SoapEnvelope)
    parser := xml.NewDecoder(bytes.NewBufferString(in))
    err = parser.DecodeElement(&envelope, nil)
    if err != nil {}
    resp.Body.Close()
    return envelope
}

func main () {
    env := GetSoapEnvelope()
    fmt.Printf("%v", env.Body.VerifyEmailResponse.VerifyEmailResult.ResponseText)
}

type SoapItem struct {
    Number int
}

type SoapItems struct {
    Items []SoapItem "return>item"
}

    if err != nil {fmt.Println(err.Error())}
*/

