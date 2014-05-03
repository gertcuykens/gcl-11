package my.endpoints;

import android.app.Activity;
import android.content.Context;
import android.os.AsyncTask;
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

class SubmitTask extends AsyncTask<Context, Void, String> {
    Context context;
    Service service;

    @Override
    protected String doInBackground(Context... arg) {
        context = arg[0];

        Global g = Global.getInstance();
        String m = g.getMessage();

        String err=null;
        Message message =new Message();
        message.setMessage(m);
        List<Message> list=new ArrayList<Message>();
        list.add(message);
        Entity entity=new Entity();
        entity.setList(list);

        try {
            Service.Builder  endpoints = new Service.Builder(AndroidHttp.newCompatibleTransport(), new GsonFactory(), null);
            service = endpoints.build();
            service.datastore().submit(entity).execute();
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

        new ListTask().execute(context);
        return err;
    }

    @Override
    protected void onPostExecute(String err) {
        if (err!=null) MainActivity.toaster(context, err);
    }

}
