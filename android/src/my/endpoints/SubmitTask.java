package my.endpoints;

import android.app.Activity;
import android.content.Context;
import android.os.AsyncTask;
import android.util.Log;
import com.appspot.gcl_11.service.Service;
import com.appspot.gcl_11.service.model.Entity;
import com.appspot.gcl_11.service.model.Message;
import com.facebook.Session;
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

        Session session = Global.createSession(context);
        if (session.isOpened()) {
            Global g = Global.getInstance();
            String m = g.getMessage();
            sendRequest1(m, session.getAccessToken());
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
                            headers.setAuthorization("Bearer " + token2);
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
                    Log.i("graph", "----------------------------");
                    Log.i("graph", "GooglePlay Services not found! " + availabilityException.getConnectionStatusCode());
                    Log.i("graph", "----------------------------");
                } catch (UserRecoverableAuthIOException userRecoverableException) {
                    ((Activity) context).startActivity(userRecoverableException.getIntent());
                    Log.i("graph", "----------------------------");
                    Log.i("graph", "User Recoverable Auth IO Exception!");
                    Log.i("graph", "----------------------------");
                } catch (IOException e) {
                    //e.printStackTrace();
                    Log.i("graph", "----------------------------");
                    Log.i("graph", "IO Exception!");
                    Log.i("graph", "----------------------------");
                }

                new ListTask().execute(context);
            }

        };

        t.start();
    }

    //@Override
    //protected void onPostExecute(String err) {
    //    if (err!=null) MainActivity.toaster(context, err);
    //}

}
