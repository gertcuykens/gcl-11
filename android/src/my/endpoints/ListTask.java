package my.endpoints;

import android.app.Activity;
import android.content.Context;
import android.os.AsyncTask;
import android.widget.TextView;
import com.appspot.gcl_11.service.Service;
import com.appspot.gcl_11.service.model.Entity;
import com.appspot.gcl_11.service.model.Message;
import com.google.api.client.extensions.android.http.AndroidHttp;
import com.google.api.client.googleapis.extensions.android.gms.auth.GooglePlayServicesAvailabilityIOException;
import com.google.api.client.googleapis.extensions.android.gms.auth.UserRecoverableAuthIOException;
import com.google.api.client.json.gson.GsonFactory;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

class ListTask extends AsyncTask<Context, Void, String> {
    Context context;
    Service service;

    @Override
    protected String doInBackground(Context... arg) {
        context = arg[0];

        String err=null;
        Entity entity=new Entity();
        Entity result=new Entity();

        try {
            Service.Builder  endpoints = new Service.Builder(AndroidHttp.newCompatibleTransport(), new GsonFactory(), null);
            service = endpoints.build();
            result=service.datastore().list(entity).execute();
        } catch (final GooglePlayServicesAvailabilityIOException availabilityException) {
            //int statusCode = GooglePlayServicesUtil.isGooglePlayServicesAvailable(c[0]);
            //int statusCode = availabilityException.getConnectionStatusCode();
            //GooglePlayServicesUtil.getErrorDialog(statusCode, this, 0).show();
            err= "GooglePlay Services not found! "+availabilityException.getConnectionStatusCode();
        } catch (UserRecoverableAuthIOException userRecoverableException) {
            ((Activity) context).startActivity(userRecoverableException.getIntent());
            err= "User Recoverable Auth IO Exception!";
        } catch (IOException e) {
            e.printStackTrace();
            err= "IO Exception!";
        }

        printEntity(context, result);
        return err;
    }

    @Override
    protected void onPostExecute(String err) {
        if (err!=null) MainActivity.toaster(context, err);
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
                        for (Message message : entity.getList()) {
                            s = s + String.format("%s\n", message.getMessage());
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

}