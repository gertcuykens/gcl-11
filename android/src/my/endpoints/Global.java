package my.endpoints;

public class Global {
    private static Global instance;
    static final String APP_ID = "249348058430770";
    private static String message;
    private static String graph;
    private Global(){}
    public String getMessage(){return message;}
    public String getGraph(){return graph;}
    public void setMessage(String s){Global.message=s;}
    public void setGraph(String s){Global.graph=s;}
    public static synchronized Global getInstance(){
        if(instance==null){instance=new Global();}
        return instance;
    }
}
