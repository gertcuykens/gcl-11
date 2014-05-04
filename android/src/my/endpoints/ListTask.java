package my.endpoints;

import android.app.Activity;
import android.content.Context;
import android.os.AsyncTask;
import android.util.Log;
import android.widget.TextView;
import com.appspot.gcl_11.service.Service;
import com.appspot.gcl_11.service.model.Entity;
import com.appspot.gcl_11.service.model.Message;
import com.facebook.*;
import com.google.api.client.extensions.android.http.AndroidHttp;
import com.google.api.client.googleapis.extensions.android.gms.auth.GooglePlayServicesAvailabilityIOException;
import com.google.api.client.googleapis.extensions.android.gms.auth.UserRecoverableAuthIOException;
import com.google.api.client.http.HttpHeaders;
import com.google.api.client.http.HttpRequest;
import com.google.api.client.http.HttpRequestInitializer;
import com.google.api.client.http.HttpTransport;
import com.google.api.client.json.gson.GsonFactory;

import java.io.IOException;

class ListTask extends AsyncTask<Context, Void, Void> {
    Context context;

    @Override
    protected Void doInBackground(Context... arg) {
        context = arg[0];

        Settings.addLoggingBehavior(LoggingBehavior.INCLUDE_ACCESS_TOKENS);
        Session.StatusCallback callback = new Session.StatusCallback() {
            public void call(Session session, SessionState state, Exception exception) {
                if (exception != null) {
                    Log.i("graph", "----------------------------");
                    Log.i("graph", String.valueOf(exception));
                    Log.i("graph", "----------------------------");
                    session.close();
                    session = createSession();
                    sendRequest1(session.getAccessToken());
                } else {
                    session = createSession();
                    sendRequest1(session.getAccessToken());
                }
            }
        };

        Session session = createSession();
        if (session.isOpened()) {
            Log.i("graph", "----------------------------");
            Log.i("graph", session.getAccessToken());
            Log.i("graph", "----------------------------");
            sendRequest1(session.getAccessToken());
        } else {
            session.openForRead( new Session.OpenRequest((Activity) context).setCallback(callback) );
        }

        return null;
    }

    private void sendRequest1(String token1) {

        final String token2= token1;
        Thread t = new Thread() {
            public void run() {
                Entity entity = new Entity();
                Entity result = new Entity();
                try {
                    class Init implements HttpRequestInitializer {
                        public void initialize(HttpRequest request) {
                            HttpHeaders headers = new HttpHeaders();
                            headers.setAuthorization("Bearer "+token2);
                            request.setHeaders(headers);
                        }
                    }

                    HttpTransport transport = AndroidHttp.newCompatibleTransport();
                    Service.Builder endpoints = new Service.Builder(transport, new GsonFactory(), new Init()).setApplicationName(Global.APP_NAME);
                    Service service = endpoints.build();
                    result = service.datastore().list(entity).execute();
                } catch (final GooglePlayServicesAvailabilityIOException availabilityException) {
                    //int statusCode = GooglePlayServicesUtil.isGooglePlayServicesAvailable(c[0]);
                    //int statusCode = availabilityException.getConnectionStatusCode();
                    //GooglePlayServicesUtil.getErrorDialog(statusCode, this, 0).show();
                    //err = "GooglePlay Services not found! " + availabilityException.getConnectionStatusCode();
                } catch (UserRecoverableAuthIOException userRecoverableException) {
                    ((Activity) context).startActivity(userRecoverableException.getIntent());
                    //err = "User Recoverable Auth IO Exception!";
                } catch (IOException e) {
                    e.printStackTrace();
                    //err = "IO Exception!";
                }
                //if (err!=null) MainActivity.toaster(context, err);
                printEntity(context, result);
            }
        };

        t.start();
    }

    private void printEntity(Context c1, Entity e1) {

        class PrintEntity implements Runnable {
            Context context;
            Entity entity;
            PrintEntity(Context c2, Entity e2) {
                context = c2;
                entity = e2;
            }
            public void run() {
                ((Activity)context).runOnUiThread(new Runnable(){
                    public void run(){

                        String s="";
                        if (entity!=null) {
                            if (!entity.isEmpty()) {
                                for (Message message : entity.getList()) {
                                    s = s + String.format("%s\n", message.getMessage());
                                }
                            }
                        }
                        TextView textViewResults = (TextView) ((Activity)context).findViewById(R.id.textViewResults);
                        textViewResults.setText(s);

                    }
                });
            }
        }

        Thread t = new Thread(new PrintEntity(c1, e1));
        t.start();

    }

    private Session createSession() {
        Session activeSession = Session.getActiveSession();
        if (activeSession == null || activeSession.getState().isClosed()) {
            activeSession = new Session.Builder(context).setApplicationId(Global.APP_ID).build();
            Session.setActiveSession(activeSession);
        }
        return activeSession;
    }

    //@Override
    //protected void onPostExecute(String err) {
    //   if (err!=null) MainActivity.toaster(context, err);
    //}

}