package com.appspot;

import android.accounts.AccountManager;
import android.app.Activity;
import android.content.Context;
import android.content.Intent;
import android.os.AsyncTask;
import android.os.Bundle;
import android.util.Pair;
import android.view.Gravity;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;
import android.widget.TextView;
import android.widget.Toast;

//import com.appspot.gcl_13.rest.Rest1;
//import com.appspot.gcl_13.rest.model.Multiply;

//import com.google.android.gms.common.AccountPicker;
//import com.google.android.gms.common.GooglePlayServicesUtil;
import com.google.api.client.extensions.android.http.AndroidHttp;
import com.google.api.client.googleapis.extensions.android.gms.auth.GoogleAccountCredential;
import com.google.api.client.googleapis.extensions.android.gms.auth.GooglePlayServicesAvailabilityIOException;
import com.google.api.client.googleapis.extensions.android.gms.auth.UserRecoverableAuthIOException;
import com.google.api.client.json.gson.GsonFactory;
import com.google.api.services.storage.Storage;

import java.io.FileOutputStream;
import java.io.IOException;
import java.io.OutputStream;
import java.util.Arrays;

public class EndpointsActivity extends Activity implements View.OnClickListener {
    private GoogleAccountCredential user;
    private GoogleAccountCredential user2;
    //private Rest service;
    private Button userButton;
    private TextView userStatus;
    private EditText getGreetingValue;
    private EditText greetingValue;
    private EditText multiplyValue;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.endpoints);

        //String AUDIENCE = "server:client_id:1093123836049-ilqfjb2s2tqal2fobuduj8b790hnnfju.apps.googleusercontent.com";
        // user = GoogleAccountCredential.usingAudience(this, AUDIENCE);

        String SCOPE="https://www.googleapis.com/auth/devstorage.read_only";
        user2=GoogleAccountCredential.usingOAuth2(this, Arrays.asList(SCOPE.split(" ")));

        //Rest.Builder endpoints = new Rest.Builder(AndroidHttp.newCompatibleTransport(), new GsonFactory(), null);
        //service = endpoints.build();

        userStatus = (TextView) findViewById(R.id.userStatus);
        getGreetingValue = (EditText) findViewById(R.id.getGreetingValue);
        greetingValue = (EditText) findViewById(R.id.greetingValue);
        multiplyValue = (EditText) findViewById(R.id.multiplyValue);

        userButton = (Button) findViewById(R.id.userButton);
        userButton.setOnClickListener(this);

        Button getGreetingButton = (Button) findViewById(R.id.getGreetingButton);
        getGreetingButton.setOnClickListener(this);

        Button getListButton = (Button) findViewById(R.id.getListButton);
        getListButton.setOnClickListener(this);

        Button multiplyButton = (Button) findViewById(R.id.multiplyButton);
        multiplyButton.setOnClickListener(this);

        Button authenticatedButton = (Button) findViewById(R.id.authenticatedButton);
        authenticatedButton.setOnClickListener(this);

        Button soapButton = (Button) findViewById(R.id.soapButton);
        soapButton.setOnClickListener(this);

        Button datastoreButton = (Button) findViewById(R.id.datastoreButton);
        datastoreButton.setOnClickListener(this);

        Button storageButton = (Button) findViewById(R.id.storageButton);
        storageButton.setOnClickListener(this);
    }

    @Override
    public void onClick(View view) {
        Context context = view.getContext();
        switch(view.getId()) {
            case R.id.userButton:startActivityForResult(user2.newChooseAccountIntent(), 1); break;
            case R.id.getGreetingButton: new RestTask().execute(Pair.create(context, 1)); break;
            case R.id.getListButton: new RestTask().execute(Pair.create(context, 2)); break;
            case R.id.multiplyButton: new RestTask().execute(Pair.create(context, 3)); break;
            case R.id.authenticatedButton: new RestTask().execute(Pair.create(context, 4)); break;
            case R.id.soapButton: new RestTask().execute(Pair.create(context, 5)); break;
            case R.id.datastoreButton: new RestTask().execute(Pair.create(context, 6)); break;
            case R.id.storageButton: new StorageTask().execute(context); break;
        }
    }

    @Override
    protected void onActivityResult(int requestCode, int resultCode, Intent data) {
        super.onActivityResult(requestCode, resultCode, data);
        switch (requestCode) {
            case 1: //REQUEST_ACCOUNT_PICKER
                if (data != null && data.getExtras() != null) {
                    String accountName = data.getExtras().getString(AccountManager.KEY_ACCOUNT_NAME);
                    if (accountName != null) {
                        user2.setSelectedAccountName(accountName);
                        userStatus.setText(user2.getSelectedAccountName());
                        //Rest.Builder endpoints = new Rest.Builder(AndroidHttp.newCompatibleTransport(), new GsonFactory(), user2);
                        //service = endpoints.build();
                    }
                }
                break;
            case 2: //TODO: ???
                break;
        }
    }

    private class RestTask extends AsyncTask<Pair<Context,Integer>, Void, Pair<Context,String>> {
        @Override
        protected Pair doInBackground(Pair<Context,Integer>... p) {
            String text = null;
            int i1 = Integer.parseInt(getGreetingValue.getText().toString());
            int i2 = Integer.parseInt(multiplyValue.getText().toString());
            //Multiply m = new Multiply();
            //m.setMessage(greetingValue.getText().toString());
            //m.setTimes(i2);
            try {
                switch(p[0].second) {
                    //case 1: text = service.greetings().getGreeting(i1).execute().getMessage(); break;
                    //case 2: text = service.greetings().listGreeting().execute().getItems().toString(); break;
                    //case 3: text = service.greetings().multiply(m).execute().getMessage(); break;
                    //case 4: text = service.greetings().authed().execute().getMessage(); break;
                    //case 5: text = service.greetings().soap().execute().getMessage(); break;
                    //case 6: text = service.greetings().datastore().execute().getMessage(); break;
                }
            } catch (Exception e) {
                e.printStackTrace();
            }
            return Pair.create(p[0].first, text);
        }
        @Override
        protected void onPostExecute(Pair<Context,String> p) {toaster(p.first, p.second);}
    }

    private class StorageTask extends AsyncTask<Context, Void, Pair<Context,String>> {
        @Override
        protected Pair doInBackground(Context... c) {
            String err;
            com.google.api.services.storage.Storage storageService = new com.google.api.services.storage.Storage.Builder(AndroidHttp.newCompatibleTransport(), new GsonFactory(), user2).setApplicationName("gcl-storage").build();
            try {
                Storage.Objects.Get getObject = storageService.objects().get("gcl-storage", "test.txt");
                String appPath = c[0].getFilesDir().getAbsolutePath();
                java.io.File parentDir = new java.io.File(appPath);
                OutputStream out = new FileOutputStream(new java.io.File(parentDir,"GERT_TEST.TXT"));
                getObject.getMediaHttpDownloader().setDirectDownloadEnabled(true);
                getObject.executeMediaAndDownloadTo(out);
                err= "Download complete "+parentDir.getPath();
            } catch (final GooglePlayServicesAvailabilityIOException availabilityException) {
                //int statusCode = GooglePlayServicesUtil.isGooglePlayServicesAvailable(c[0]);
                //int statusCode = availabilityException.getConnectionStatusCode();
                //GooglePlayServicesUtil.getErrorDialog(statusCode, this, 0).show();
                err= "GooglePlay Services not found! "+availabilityException.getConnectionStatusCode();
            } catch (UserRecoverableAuthIOException userRecoverableException) {
                startActivityForResult(userRecoverableException.getIntent(), 2);
                err= "User Recoverable Auth IO Exception!";
            } catch (IOException e) {
                err= "IO Exception!";
            }
            return Pair.create(c[0], err);
        }
        @Override
        protected void onPostExecute(Pair<Context,String> p) {toaster(p.first, p.second);}
    }

    private void toaster(Context c, String s) {
        Toast toast = Toast.makeText(c, s, Toast.LENGTH_SHORT);
        toast.setGravity(Gravity.TOP|Gravity.RIGHT, 0, 0);
        toast.show();
    }

}