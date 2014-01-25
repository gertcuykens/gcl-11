package com.appspot;

import android.app.Activity;
import android.content.Context;
import android.os.AsyncTask;
import android.util.Pair;
import com.google.api.client.extensions.android.http.AndroidHttp;
import com.google.api.client.googleapis.extensions.android.gms.auth.GoogleAccountCredential;
import com.google.api.client.googleapis.extensions.android.gms.auth.GooglePlayServicesAvailabilityIOException;
import com.google.api.client.googleapis.extensions.android.gms.auth.UserRecoverableAuthIOException;
import com.google.api.client.json.gson.GsonFactory;
import com.google.api.services.storage.Storage;

import java.io.FileOutputStream;
import java.io.IOException;
import java.io.OutputStream;

class StorageTask extends AsyncTask<Pair<Context,GoogleAccountCredential>, Void, Pair<Context,String>> {

    @Override
    protected Pair doInBackground(Pair<Context,GoogleAccountCredential>... arg) {
        String err;
        Context c = arg[0].first;
        GoogleAccountCredential u = arg[0].second;
        com.google.api.services.storage.Storage storageService = new com.google.api.services.storage.Storage.Builder(AndroidHttp.newCompatibleTransport(), new GsonFactory(), u).setApplicationName("gcl-storage").build();
        try {
            String appPath = c.getFilesDir().getAbsolutePath();
            java.io.File parentDir = new java.io.File(appPath);
            OutputStream out = new FileOutputStream(new java.io.File(parentDir,"GERT_TEST.TXT"));
            Storage.Objects.Get getObject = storageService.objects().get("gcl-storage", "test.txt");
            getObject.getMediaHttpDownloader().setDirectDownloadEnabled(true);
            getObject.executeMediaAndDownloadTo(out);
            err= "Download complete for "+u.getSelectedAccountName()+" at "+parentDir.getPath();
        } catch (final GooglePlayServicesAvailabilityIOException availabilityException) {
            //int statusCode = GooglePlayServicesUtil.isGooglePlayServicesAvailable(c[0]);
            //int statusCode = availabilityException.getConnectionStatusCode();
            //GooglePlayServicesUtil.getErrorDialog(statusCode, this, 0).show();
            err= "GooglePlay Services not found! "+availabilityException.getConnectionStatusCode();
        } catch (UserRecoverableAuthIOException userRecoverableException) {
            ((Activity) c).startActivity(userRecoverableException.getIntent());
            err= "User Recoverable Auth IO Exception!";
        } catch (IOException e) {
            err= "IO Exception!";
        }
        return Pair.create(c, err);
    }

    @Override
    protected void onPostExecute(Pair<Context,String> p) {
        MainActivity.toaster(p.first, p.second);
    }

}
