package my.endpoints;

import android.app.Activity;
import android.app.AlertDialog;
import android.content.Context;
import android.os.AsyncTask;
import android.os.Bundle;
import android.util.Log;
import android.widget.TextView;
import com.appspot.gcl_11.service.Service;
import com.appspot.gcl_11.service.model.Entity;
import com.appspot.gcl_11.service.model.Message;
import com.facebook.*;
import com.facebook.model.GraphObject;
import com.google.api.client.extensions.android.http.AndroidHttp;
import com.google.api.client.googleapis.extensions.android.gms.auth.GooglePlayServicesAvailabilityIOException;
import com.google.api.client.googleapis.extensions.android.gms.auth.UserRecoverableAuthIOException;
import com.google.api.client.json.gson.GsonFactory;
import org.json.JSONException;
import org.json.JSONObject;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

class RequestTask extends AsyncTask<Context, Void, Void> {
    Context context;

    @Override
    protected Void doInBackground(Context... arg) {
        context = arg[0];

        Settings.addLoggingBehavior(LoggingBehavior.INCLUDE_ACCESS_TOKENS);
        Session.StatusCallback callback = new Session.StatusCallback() {
            public void call(Session session, SessionState state, Exception exception) {
                Global g = Global.getInstance();
                String m = g.getMessage();
                String f = g.getGraph();
                Bundle p = new Bundle();
                p.putString("message",m);
                p.putString("name","");
                p.putString("link","");
                p.putString("picture","");
                //p.putString("access_token","");
                if (exception != null) {session = createSession(g.APP_ID);}
                sendRequest1(context, session, f, p);
            }
        };

        Global g = Global.getInstance();
        String m = g.getMessage();
        String f = g.getGraph();
        Bundle p = new Bundle();
        p.putString("message",m);
        p.putString("name","");
        p.putString("link","");
        p.putString("picture","");
        //p.putString("access_token","");

        Session session = createSession(g.APP_ID);
        if (session.isOpened()) {
            sendRequest1(context, session, f, p);
        } else {
            session.openForRead(new Session.OpenRequest((Activity) context).setCallback(callback));
        }

        return null;
    }

    private void sendRequest1(Context c1, Session s1, String g1, Bundle p1) {

        class RunUI implements Runnable {
            Context context;
            Session session;
            String graph;
            Bundle post;

            RunUI(Context c2, Session s2, String g2, Bundle p2) {
                context = c2;
                session = s2;
                graph = g2;
                post = p2;
            }

            public void run() {
                ((Activity)context).runOnUiThread(new Runnable(){
                    public void run(){

                        Request.Callback callback = new Request.Callback() {
                            public void onCompleted(Response response) {
                                GraphObject graphObject = response.getGraphObject();
                                FacebookRequestError error = response.getError();
                                String s = "";

                                if (error != null) {
                                    s = s + String.format("Error gcl-11: %s\n", error.getErrorMessage());
                                    Log.e("FACEBOOK ERROR", "" + error.getErrorMessage());
                                } else {
                                    JSONObject graphResponse = graphObject.getInnerJSONObject();
                                    try {
                                        String postId = graphResponse.getString("id");
                                        s = s + String.format("%s: gcl-11\n", postId);
                                        Log.i ("graph", postId);
                                    } catch (JSONException e) {
                                        s = s + String.format("Error gcl-11: %s\n", e.getMessage());
                                        Log.e("FACEBOOK ERROR", e.getMessage());
                                    }
                                }

                                MainActivity.toaster(context, s);
                            }
                        };

                        Request request = new Request(session, graph, post, HttpMethod.POST, callback);
                        RequestAsyncTask task = new RequestAsyncTask(request);
                        task.execute();

                    }
                });
            }
        }

        Thread t = new Thread(new RunUI(c1, s1, g1, p1));
        t.start();

    }

    private Session createSession(String APP_ID) {
        Session activeSession = Session.getActiveSession();
        if (activeSession == null || activeSession.getState().isClosed()) {
            activeSession = new Session.Builder(context).setApplicationId(APP_ID).build();
            Session.setActiveSession(activeSession);
        }
        return activeSession;
    }

}

//@Override
//protected void onPostExecute(String err) {
//    if (err!=null) MainActivity.toaster(context, err);
//}

/*
    private void sendRequests(Bundle postParams, String graph) {

        Request request =new Request(session, "me", null, null, new Request.Callback() {
            public void onCompleted(Response response) {
                GraphObject graphObject = response.getGraphObject();
                FacebookRequestError error = response.getError();
                String s = "";

                if (graphObject != null) {
                    if (graphObject.getProperty("id") != null) {
                        s = s + String.format("%s: %s\n", graphObject.getProperty("id"), graphObject.getProperty("name"));
                    } else {
                        s = s + String.format("%s: <no such id>\n", "me");
                    }
                } else if (error != null) {
                    s = s + String.format("Error: %s", error.getErrorMessage());
                }

                MainActivity.toaster(context, s);
            }
        });

        //pendingRequest = false;

        //request.executeAsync();

    }
*/