package my.endpoints;

import android.app.Activity;
import android.content.Context;
import android.os.AsyncTask;
import com.appspot.gcl_11.service.Service;
import com.appspot.gcl_11.service.model.Entity;
import com.appspot.gcl_11.service.model.Message;
import com.facebook.LoggingBehavior;
import com.facebook.Session;
import com.facebook.SessionState;
import com.facebook.Settings;
import com.google.api.client.extensions.android.http.AndroidHttp;
import com.google.api.client.googleapis.extensions.android.gms.auth.GooglePlayServicesAvailabilityIOException;
import com.google.api.client.googleapis.extensions.android.gms.auth.UserRecoverableAuthIOException;
import com.google.api.client.http.HttpHeaders;
import com.google.api.client.http.HttpRequest;
import com.google.api.client.http.HttpRequestInitializer;
import com.google.api.client.http.HttpTransport;
import com.google.api.client.json.gson.GsonFactory;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

class SubmitTask extends AsyncTask<Context, Void, Void> {
    Context context;

    @Override
    protected Void doInBackground(Context... arg) {
        context = arg[0];

        Settings.addLoggingBehavior(LoggingBehavior.INCLUDE_ACCESS_TOKENS);
        Session.StatusCallback callback = new Session.StatusCallback() {
            public void call(Session session, SessionState state, Exception exception) {
                Global g = Global.getInstance();
                String m = g.getMessage();
                if (exception != null) {session = createSession(Global.APP_ID);}
                sendRequest1(m, session.getAccessToken());
            }
        };

        Global g = Global.getInstance();
        String m = g.getMessage();
        Session session = createSession(Global.APP_ID);

        if (session.isOpened()) {
            sendRequest1(m, session.getAccessToken());
        } else {
            session.openForRead(new Session.OpenRequest((Activity) context).setCallback(callback));
        }

        return null;
    }

    private void sendRequest1(String m1, String token1) {
        final String token2 = token1;
        final String m2 = m1;

        Thread t = new Thread() {

            public void run() {
                Message message = new Message();
                message.setMessage(m2);
                List<Message> list = new ArrayList<Message>();
                list.add(message);
                Entity entity = new Entity();
                entity.setList(list);

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
                    service.datastore().submit(entity).execute();
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
                new ListTask().execute(context);
            }

        };

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

    //@Override
    //protected void onPostExecute(String err) {
    //    if (err!=null) MainActivity.toaster(context, err);
    //}

}