package my.endpoints;

import android.content.Context;
import com.facebook.Session;

public class Global {
    private static Global instance;
    static final String APP_ID = "249348058430770";
    static final String APP_NAME = "gcl-11";
    private static String message;
    private static String graph;
    private Global(){}
    public String getMessage(){return message;}
    public String getGraph(){return graph;}
    public void setMessage(String s){Global.message=s;}
    public void setGraph(String s){Global.graph=s;}
    public static Session createSession(Context context) {
        Session activeSession = Session.getActiveSession();
        if (activeSession == null || activeSession.getState().isClosed()) {
            activeSession = new Session.Builder(context).setApplicationId(Global.APP_ID).build();
            Session.setActiveSession(activeSession);
        }
        return activeSession;
    }
    public static synchronized Global getInstance(){
        if(instance==null){instance=new Global();}
        return instance;
    }
}
