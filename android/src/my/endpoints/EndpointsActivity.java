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
        credential.setSelectedAccountName("gert.cuykens@gmail.com");

        userStatus = (TextView) findViewById(R.id.userStatus);
        userStatus.setText("gert.cuykens@gmail.com");

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

    private class RestTask extends AsyncTask<Void, Void, String> {
        private Context mContext;
        public RestTask(Context context) {mContext = context;}
        @Override
        protected String doInBackground(Void... unused) {
            String text = null;
            try {
                EndpointsClient.Builder endpoints = new EndpointsClient.Builder(AndroidHttp.newCompatibleTransport(), new GsonFactory(), credential);
                EndpointsClient service = endpoints.build();
                Message message = new Message();
                message.setMessage("hello ");

                Message response = service.get("response/0").execute();
                //Message response = service.get("response").execute();
                //Message response = service.post("response/2",message).execute();
                //Message response = service.post("greetings/authed",message).execute();
                //Message response = service.get("greetings/soap").execute()
                //Message response = service.get("greetings/datastore").execute();

                text=response.getMessage();
                //text=response.getItems().toString();
            } catch (IOException e) {
                e.printStackTrace();
            }
            return text;
        }
        @Override
        protected void onPostExecute(String text) {toaster(mContext, text);}
    }

    private void toaster(Context c, String s) {
        Toast toast = Toast.makeText(c, s, Toast.LENGTH_SHORT);
        toast.setGravity(Gravity.TOP|Gravity.RIGHT, 0, 0);
        toast.show();
    }

}