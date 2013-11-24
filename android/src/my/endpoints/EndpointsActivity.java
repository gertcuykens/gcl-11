package my.endpoints;

import android.app.Activity;
import android.content.Context;
import android.os.AsyncTask;
import android.os.Bundle;
import android.view.Gravity;
import android.view.View;
import android.widget.Toast;
import com.google.api.client.extensions.android.http.AndroidHttp;
import com.google.api.client.googleapis.extensions.android.gms.auth.GoogleAccountCredential;
import com.google.api.client.json.gson.GsonFactory;

import java.io.IOException;

public class EndpointsActivity extends Activity implements View.OnClickListener {
    public static final String AUDIENCE = "server:client_id:522156758812-09f5qv0e4gqjdjqfocerqcud5m5jutau.apps.googleusercontent.com";
    GoogleAccountCredential credential;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.endpoints);
        View userButton = findViewById(R.id.userButton);
        userButton.setOnClickListener(this);
        View getGreetingButton = findViewById(R.id.getGreetingButton);
        getGreetingButton.setOnClickListener(this);
        credential = GoogleAccountCredential.usingAudience(this, AUDIENCE);
    }

    private class RestTask extends AsyncTask<Void, Void, String> {
        private Context mContext;

        public RestTask(Context context) {
            mContext = context;
        }

        @Override
        protected String doInBackground(Void... unused) {
            String message = null;
            try {
                Endpoints.Builder endpoint = new Endpoints.Builder(AndroidHttp.newCompatibleTransport(), new GsonFactory(), credential);
                Endpoints service = endpoint.build();
                EndpointsResponse response = service.rest().getGreeting("0").execute();
                message=response.getMessage();
            } catch (IOException e) {
                e.printStackTrace();
            }
            return message;
        }
        @Override
        protected void onPostExecute(String message) {
            Toast toast = Toast.makeText(mContext, message, Toast.LENGTH_SHORT);
            toast.setGravity(Gravity.TOP|Gravity.RIGHT, 0, 0);
            toast.show();
        }
    }

    @Override
    public void onClick(View view) {
        Context context = view.getContext();
        switch(view.getId()) {
            case R.id.userButton:
                Toast toast = Toast.makeText(context, "Not implemented yet", Toast.LENGTH_SHORT);
                toast.setGravity(Gravity.TOP|Gravity.RIGHT, 0, 0);
                toast.show();
                break;
            case R.id.getGreetingButton:
                new RestTask(context).execute();
                break;
        }
    }

}