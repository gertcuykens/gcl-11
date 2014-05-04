package my.endpoints.deprecated;
/*
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
import com.appspot.gcl_11.rest.Rest;
import com.appspot.gcl_11.rest.model.Multiply;
import com.google.api.client.extensions.android.http.AndroidHttp;
import com.google.api.client.googleapis.extensions.android.gms.auth.GoogleAccountCredential;
import com.google.api.client.json.gson.GsonFactory;

public class MainActivity3 extends Activity implements View.OnClickListener {
    private GoogleAccountCredential user;
    private Rest service;
    private Button userButton;
    private TextView userStatus;
    private EditText getGreetingValue;
    private EditText greetingValue;
    private EditText multiplyValue;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.endpoints1);

        String AUDIENCE = "server:client_id:522156758812-09f5qv0e4gqjdjqfocerqcud5m5jutau.apps.googleusercontent.com";
        user = GoogleAccountCredential.usingAudience(this, AUDIENCE);

        Rest.Builder endpoints = new Rest.Builder(AndroidHttp.newCompatibleTransport(), new GsonFactory(), null);
        service = endpoints.build();

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
    }

    @Override
    public void onClick(View view) {
        Context context = view.getContext();
        switch(view.getId()) {
            case R.id.userButton:
                if (userButton.getText().equals("Sign Out")) {
                    Rest.Builder endpoints = new Rest.Builder(AndroidHttp.newCompatibleTransport(), new GsonFactory(), null);
                    service = endpoints.build();
                    userStatus.setText("Not signed in");
                    userButton.setText("Sign In");
                }else{
                    startActivityForResult(user.newChooseAccountIntent(), 1);
                }
                break;
            case R.id.getGreetingButton: new RestTask().execute(Pair.create(context, 1)); break;
            case R.id.getListButton: new RestTask().execute(Pair.create(context, 2)); break;
            case R.id.multiplyButton: new RestTask().execute(Pair.create(context, 3)); break;
            case R.id.authenticatedButton: new RestTask().execute(Pair.create(context, 4)); break;
            case R.id.soapButton: new RestTask().execute(Pair.create(context, 5)); break;
            case R.id.datastoreButton: new RestTask().execute(Pair.create(context, 6)); break;
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
                        user.setSelectedAccountName(accountName);
                        Rest.Builder endpoints = new Rest.Builder(AndroidHttp.newCompatibleTransport(), new GsonFactory(), user);
                        service = endpoints.build();
                        userStatus.setText(user.getSelectedAccountName());
                        userButton.setText("Sign Out");
                    }
                }
                break;
        }
    }

    private class RestTask extends AsyncTask<Pair<Context,Integer>, Void, Pair<Context,String>> {
        @Override
        protected Pair doInBackground(Pair<Context,Integer>... p) {
            String text = null;
            int i1 = Integer.parseInt(getGreetingValue.getText().toString());
            int i2 = Integer.parseInt(multiplyValue.getText().toString());
            Multiply m = new Multiply();
            m.setMessage(greetingValue.getText().toString());
            m.setTimes(i2);
            try {
                switch(p[0].second) {
                    case 1: text = service.greetings().getGreeting(i1).execute().getMessage(); break;
                    case 2: text = service.greetings().listGreeting().execute().getItems().toString(); break;
                    case 3: text = service.greetings().multiply(m).execute().getMessage(); break;
                    case 4: text = service.greetings().authed().execute().getMessage(); break;
                    case 5: text = service.greetings().soap().execute().getMessage(); break;
                    case 6: text = service.greetings().datastore().execute().getMessage(); break;
                }
            } catch (Exception e) {
                e.printStackTrace();
            }
            return Pair.create(p[0].first, text);
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
*/