package my.endpoints;

import android.app.Activity;
import android.content.Context;
import android.os.AsyncTask;
import android.os.Bundle;
import android.util.Log;
import com.facebook.*;
import com.facebook.model.GraphObject;
import org.json.JSONException;
import org.json.JSONObject;

class RequestTask extends AsyncTask<Context, Void, Void> {
    Context context;

    @Override
    protected Void doInBackground(Context... arg) {
        context = arg[0];

        Session session = Global.createSession(context);
        if (session.isOpened()) {
            Global g = Global.getInstance();
            String m = g.getMessage();
            String f = g.getGraph();
            Bundle p = new Bundle();
            p.putString("message",m);
            p.putString("name","");
            p.putString("link","");
            p.putString("picture","");
            //p.putString("access_token","");
            sendRequest1(context, session, f, p);
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
                                    s = s + String.format("Error gcl-11: %s", error.getErrorMessage());
                                    Log.i("graph", "----------------------------");
                                    Log.e("graph", error.getErrorMessage());
                                    Log.i("graph", "----------------------------");
                                } else {
                                    JSONObject graphResponse = graphObject.getInnerJSONObject();
                                    try {
                                        String postId = graphResponse.getString("id");
                                        s = s + String.format("%s: gcl-11", postId);
                                        Log.i("graph", "----------------------------");
                                        Log.i("graph", postId);
                                        Log.i("graph", "----------------------------");
                                    } catch (JSONException e) {
                                        s = s + String.format("Error gcl-11: %s", e.getMessage());
                                        Log.i("graph", "----------------------------");
                                        Log.e("graph", e.getMessage());
                                        Log.i("graph", "----------------------------");
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

    //@Override
    //protected void onPostExecute(String err) {
    //    if (err!=null) MainActivity.toaster(context, err);
    //}

}
