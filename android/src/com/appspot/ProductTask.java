package com.appspot;

import android.content.Context;
import android.os.AsyncTask;
import android.util.Pair;
import com.google.api.client.googleapis.extensions.android.gms.auth.GoogleAccountCredential;

class ProductTask extends AsyncTask<Pair<Context,GoogleAccountCredential>, Void, Pair<Context,String>> {

    @Override
    protected Pair doInBackground(Pair<Context,GoogleAccountCredential>... arg) {
        String err;
        Context c = arg[0].first;
        GoogleAccountCredential u = arg[0].second;

        err="NOT IMP";
        return Pair.create(c, err);
    }

    @Override
    protected void onPostExecute(Pair<Context,String> p) {
        MainActivity.toaster(p.first, p.second);
    }

}