package my.endpoints.deprecated;
/*
import android.app.Activity;
import android.content.Context;
import android.os.AsyncTask;
import android.util.Pair;
import com.appspot.gcl_13.rest.Rest;
import com.google.api.client.extensions.android.http.AndroidHttp;
import com.google.api.client.googleapis.extensions.android.gms.auth.GoogleAccountCredential;
import com.google.api.client.googleapis.extensions.android.gms.auth.GooglePlayServicesAvailabilityIOException;
import com.google.api.client.googleapis.extensions.android.gms.auth.UserRecoverableAuthIOException;
import com.google.api.client.json.gson.GsonFactory;

import java.io.IOException;

class ServiceTask extends AsyncTask<Pair<Context,GoogleAccountCredential>, Void, Pair<Context,String>> {

    @Override
    protected Pair doInBackground(Pair<Context,GoogleAccountCredential>... arg) {
        String err;
        Context c = arg[0].first;
        GoogleAccountCredential u = arg[0].second;

        Rest.Builder endpoints = new Rest.Builder(AndroidHttp.newCompatibleTransport(), new GsonFactory(), u);
        Rest service = endpoints.build();

        try {
            err=service.google().storage().execute().getMessage();
        } catch (final GooglePlayServicesAvailabilityIOException availabilityException) {
            //int statusCode = GooglePlayServicesUtil.isGooglePlayServicesAvailable(c[0]);
            //int statusCode = availabilityException.getConnectionStatusCode();
            //GooglePlayServicesUtil.getErrorDialog(statusCode, this, 0).show();
            err= "GooglePlay Services not found! "+availabilityException.getConnectionStatusCode();
        } catch (UserRecoverableAuthIOException userRecoverableException) {
            ((Activity) c).startActivity(userRecoverableException.getIntent());
            err= "User Recoverable Auth IO Exception!";
        } catch (IOException e) {
            e.printStackTrace();
            err= "IO Exception!";
        }

        return Pair.create(c, err);
    }

    @Override
    protected void onPostExecute(Pair<Context,String> p) {
        MainActivity.toaster(p.first, p.second);
    }

}
*/