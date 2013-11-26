package my.endpoints;

import android.accounts.AccountManager;
import android.app.Activity;
import android.content.Context;
import android.content.Intent;
import android.os.AsyncTask;
import android.os.Bundle;
import android.view.Gravity;
import android.view.View;
import android.widget.Button;
import android.widget.TextView;
import android.widget.Toast;
import com.google.api.client.extensions.android.http.AndroidHttp;
import com.google.api.client.googleapis.extensions.android.gms.auth.GoogleAccountCredential;
import com.google.api.client.json.gson.GsonFactory;

import java.io.IOException;

public class EndpointsActivity extends Activity implements View.OnClickListener {
    private GoogleAccountCredential credential;
    private TextView userStatus;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.endpoints);

        String AUDIENCE = "server:client_id:522156758812-09f5qv0e4gqjdjqfocerqcud5m5jutau.apps.googleusercontent.com";
        credential = GoogleAccountCredential.usingAudience(this, AUDIENCE);

        userStatus = (TextView) findViewById(R.id.userStatus);

        Button userButton = (Button) findViewById(R.id.userButton);
        userButton.setOnClickListener(this);

        Button getGreetingButton = (Button) findViewById(R.id.getGreetingButton);
        getGreetingButton.setOnClickListener(this);
    }

    @Override
    protected void onActivityResult(int requestCode, int resultCode, Intent data) {
        super.onActivityResult(requestCode, resultCode, data);
        switch (requestCode) {
            case 1: //REQUEST_ACCOUNT_PICKER
                if (data != null && data.getExtras() != null) {
                    String accountName = data.getExtras().getString(AccountManager.KEY_ACCOUNT_NAME);
                    if (accountName != null) {
                        credential.setSelectedAccountName(accountName);
                        userStatus.setText(credential.getSelectedAccountName());
                    }
                }
                break;
        }
    }

    private class RestTask extends AsyncTask<Void, Void, String> {
        private Context mContext;
        public RestTask(Context context) {mContext = context;}
        @Override
        protected String doInBackground(Void... unused) {
            String message = null;
            try {
                Endpoints.Builder endpoint = new Endpoints.Builder(AndroidHttp.newCompatibleTransport(), new GsonFactory(), credential)
                        .setApplicationName("gcl-11");
                Endpoints service = endpoint.build();
                EndpointsResponse response = service.rest().getGreeting("0").execute();
                message=response.getMessage();
            } catch (IOException e) {
                e.printStackTrace();
            }
            return message;
        }
        @Override
        protected void onPostExecute(String message) {toaster(mContext, message);}
    }

    @Override
    public void onClick(View view) {
        Context context = view.getContext();
        switch(view.getId()) {
            case R.id.userButton:
                startActivityForResult(credential.newChooseAccountIntent(),1);
                break;
            case R.id.getGreetingButton:
                new RestTask(context).execute();
                break;
        }
    }

    private void toaster(Context c, String s) {
        Toast toast = Toast.makeText(c, s, Toast.LENGTH_SHORT);
        toast.setGravity(Gravity.TOP|Gravity.RIGHT, 0, 0);
        toast.show();
    }

}